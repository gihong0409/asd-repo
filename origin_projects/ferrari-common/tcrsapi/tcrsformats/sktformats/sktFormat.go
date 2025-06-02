package sktformats

import "encoding/xml"

type RspSKT struct {
	ERROR_CODE string `json:"ERROR_CODE"`
	ERROR_MSG  string `json:"ERROR_MSG"`
}

type JoinSearchReq struct {
	SVC_NUM string `qstring:"SVC_NUM"`
	PROD_ID string `qstring:"PROD_ID"`
}

type RspMain struct {
	XMLName xml.Name    `xml:"RESPONSE" json:"-"`
	Header  RspHeader   `xml:"HEADER" json:"Header"`
	Body    interface{} `xml:"BODY" json:"Body"`
}

type RspMainHub struct {
	RESPONSE struct {
		Header RspHeaderHub `json:"HEADER"`
		Body   interface{}  `json:"BODY"`
	} `json:"RESPONSE"`
}

type RspHeader struct {
	Result     string `xml:"RESULT" json:"Result"`
	ReturnCode string `xml:"RESULT_CODE" json:"ReturnCode"`
	ReturnDesc string `xml:"RESULT_MESSAGE" json:"ReturnDesc"`
}

type RspHeaderHub struct {
	Result       string `xml:"RESULT" json:"RESULT"`
	ReturnCode   string `xml:"RESULT_CODE" json:"RESULT_CODE"`
	ReturnDesc   string `xml:"RESULT_MESSAGE"  json:"RESULT_MESSAGE"`
	ResponseCode string `xml:"RESPONSE_CODE"  json:"RESPONSE_CODE"`
}

type JoinSearchRsp struct {
	JOIN_YN string `xml:"JOIN_YN" json:"JOIN_YN"`
	PROD_ID string `xml:"PROD_ID" json:"PROD_ID"`
	JOIN_DT string `xml:"JOIN_DT" json:"JOIN_DT"`
}

type JoinDelReq struct {
	OP_CD   string `qstring:"OP_CD" json:"OP_CD"`
	SVC_NUM string `qstring:"SVC_NUM" json:"SVC_NUM"`
	PROD_ID string `qstring:"PROD_ID" json:"PROD_ID"`
}

type SVCMgmtNumReq struct {
	SVC_MGMT_NUM string `xml:"SVC_MGMT_NUM" json:"SVC_MGMT_NUM"`
}

type JoinDelRsp struct {
	SVC_MGMT_NUM string `xml:"SVC_MGMT_NUM" json:"SVC_MGMT_NUM"`
}

type JoinDelRspHub struct {
	ADDED_SVC_RESULT_DATA JoinDelRsp `xml:"ADDED_SVC_RESULT_DATA" json:"ADDED_SVC_RESULT_DATA"`
}

type UserInfoRsp struct {
	SSN_BIRTH_DT string `xml:"SSN_BIRTH_DT" json:"SSN_BIRTH_DT"`
	//SSN_SEX_CD   string `xml:"SSN_SEX_CD" json:"SSN_SEX_CD"`
	//CUST_TYP_CD  string `xml:"CUST_TYP_CD" json:"CUST_TYP_CD"`
	EQP_MDL_CD          string `xml:"EQP_MDL_CD" json:"EQP_MDL_CD"`
	SVC_ST_CD           string `xml:"SVC_ST_CD" json:"SVC_ST_CD"`
	SVC_CHG_RSN_CD      string `xml:"SVC_CHG_RSN_CD" json:"SVC_CHG_RSN_CD"`
	IMSI_NUM_N_MAC_ADDR string `xml:"IMSI_NUM_N_MAC_ADDR" json:"IMSI_NUM_N_MAC_ADDR"`
	EQP_SER_NUM         string `xml:"EQP_SER_NUM" json:"EQP_SER_NUM"`
	IMEI_NUM            string `xml:"IMEI_NUM" json:"IMEI_NUM"`
	USIM_MDL_CD         string `xml:"USIM_MDL_CD" json:"USIM_MDL_CD"`
	USIM_SER_NUM        string `xml:"USIM_SER_NUM" json:"USIM_SER_NUM"`
	SVC_MGMT_NUM        string `xml:"SVC_MGMT_NUM" json:"SVC_MGMT_NUM"`
}

type MVNORsp struct {
	CUST_TYP_CD string `xml:"CUST_TYP_CD" json:"CUST_TYP_CD"`
	EQP_MDL_CD  string `xml:"EQP_MDL_CD" json:"EQP_MDL_CD"`
}

type SVCNumReq struct {
	SVC_NUM string `qstring:"SVC_NUM" json:"SVC_NUM"`
}

type LostPhoneCheckReq struct {
	EQP_MDL_CD   string `json:"EQP_MDL_CD"`   //단말기 모델코드
	EQP_SER_NUM  string `json:"EQP_SER_NUM"`  //단말기 일련번호
	SVC_MGMT_NUM string `json:"SVC_MGMT_NUM"` //서비스 관리코드
}

type LostPhoneCheckRSP struct {
	EqpLossInfoList EQP_LOSS_INQ_LIST `xml:"EQP_LOSS_INQ_LIST" json:"EQP_LOSS_INQ_LIST"`
}

type EQP_LOSS_INQ_LIST struct {
	EqpLossInfo EQP_LOSS_INQ `xml:"EQP_LOSS_INQ" json:"EQP_LOSS_INQ"`
}

type EQP_LOSS_INQ struct {
	LOSS_YN string `xml:"LOSS_YN" json:"LOSS_YN"` //분실여부
	RCV_DT  string `xml:"RCV_DT" json:"RCV_DT"`   //분실등록일자
}

type LostPhoneCheckRSP_Hub struct {
	EqpLossInfoList EQP_LOSS_INQ_LIST_Hub `xml:"EQP_LOSS_INQ_LIST" json:"EQP_LOSS_INQ_LIST"`
}

type EQP_LOSS_INQ_LIST_Hub struct {
	EqpLossInfo []EQP_LOSS_INQ `xml:"EQP_LOSS_INQ" json:"EQP_LOSS_INQ"`
}

type DeviceInfoRSP struct {
	USG_EQP_MDL_CD  string `xml:"USG_EQP_MDL_CD" json:"USG_EQP_MDL_CD"`   //현소유자 단말기 코드
	USG_EQP_MDL_NM  string `xml:"USG_EQP_MDL_NM" json:"USG_EQP_MDL_NM"`   //현소유자 단말기 모델명
	USG_EQP_SER_NUM string `xml:"USG_EQP_SER_NUM" json:"USG_EQP_SER_NUM"` //현소유자 단말기 일련번호
	SVC_EQP_ST_CD   string `xml:"SVC_EQP_ST_CD" json:"SVC_EQP_ST_CD"`     //서비스 단말기 상태코드
}

type SMSReq struct {
	CONSUMER_ID   string `qstring:"CONSUMER_ID"`
	RPLY_PHON_NUM string `qstring:"RPLY_PHON_NUM"`
	TITLE         string `qstring:"TITLE"`
	PHONE         string ``
}

type SMSRsp struct {
	RETURN string `json:"RETURN"`
}

type RSPMMSHeader struct {
	TransactionID string `json:"TransactionID"`
}

type RSPMMSBodyStatus struct {
	StatusCode string `json:"StatusCode"`
	StatusText string `json:"StatusText"`
}

type RSPMMSBodySubmitRsp struct {
	Status     RSPMMSBodyStatus `json:"Status"`
	MM7Version string           `json:"MM7Version"`
}

type RSPMMSBody struct {
	SubmitRsp *RSPMMSBodySubmitRsp `json:"SubmitRsp"`
}

type RSPMMS struct {
	Header RSPMMSHeader `json:"Header"`
	Body   RSPMMSBody   `json:"Body"`
}

type LongtimeNotUseTrID struct {
	TrID           string `json:"TR_ID"`
	ResultCode     string `json:"RESULT_CODE"`
	ResultMsg      string `json:"RESULT_STATUS"`
	SrcResultCode  string `json:"SRC_RESULT_CODE"`
	ChargingMethod string `json:"CHARGING_METHOD"`
	ChargeAmount   string `json:"CHARGE_AMOUNT"`
	Balance        string `json:"BALANCE"`
}

type LongtimeNotUseCharging struct {
	TrID           string `json:"TR_ID"`
	MsgVersion     string `json:"MSG_VERSION"`
	ProcCode       int    `json:"PROC_CODE"`
	PtID           string `json:"PT_ID"`
	SystemName     string `json:"SYSTEM_NAME"`
	DcmfPID        string `json:"DCMF_PID"`
	SystemDivision string `json:"SYSTEM_DIVISION"`
	ChargingID     string `json:"CHARGING_ID"`
	ChargingIND    int    `json:"CHARGING_IND"`
	ChargingMethod string `json:"CHARGING_METHOD"`
	ChargingExcept int    `json:"CHARGING_EXCEPT"`
	CallingID      string `json:"CALLING_ID"`
	CallingIND     int    `json:"CALLING_IND"`
	ChargePivot    int    `json:"CHARGE_PIVOT"`
	ChargeAmount   string `json:"CHARGE_AMOUNT"`
	URL1           string `json:"URL1"`
	CDS            int    `json:"C_D_S"`
	DataSize       int    `json:"DATA_SIZE"`
	SvcOperation   string `json:"SVC_OPERATION"`
	UaFlag         string `json:"UA_FLAG"`
	DeviceIP       string `json:"DEVICE_IP"`
	Protocal       string `json:"PROTOCOL"`
	BpStatus       string `json:"BP_STATUS"`
}

type UseServiceReq struct {
	PNumber string `xml:"PNumber" json:"PNumber"`
}
