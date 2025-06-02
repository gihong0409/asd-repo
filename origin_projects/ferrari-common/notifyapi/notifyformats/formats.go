package notifyformats

type ReqHeader struct {
	CmdType int `json:"CmdType"`
}

type ReqBodySMCommon struct {
	//서비스명 ex) 스마트피싱보호
	ServiceName string `json:"ServiceName"`
	//서버명 ex) bentley-cbd-kt
	ServerName string `json:"ServerName"`
	//플랫폼 ex) kt, skt, lgup ...
	Platform string `json:"Platform"`
	//메시지 내용 ex) [KT 해지자 확인 시작]대상 425341건
	Msg string `json:"Msg"`
}

type ReqBodyWebhookNotification struct {
	ReqBodySMCommon
}

type ReqBodyDBInsert struct {
	//0: 배치 시작 | 1: 배치 완료 | 2: 단순 알림용 배치
	BatchType int `json:"BatchType"`
	ReqBodySMCommon
}

type QueryBody struct {
	Data []interface{} `json:"Data"`
}

type RspHeader struct {
	CmdType int    `json:"CmdType"`
	ErrCode int    `json:"ErrCode"`
	ErrMsg  string `json:"ErrMsg"`
}

type ReqBodySlackMsgByPass struct {
	Msg         string        `json:"Msg"`
	ChannelID   string        `json:"ChannelID"`
	Attachments []Attachments `json:"Attachments"`
}

type Attachments struct {
	Color  string   `json:"color"`
	Blocks []Blocks `json:"blocks"`
}

type Blocks struct {
	Type string `json:"type"`
	Text *Text  `json:"text,omitempty"`
}

type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type NotifySlackParams struct {
	CmdType      int
	BatchType    int // CmdType 1000 일 때만 필요
	ServiceName  string
	ServerName   string
	PlatformName string
	Msg          string
	ChannelID    string        // CmdType 4000 일 때만 필요
	Attachments  []Attachments // CmdType 4000 일 때 Optional
	ConfigSet    string        // "DEV" or "LIVE"
}
