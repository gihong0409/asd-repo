package fcmformats

import "gopkg.in/maddevsio/fcm.v1"

type FcmResults struct {
	OK      bool     `json:"OK"`
	Fail    int      `json:"Fail"`
	Success int      `json:"Success"`
	Results []Result `json:"Results"`
}

type FcmMultiResults struct {
	OK      bool         `json:"OK"`
	Fail    int          `json:"Fail"`
	Success int          `json:"Success"`
	Results []fcm.Result `json:"Results"`
}

type Result struct {
	MessageID      string `json:"MessageID"`
	RegistrationID string `json:"RegistrationID"`
	Error          string `json:"Error"`
}

type RspFCM struct {
	ERROR_CODE string      `json:"ERROR_CODE"`
	ERROR_MSG  string      `json:"ERROR_MSG"`
	FCM_RSP    interface{} `json:"FCM_RSP"`
}

type RspMultiFCM struct {
	ERROR_CODE string          `json:"ERROR_CODE"`
	ERROR_MSG  string          `json:"ERROR_MSG"`
	FCM_RSP    FcmMultiResults `json:"FCM_RSP"`
}

/*
PushType	1000: 서버 접속 요청, 9999 : 해지 - Noti 생성 후 앱 종료
ConnectIP	접속 IP
ConnectPort 접속 PORt
*/
type PushData struct {
	PushType string `json:"PushType"`
	TransID  string `json:"TransID"`
}
