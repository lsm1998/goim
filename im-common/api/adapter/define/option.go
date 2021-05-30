package define

// 这里定义所有上行消息 RPCInput 的可选项。内建的名称
type goimOpt int8

const (
	// 长连接鉴权时使用
	UID    goimOpt = 1
	Token  goimOpt = 2
	AppID  goimOpt = 3
	ConnID goimOpt = 4 // 同一uid下connid唯一
	// 长连接鉴权时使用，切换房间时使用
	SubscribeRoom      goimOpt = 5
	IsAnonymousUser    goimOpt = 6
	HeartbeatThreshold goimOpt = 7
	// 连接相关信息
	ClientIP        goimOpt = 8
	ClientPort      goimOpt = 9
	AccessPointIP   goimOpt = 10
	AccessPointPort goimOpt = 11
	AccessToken     goimOpt = 12 //鉴权时需要登录token
)
