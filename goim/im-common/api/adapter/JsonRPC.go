package adapter

type JsonInput struct {
	ServerName string           `json:"ServerName"`
	Func       string           `json:"Func"`
	Req        string           `json:"Req"`
	Opt        map[int32]string `json:"Opt"`
}

type JsonOutput struct {
	Ret        int32            `json:"Ret"`
	Rsp        interface{}      `json:"Rsp"`
	Opt        map[int32]string `json:"Opt"`
	Desc       string           `json:"Desc"`
	ServerName string           `json:"ServerName"`
	Func       string           `json:"Func"`
}
