package tcrsformats

type ReqHeader struct {
	CmdType     string `json:"CmdType"`
	RequestID   string `json:"RequestID"`
	CallAppName string `json:"CallAppName"`
}

type ReqBodySMS struct {
	MsgType        int    `json:"MsgType"`
	RecvPNumber    string `json:"RecvPNumber"`
	Title          string `json:"Title"`
	SendMsg        string `json:"SendMsg"`
	SendFileBase64 string `json:"SendFileBase64"`
}

type ReqBodyMultiMMS struct {
	MsgType        int      `json:"MsgType"`
	RecvPNumber    []string `json:"RecvPNumber"`
	Title          string   `json:"Title"`
	SendMsg        string   `json:"SendMsg"`
	SendFileBase64 string   `json:"SendFileBase64"`
}

type ReqBodyPNumber struct {
	PNumber string `json:"PNumber"`
}

type ReqBodyMGMTNUM struct {
	SVC_MGMT_NUM string `json:"SVC_MGMT_NUM"`
}

type ReqBodyLostPhoneCheck struct {
	EQP_MDL_CD   string `json:"EQP_MDL_CD"`   //단말기모델코드
	EQP_SER_NUM  string `json:"EQP_SER_NUM"`  //단말기일련번호
	SVC_MGMT_NUM string `json:"SVC_MGMT_NUM"` //서비스관리번호
}

type ReqBodyTerms struct {
	PNumber string `json:"PNumber"`
	Terms1  string `json:"Terms1"`
	Terms2  string `json:"Terms2"`
	Terms3  string `json:"Terms3"`
	Terms4  string `json:"Terms4"`
	Terms5  string `json:"Terms5"`
}

// type ReqBodyLMS struct {
// 	MsgType     int    `json:"MsgType"`
// 	RecvPNumber string `json:"RecvPNumber"`
// 	SendMsg     string `json:"SendMsg"`
// 	Title       string `json:"Title"`
// }

type ReqBodyPush struct {
	RecvPushToken string `json:"RecvPushToken"`
	SendMsg       string `json:"SendMsg"`
	Title         string `json:"Title"`
	PushData      string `json:"PushData"`
}

type ReqBodyMultiPush struct {
	RecvPushToken []string `json:"RecvPushToken"`
	SendMsg       string   `json:"SendMsg"`
	Title         string   `json:"Title"`
	PushType      string   `json:"PushType"`
	TransID       string   `json:"TransID"`
	EventID       string   `json:"EventID"`
	ReqPNumber    string   `json:"ReqPNumber"`
	ResPNumber    string   `json:"ResPNumber"`
}

// type ReqBodyMMS struct {
// 	MsgType     int    `json:"MsgType"`
// 	RecvPNumber string `json:"RecvPNumber"`
// 	SendMsg     string `json:"SendMsg"`
// 	Title       string `json:"Title"`
// }

type ReqTelegramSendBody struct {
	BotToken    string `json:"BotToken"`
	ChatID      string `json:"ChatID"`
	SendMessage string `json:"SendMessage"`
}

type RspHeader struct {
	CmdType string `json:"CmdType"`
	ErrCode int    `json:"ErrCode"`
	ErrMsg  string `json:"ErrMsg"`
}

type RspSearchTelecomBody struct {
	Telecom     int    `json:"Telecom"`
	TelecomName string `json:"TelecomName"`
}

type DURetHeader struct {
	ErrCode int    `json:"ErrCode"`
	ErrMsg  string `json:"ErrMsg"`
}

func (_this *RspHeader) SetMsg(nMsgCode int) {
	_this.ErrCode = nMsgCode
	_this.ErrMsg = GetMsg(nMsgCode)
}

type Model struct {
	Model string `json:"Model"`
}
