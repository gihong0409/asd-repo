package dmrsformats

//ReqHeader Request Headers
//ReqHeader is only one style

type DMRSInfo struct {
	DMRSURL string
	XMLNAME string
	APPNAME string
}

type ReqHeader struct {
	CmdType   string `json:"CmdType"`
	CallApp   string `json:"CallApp"`
	XMLName   string `json:"XMLName"`
	Query     string `json:"Query"`
	RequestID string `json:"RequestID"`
}

// ReqBody Queries RequestBody
type ReqBody struct {
	Data []interface{} `json:"Data"`
}

type ReqBody_Replace struct {
	Data    []interface{} `json:"Data"`
	Replace string        `json:"Replace"`
}

type ReqBodyReplace struct {
	Data    []interface{} `json:"Data"`
	Replace []interface{} `json:"Replace"`
}

type PushData struct {
	PushType string `json:"PushType"` // PUSHTYPE : 1000 접속 요청
	TransID  string `json:"TransID"`
}
