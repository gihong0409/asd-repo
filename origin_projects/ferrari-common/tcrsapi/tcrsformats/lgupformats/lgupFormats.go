package lgupformats

type RspLGUP struct {
	ErrCode int    `json:"ErrCode"`
	ErrMsg  string `json:"ErrMsg"`
}

type RspLGUPDURL struct {
	RetURL string `json:"RetURL"`
}

func (_this *RspLGUP) SetMsg(nMsgCode int) {
	_this.ErrCode = nMsgCode
	_this.ErrMsg = "" //formats.GetMsg(nMsgCode)
}

type ReqBodyNCASPNumber struct {
	PNumber   string `json:"PNumber"`
	OTPNumber string `json:"OTPNumber"`
}

type RSPUserInfo struct {
	RESPCODE      int    `json:"RESPCODE"`
	RESPMSG       string `json:"RESPMSG"`
	//Age           string `json:"AGE_OUT"`
	CTN_STUS_CODE string `json:"CTN_STUS_CODE"` //CTN 상태 코드(A:정상/S:일시 중지)
	//PRE_PAY_CODE  string `json:"PRE_PAY_CODE"`
	//REF_TYPE_CODE string `json:"REF_TYPE_CODE"`
	MDL_VALUE          string `json:"MDL_VALUE"`    //단말속성정보
	UNIT_MDL           string `json:"UNIT_MDL"`     //단말기명
	YOUNG_FEE_YN       string `json:"YOUNG_FEE_YN"` //청소년 정보료 상한요금제
	SVC_AUTH_DT        string `json:"SVC_AUTH_DT"`
	UNIT_LOSS_YN_CODE  string `json:"UNIT_LOSS_YN_CODE"` //분실여부
	REAL_BIRTH_PERS_ID string `json:"REAL_BIRTH_PERS_ID,omitempty"`
	SUB_BIRTH_PERS_ID  string `json:"SUB_BIRTH_PERS_ID,omitempty"`
}

type RSPBodyHeader struct {
	ErrCode int    `json:"ErrCode"`
	ErrMsg  string `json:"ErrMsg`
	CmdType string `json:"CmdType`
}

type RspLgupSMS struct {
	RetSMSBody    RetSMSBody    `json:"Body"`
	RSPBodyHeader RSPBodyHeader `json:"Header"`
}

//RESPCODE=00  &RESPMSG=NsIZG867RHU=  &CTN=5+SK0+KXltq8d5UehSNYtg==  &SUB_NO=s7g9MYiyvA8r/2TqrTjHvA==
type IFDATARsp struct {
	RESPCODE int    `qstring:"RESPCODE" json:"RESPCODE"`
	RESPMSG  string `qstring:"RESPMSG" json:"RESPMSG"`
	//AGE_OUT        string `qstring:"AGE_OUT" json:"AGE_OUT"`
	CTN_STUS_CODE string `qstring:"CTN_STUS_CODE" json:"CTN_STUS_CODE"` //(USERINFO 대체)CTN 상태 코드(A:정상/S:일시 중지)
	//CUST_TYPE_CODE string `qstring:"CUST_TYPE_CODE" json:"CUST_TYPE_CODE"`
	//PRE_PAY_CODE   string `qstring:"PRE_PAY_CODE" json:"PRE_PAY_CODE"`
	//REF_TYPE_CODE  string `qstring:"REF_TYPE_CODE" json:"REF_TYPE_CODE"`
	UNIT_MDL string `qstring:"UNIT_MDL" json:"UNIT_MDL"` //단말기명
	//SVC_AUTH          string `qstring:"SVC_AUTH" json:"SVC_AUTH"`
	SVC_AUTH_DT        string `qstring:"SVC_AUTH_DT" json:"SVC_AUTH_DT"`
	MDL_VALUE          string `qstring:"MDL_VALUE" json:"MDL_VALUE"`                 //단말속성정보
	YOUNG_FEE_YN       string `qstring:"YOUNG_FEE_YN" json:"YOUNG_FEE_YN"`           //청소년 정보료 상한요금제
	UNIT_LOSS_YN_CODE  string `qstring:"UNIT_LOSS_YN_CODE" json:"UNIT_LOSS_YN_CODE"` //(USERINFO 대체)분실여부
	REAL_BIRTH_PERS_ID string `qstring:"REAL_BIRTH_PERS_ID" json:"REAL_BIRTH_PERS_ID,omitempty"`
	SUB_BIRTH_PERS_ID  string `qstring:"SUB_BIRTH_PERS_ID" json:"SUB_BIRTH_PERS_ID,omitempty"`
}

// AGE_OUT : 만 나이 일반고객일경우 "만 나이" 법인일경우 "G"
// CTN_STUS_CODE : (USERINFO 대체)CTN 상태 코드(A:정상/S:일시 중지)
// CUST_TYPE_CODE : (USERINFO 대체)개인,법인구분(I:개인/G:법인)
// PRE_PAY_CODE : (USERINFO 대체)선불가입여부(선불가입자:PPS)
// REF_TYPE_CODE : MVNO코드 + 별정코드
// UNIT_MDL : 단말기명

type OTPReq struct {
	Url     string `qstring:"url"`
	SOC     string `qstring:"soc"`
	OPID    string `qstring:"p_code"`
	PCODE   string `qstring:"op_id"`
	PNumber string `qstring:"ctn"`
}

type IFDATAReq struct {
	CPTYPE   string
	CPID     string
	CPPWD    string
	CASECODE string
	CTN      string
}

type CASReq struct {
	ReqType string `qstring:"req_type"`
	Url     string `qstring:"url"`
	SOC     string `qstring:"soc"`
	PNumber string `qstring:"ctn"`
	PCODE   string `qstring:"op_id"`
	OPID    string `qstring:"p_code"`
	SiteUrl string `qstring:"siteURL"`
	CmpnyId string `qstring:"cmpnyId"`
}

type CASDCReq struct {
	ReqType         string `qstring:"REQ_TYPE"`
	URL             string `qstring:"RETURN_URL"`
	SOC             string `qstring:"SOC"`
	PNumber         string `qstring:"CTN"`
	DiscountCode    string `qstring:"DC_CD"`
	DiscountStartDT string `qstring:"DC_STR_DT"`
	DiscountEndDT   string `qstring:"DC_END_DT"`
	SourceIP        string `qstring:"SVC_IP"`
}

type RetSMSBody struct {
	TransactionID string `json:"TransactionID"`
	StatusCode    string `json:"StatusCode"`
	StatusText    string `json:"StatusText"`
	MM7Version    string `json:"MM7Version"`
	MessageID     string `json:"MessageID"`
}
