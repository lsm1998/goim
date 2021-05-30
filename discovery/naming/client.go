package naming

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	ecode "github.com/bilibili/kratos/pkg/ecode"
	log "github.com/bilibili/kratos/pkg/log"
	http "github.com/bilibili/kratos/pkg/net/http/blademaster"
	xtime "github.com/bilibili/kratos/pkg/time"
)

const (
	_registerURL = "http://%s/discovery/register"
	_setURL      = "http://%s/discovery/set"
	_cancelURL   = "http://%s/discovery/cancel"
	_renewURL    = "http://%s/discovery/renew"
	_pollURL     = "http://%s/discovery/polls"

	_registerGap = 30 * time.Second

	_statusUP = "1"

	_appid = "infra.discovery"
)

var (
	_ Builder  = &Discovery{}
	_ Registry = &Discovery{}

	// ErrDuplication duplication treeid.
	ErrDuplication = errors.New("discovery: instance duplicate registration")
)

// Config discovery configures.
type Config struct {
	Nodes  []string
	Region string
	Zone   string
	Env    string
	Host   string
}

type appData struct {
	Instances map[string][]*Instance `json:"instances"`
	LastTs    int64                  `json:"latest_timestamp"`
}

// Discovery is discovery client.
type Discovery struct {
	c          *Config
	once       sync.Once
	ctx        context.Context
	cancelFunc context.CancelFunc
	httpClient *http.Client

	node    atomic.Value
	nodeIdx uint64

	mutex       sync.RWMutex
	apps        map[string]*appInfo
	registry    map[string]struct{}
	lastHost    string
	cancelPolls context.CancelFunc

	delete chan *appInfo
}

type appInfo struct {
	resolver map[*Resolve]struct{}
	zoneIns  atomic.Value
	lastTs   int64 // latest timestamp
}

func fixConfig(c *Config) {
	if len(c.Nodes) == 0 {
		panic("conf nodes can not be nil")
	}
	if c.Region == "" {
		c.Region = os.Getenv("REGION")
	}
	if c.Zone == "" {
		c.Zone = os.Getenv("ZONE")
	}
	if c.Env == "" {
		c.Env = os.Getenv("DEPLOY_ENV")
	}
	if c.Host == "" {
		c.Host, _ = os.Hostname()
	}
}

// New new a discovery client.
func New(c *Config) (d *Discovery) {
	fixConfig(c)
	ctx, cancel := context.WithCancel(context.Background())
	d = &Discovery{
		c:          c,
		ctx:        ctx,
		cancelFunc: cancel,
		apps:       map[string]*appInfo{},
		registry:   map[string]struct{}{},
		delete:     make(chan *appInfo, 10),
	}
	// httpClient
	cfg := &http.ClientConfig{
		Dial:      xtime.Duration(3 * time.Second),
		KeepAlive: xtime.Duration(40 * time.Second),
		Timeout:   xtime.Duration(40 * time.Second),
	}
	d.httpClient = http.NewClient(cfg)
	// discovery self
	resolver := d.Build(_appid)
	event := resolver.Watch()
	_, ok := <-event
	if !ok {
		panic("discovery watch failed")
	}
	ins, ok := resolver.Fetch()
	if ok {
		d.newSelf(ins.Instances)
	}
	go d.selfproc(resolver, event)
	return
}

func (d *Discovery) selfproc(resolver Resolver, event <-chan struct{}) {
	for {
		_, ok := <-event
		if !ok {
			return
		}
		zones, ok := resolver.Fetch()
		if ok {
			d.newSelf(zones.Instances)
		}
	}
}

func (d *Discovery) newSelf(zones map[string][]*Instance) {
	d.mutex.Lock()
	ins, ok := zones[d.c.Zone]
	d.mutex.Unlock()
	if !ok {
		return
	}
	var nodes []string
	for _, in := range ins {
		for _, addr := range in.Addrs {
			u, err := url.Parse(addr)
			if err == nil && u.Scheme == "http" {
				nodes = append(nodes, u.Host)
			}
		}
	}
	// diff old nodes
	var olds int
	for _, n := range nodes {
		if node, ok := d.node.Load().([]string); ok {
			for _, o := range node {
				if o == n {
					olds++
					break
				}
			}
		}
	}
	if len(nodes) == olds {
		return
	}
	// FIXME: we should use rand.Shuffle() in golang 1.10
	shuffle(len(nodes), func(i, j int) {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	})
	d.node.Store(nodes)
}

// Build discovery resovler builder.
func (d *Discovery) Build(appid string) Resolver {
	r := &Resolve{
		id:    appid,
		d:     d,
		event: make(chan struct{}, 1),
	}
	d.mutex.Lock()
	app, ok := d.apps[appid]
	if !ok {
		app = &appInfo{
			resolver: make(map[*Resolve]struct{}),
		}
		d.apps[appid] = app
		cancel := d.cancelPolls
		if cancel != nil {
			cancel()
		}
	}
	app.resolver[r] = struct{}{}
	d.mutex.Unlock()
	if ok {
		select {
		case r.event <- struct{}{}:
		default:
		}
	}
	log.Info("discovery: AddWatch(%s) already watch(%v)", appid, ok)
	d.once.Do(func() {
		go d.serverproc()
	})
	return r
}

// Scheme return discovery's scheme
func (d *Discovery) Scheme() string {
	return "discovery"
}

// Resolve discveory resolver.
type Resolve struct {
	id    string
	event chan struct{}
	d     *Discovery
}

// Watch watch instance.
func (r *Resolve) Watch() <-chan struct{} {
	return r.event
}

// Fetch fetch resolver instance.
func (r *Resolve) Fetch() (ins *InstancesInfo, ok bool) {
	r.d.mutex.RLock()
	app, ok := r.d.apps[r.id]
	r.d.mutex.RUnlock()
	if ok {
		ins, ok = app.zoneIns.Load().(*InstancesInfo)
		return
	}
	return
}

// Close close resolver.
func (r *Resolve) Close() error {
	r.d.mutex.Lock()
	if app, ok := r.d.apps[r.id]; ok && len(app.resolver) != 0 {
		delete(app.resolver, r)
		// TODO: delete app from builder
	}
	r.d.mutex.Unlock()
	return nil
}

// Reload reload the config
func (d *Discovery) Reload(c *Config) {
	fixConfig(c)
	d.mutex.Lock()
	d.c = c
	d.mutex.Unlock()
}

// Close stop all running process including discovery and register
func (d *Discovery) Close() error {
	d.cancelFunc()
	return nil
}

// Register Register an instance with discovery and renew automatically
func (d *Discovery) Register(ins *Instance) (cancelFunc context.CancelFunc, err error) {
	d.mutex.Lock()
	if _, ok := d.registry[ins.AppID]; ok {
		err = ErrDuplication
	} else {
		d.registry[ins.AppID] = struct{}{}
	}
	d.mutex.Unlock()
	if err != nil {
		return
	}

	ctx, cancel := context.WithCancel(d.ctx)
	if err = d.register(ctx, ins); err != nil {
		d.mutex.Lock()
		delete(d.registry, ins.AppID)
		d.mutex.Unlock()
		cancel()
		return
	}
	ch := make(chan struct{}, 1)
	cancelFunc = context.CancelFunc(func() {
		cancel()
		<-ch
	})
	go func() {
		ticker := time.NewTicker(_registerGap)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := d.renew(ctx, ins); err != nil && ecode.EqualError(ecode.NothingFound, err) {
					_ = d.register(ctx, ins)
				}
			case <-ctx.Done():
				_ = d.cancel(ins)
				ch <- struct{}{}
				return
			}
		}
	}()
	return
}

// register Register an instance with discovery
func (d *Discovery) register(ctx context.Context, ins *Instance) (err error) {
	d.mutex.RLock()
	c := d.c
	d.mutex.RUnlock()

	var metadata []byte
	if ins.Metadata != nil {
		if metadata, err = json.Marshal(ins.Metadata); err != nil {
			log.Error("discovery:register instance Marshal metadata(%v) failed!error(%v)", ins.Metadata, err)
		}
	}
	res := new(struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	})
	uri := fmt.Sprintf(_registerURL, d.pickNode())
	params := d.newParams(c)
	params.Set("appid", ins.AppID)
	for _, addr := range ins.Addrs {
		params.Add("addrs", addr)
	}
	params.Set("version", ins.Version)
	params.Set("status", _statusUP)
	params.Set("metadata", string(metadata))
	if err = d.httpClient.Post(ctx, uri, "", params, &res); err != nil {
		d.switchNode()
		log.Error("discovery: register client.Get(%v)  zone(%s) env(%s) appid(%s) addrs(%v) error(%v)",
			uri, c.Zone, c.Env, ins.AppID, ins.Addrs, err)
		return
	}
	if ec := ecode.Int(res.Code); !ecode.EqualError(ecode.OK, ec) {
		log.Warn("discovery: register client.Get(%v)  env(%s) appid(%s) addrs(%v) code(%v)", uri, c.Env, ins.AppID, ins.Addrs, res.Code)
		err = ec
		return
	}
	log.Info("discovery: register client.Get(%v) env(%s) appid(%s) addrs(%s) success", uri, c.Env, ins.AppID, ins.Addrs)
	return
}

// renew Renew an instance with discovery
func (d *Discovery) renew(ctx context.Context, ins *Instance) (err error) {
	d.mutex.RLock()
	c := d.c
	d.mutex.RUnlock()

	res := new(struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	})
	uri := fmt.Sprintf(_renewURL, d.pickNode())
	params := d.newParams(c)
	params.Set("appid", ins.AppID)
	if err = d.httpClient.Post(ctx, uri, "", params, &res); err != nil {
		d.switchNode()
		log.Error("discovery: renew client.Get(%v)  env(%s) appid(%s) hostname(%s) error(%v)",
			uri, c.Env, ins.AppID, c.Host, err)
		return
	}
	if ec := ecode.Int(res.Code); !ecode.EqualError(ecode.OK, ec) {
		err = ec
		if ecode.EqualError(ecode.NothingFound, ec) {
			return
		}
		log.Error("discovery: renew client.Get(%v) env(%s) appid(%s) hostname(%s) code(%v)",
			uri, c.Env, ins.AppID, c.Host, res.Code)
		return
	}
	return
}

// cancel Remove the registered instance from discovery
func (d *Discovery) cancel(ins *Instance) (err error) {
	d.mutex.RLock()
	c := d.c
	d.mutex.RUnlock()

	res := new(struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	})
	uri := fmt.Sprintf(_cancelURL, d.pickNode())
	params := d.newParams(c)
	params.Set("appid", ins.AppID)
	// request
	if err = d.httpClient.Post(context.TODO(), uri, "", params, &res); err != nil {
		d.switchNode()
		log.Error("discovery cancel client.Get(%v) env(%s) appid(%s) hostname(%s) error(%v)",
			uri, c.Env, ins.AppID, c.Host, err)
		return
	}
	if ec := ecode.Int(res.Code); !ecode.EqualError(ecode.OK, ec) {
		log.Warn("discovery cancel client.Get(%v)  env(%s) appid(%s) hostname(%s) code(%v)",
			uri, c.Env, ins.AppID, c.Host, res.Code)
		err = ec
		return
	}
	log.Info("discovery cancel client.Get(%v)  env(%s) appid(%s) hostname(%s) success",
		uri, c.Env, ins.AppID, c.Host)
	return
}

// Set set ins status and metadata.
func (d *Discovery) Set(ins *Instance) error {
	return d.set(context.Background(), ins)
}

// set set instance info with discovery
func (d *Discovery) set(ctx context.Context, ins *Instance) (err error) {
	d.mutex.RLock()
	conf := d.c
	d.mutex.RUnlock()
	res := new(struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	})
	uri := fmt.Sprintf(_setURL, d.pickNode())
	params := d.newParams(conf)
	params.Set("appid", ins.AppID)
	params.Set("version", ins.Version)
	params.Set("status", _statusUP)
	if ins.Metadata != nil {
		var metadata []byte
		if metadata, err = json.Marshal(ins.Metadata); err != nil {
			log.Error("discovery:set instance Marshal metadata(%v) failed!error(%v)", ins.Metadata, err)
			return
		}
		params.Set("metadata", string(metadata))
	}
	if err = d.httpClient.Post(ctx, uri, "", params, &res); err != nil {
		d.switchNode()
		log.Error("discovery: set client.Get(%v)  zone(%s) env(%s) appid(%s) addrs(%v) error(%v)",
			uri, conf.Zone, conf.Env, ins.AppID, ins.Addrs, err)
		return
	}
	if ec := ecode.Int(res.Code); !ecode.EqualError(ecode.OK, ec) {
		log.Warn("discovery: set client.Get(%v)  env(%s) appid(%s) addrs(%v)  code(%v)",
			uri, conf.Env, ins.AppID, ins.Addrs, res.Code)
		err = ec
		return
	}
	log.Info("discovery: set client.Get(%v) env(%s) appid(%s) addrs(%s) success", uri+"?"+params.Encode(), conf.Env, ins.AppID, ins.Addrs)
	return
}

func (d *Discovery) serverproc() {
	var (
		retry  int
		ctx    context.Context
		cancel context.CancelFunc
	)
	ticker := time.NewTicker(time.Minute * 30)
	defer ticker.Stop()
	for {
		if ctx == nil {
			ctx, cancel = context.WithCancel(d.ctx)
			d.mutex.Lock()
			d.cancelPolls = cancel
			d.mutex.Unlock()
		}
		select {
		case <-d.ctx.Done():
			return
		case <-ticker.C:
			d.switchNode()
		default:
		}
		apps, err := d.polls(ctx)
		if err != nil {
			d.switchNode()
			if ctx.Err() == context.Canceled {
				ctx = nil
				continue
			}
			nodes, _ := d.node.Load().([]string)
			if retry >= len(nodes) {
				d.tryAppendSeedNodes()
				retry = 0
			}
			time.Sleep(time.Second)
			retry++
			continue
		}
		retry = 0
		d.broadcast(apps)
	}
}

func (d *Discovery) tryAppendSeedNodes() {
	nodes, ok := d.node.Load().([]string)
	if !ok {
		return
	}
	exist := make(map[string]struct{}, len(nodes))
	for _, node := range nodes {
		exist[node] = struct{}{}
	}
	d.mutex.Lock()
	seedNodes := d.c.Nodes
	d.mutex.Unlock()
	var changed bool
	for _, node := range seedNodes {
		if _, ok := exist[node]; !ok {
			nodes = append(nodes, node)
			changed = true
		}
	}
	if changed {
		d.node.Store(nodes)
		log.Info("discovery: append seed nodes(%s)", nodes)
	}
}

func (d *Discovery) pickNode() string {
	nodes, ok := d.node.Load().([]string)
	if !ok || len(nodes) == 0 {
		return d.c.Nodes[rand.Intn(len(d.c.Nodes))]
	}
	return nodes[atomic.LoadUint64(&d.nodeIdx)%uint64(len(nodes))]
}

func (d *Discovery) switchNode() {
	atomic.AddUint64(&d.nodeIdx, 1)
}

func (d *Discovery) polls(ctx context.Context) (apps map[string]*InstancesInfo, err error) {
	var (
		lastTss []int64
		appIDs  []string
		host    = d.pickNode()
		changed bool
	)
	if host != d.lastHost {
		d.lastHost = host
		changed = true
	}
	d.mutex.RLock()
	c := d.c
	for k, v := range d.apps {
		if changed {
			v.lastTs = 0
		}
		appIDs = append(appIDs, k)
		lastTss = append(lastTss, v.lastTs)
	}
	d.mutex.RUnlock()
	if len(appIDs) == 0 {
		return
	}
	uri := fmt.Sprintf(_pollURL, host)
	res := new(struct {
		Code int                       `json:"code"`
		Data map[string]*InstancesInfo `json:"data"`
	})
	params := url.Values{}
	params.Set("env", c.Env)
	params.Set("hostname", c.Host)
	for _, appid := range appIDs {
		params.Add("appid", appid)
	}
	for _, ts := range lastTss {
		params.Add("latest_timestamp", strconv.FormatInt(ts, 10))
	}
	if err = d.httpClient.Get(ctx, uri, "", params, res); err != nil {
		if ctx.Err() != context.Canceled {
			log.Error("discovery: client.Get(%s) error(%+v)", uri+"?"+params.Encode(), err)
		}
		return
	}
	if ec := ecode.Int(res.Code); !ecode.EqualError(ecode.OK, ec) {
		if !ecode.EqualError(ecode.NotModified, ec) {
			log.Error("discovery: client.Get(%s) get error code(%d)", uri+"?"+params.Encode(), res.Code)
			err = ec
		}
		return
	}
	info, _ := json.Marshal(res.Data)
	for _, app := range res.Data {
		if app == nil || app.LastTs == 0 {
			err = ecode.ServerErr
			log.Error("discovery: client.Get(%s) latest_timestamp is 0,instances:(%s)", uri+"?"+params.Encode(), info)
			return
		}
	}
	log.Info("discovery: successfully polls(%s) instances (%s)", uri+"?"+params.Encode(), info)
	apps = res.Data
	return
}

func (d *Discovery) broadcast(apps map[string]*InstancesInfo) {
	for appID, v := range apps {
		var count int
		for zone, ins := range v.Instances {
			if len(ins) == 0 {
				delete(v.Instances, zone)
			}
			count += len(ins)
		}
		if count == 0 {
			continue
		}
		d.mutex.RLock()
		app, ok := d.apps[appID]
		d.mutex.RUnlock()
		if ok {
			app.lastTs = v.LastTs
			app.zoneIns.Store(v)
			d.mutex.RLock()
			for rs := range app.resolver {
				select {
				case rs.event <- struct{}{}:
				default:
				}
			}
			d.mutex.RUnlock()
		}
	}
}

func (d *Discovery) newParams(c *Config) url.Values {
	params := url.Values{}
	params.Set("region", c.Region)
	params.Set("zone", c.Zone)
	params.Set("env", c.Env)
	params.Set("hostname", c.Host)
	return params
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// shuffle pseudo-randomizes the order of elements.
// n is the number of elements. Shuffle panics if n < 0.
// swap swaps the elements with indexes i and j.
func shuffle(n int, swap func(i, j int)) {
	if n < 0 {
		panic("invalid argument to Shuffle")
	}

	// Fisher-Yates shuffle: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
	// Shuffle really ought not be called with n that doesn't fit in 32 bits.
	// Not only will it take a very long time, but with 2³¹! possible permutations,
	// there's no way that any PRNG can have a big enough internal state to
	// generate even a minuscule percentage of the possible permutations.
	// Nevertheless, the right API signature accepts an int n, so handle it as best we can.
	i := n - 1
	for ; i > 1<<31-1-1; i-- {
		j := int(r.Int63n(int64(i + 1)))
		swap(i, j)
	}
	for ; i > 0; i-- {
		j := int(r.Int31n(int32(i + 1)))
		swap(i, j)
	}
}
