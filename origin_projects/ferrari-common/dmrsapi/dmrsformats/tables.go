package dmrsformats

// ResHeader nomal Response Body
type RspHeader struct {
	CmdType   string `json:"CmdType"`
	ErrCode   int    `json:"ErrCode"`
	ErrMsg    string `json:"ErrMsg"`
	CallApp   string `json:"CallApp"`
	XMLName   string `json:"XMLName"`
	Query     string `json:"Query"`
	RequestID string `json:"RequestID"`
}

type ResExecQueris struct {
	LastInsertId int64
	RowsAffected int64
}

type SvcCaptcha struct {
	TransID             string `json:"TransID"`
	CryptoNumber        string `json:"CryptoNumber"`
	CryptoVerifyComplet int    `json:"CryptoVerifyComplet"`
	CryptoVerifyCnt     int    `json:"CryptoVerifyCnt"`
	RegDT               string `json:"RegDT"`
}

type SvcPartnerList struct {
	PSKey           string `json:"PSKey"`
	PCODE           string `json:"PCODE"`
	PGroup          string `json:"PGroup"`
	PName           string `json:"PName"`
	JoinURL         string `json:"JoinURL"`
	Image1          string `json:"Image1"`
	Image2          string `json:"Image2"`
	FontSize        string `json:"FontSize"`
	FontColor       string `json:"FontColor"`
	NotiContent     string `json:"NotiContent"`
	ButtonMsgBefore string `json:"ButtonMsgBefore"`
	ButtonMsgAfter  string `json:"ButtonMsgAfter"`
	DelayRewardMin  int    `json:"DelayRewardMin"`
	UiData          string `json:"UiData"`
	SKTJoin         string `json:"SKTJoin"`
	KTJoin          string `json:"KTJoin"`
	LGUPJoin        string `json:"LGUPJoin"`
	JoinBlockMsg    string `json:"JoinBlockMsg"`
}

type SvcJoinMsg struct {
	Title     string `json:"Title"`
	Content   string `json:"Content"`
	MsgType   int    `json:"MsgType"`
	DelaySec  int    `json:"DelaySec"`
	ImagePath string `json:"ImagePath"`
	Telecom   int    `json:"Telecom"`
}

// ResBody nomal Response Body
type SvcNotices struct {
	Idx          int    `json:"Idx"`
	Title        string `json:"Title"`
	Contents     string `json:"Contents"`
	ContentsType int    `json:"ContentsType"`
	EventType    int    `json:"EventType"`
	StartTime    string `json:"StartTime"`
	EndTime      string `json:"EndTime"`
	RegDT        string `json:"RegDT"`
	Total        string `json:"Total"`
}

// ResBody Member Response Body
type SvcMember struct {
	Telecom            int    `json:"Telecom"`
	PNumber            string `json:"PNumber"`
	PCode              string `json:"PCode"`
	PushToken          string `json:"PushToken"`
	ReqApp             string `json:"ReqApp"`
	ReqAppUnreg        string `json:"ReqAppUnreg"`
	Model              string `json:"Model"`
	DeviceStatus       int    `json:"DeviceStatus"`
	DeviceStatusMsg    string `json:"DeviceStatusMsg"`
	DeviceSerialNumber string `json:"DeviceSerialNumber"`
	LastLoginDT        string `json:"LastLoginDT"`
	UnRegType          int    `json:"UnRegType"`
	UnRegMsg           string `json:"UnRegMsg"`
	MemberType         int    `json:"MemberType"`
	PhoneType          string `json:"PhoneType"`
	Terms              int    `json:"Terms"`
	Pin                string `json:"Pin"`
	PinErrCnt          int    `json:"PinErrCnt"`
	PinWebErrCnt       int    `json:"PinWebErrCnt"`
	LoginErrCnt        int    `json:"LoginErrCnt"`
	Email              string `json:"Email"`
	Version            string `json:"Version"`
	Memo               string `json:"Memo"`
	GPNumber           string `json:"GPNumber"`
	RegDT              string `json:"RegDT"`
}

type SvcSpMember struct {
	Telecom int    `json:"Telecom"`
	PNumber string `json:"PNumber"`
	SpType  int    `json:"SpType"`
	RegType string `json:"RegType"`
	Memo    string `json:"memo"`
	RegDT   string `json:"RegDT"`
	RegID   string `json:"RegID"`
}

type SvcFreemembers struct {
	Idx       int    `json:"Idx"`
	PNumber   string `json:"PNumber"`
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
	Contents  string `json:"Contents"`
	RegDT     string `json:"RegDT"`
}

// ResBody Member Response Body
type SvcAuthMsg struct {
	Idx            int    `json:"Idx"`
	PCode          string `json:"PCode"`
	Telecom        int    `json:"Telecom"`
	PNumber        string `json:"PNumber"`
	AuthType       string `json:"AuthType"`
	Title          string `json:"Title"`
	Msg            string `json:"Msg"`
	TransID        string `json:"TransID"`
	ApprovalNumber int    `json:"ApprovalNumber"`
	Result         int    `json:"Result"`
	ErrCnt         int    `json:"ErrCnt"`
}

// ResBody Member Response Body
type SvcTerms struct {
	Idx           int    `json:"Idx"`
	TermsType     int    `json:"TermsType"`
	TermsTitle1   string `json:"TermsTitle1"`
	Terms1        string `json:"Terms1"`
	TermsTitle2   string `json:"TermsTitle2"`
	Terms2        string `json:"Terms2"`
	TermsTitle3   string `json:"TermsTitle3"`
	Terms3        string `json:"Terms3"`
	TermsTitle4   string `json:"TermsTitle4"`
	Terms4        string `json:"Terms4"`
	TermsTitle5   string `json:"TermsTitle5"`
	Terms5        string `json:"Terms5"`
	TermsTitle6   string `json:"TermsTitle6"`
	Terms6        string `json:"Terms6"`
	TermsTitle7   string `json:"TermsTitle7"`
	Terms7        string `json:"Terms7"`
	TermsTitle8   string `json:"TermsTitle8"`
	Terms8        string `json:"Terms8"`
	TermsTitle9   string `json:"TermsTitle9"`
	Terms9        string `json:"Terms9"`
	Order         string `json:"Order"`
	EssentialType string `json:"EssentialType"`
}

type SvcTermsQuestion struct {
	TermsTitle string `json:"TermsTitle"`
	Terms      string `json:"Terms"`
}

type SvcTransMsg struct {
	Idx        int    `json:"Idx"`
	MsgType    int    `json:"MsgType"`
	PCode      string `json:"PCode"`
	TransID    string `json:"TransID"`
	Telecom    int    `json:"Telecom"`
	PNumber    string `json:"PNumber"`
	PushToken  string `json:"PushToken"`
	PushData   string `json:"PushData"`
	Title      string `json:"Title"`
	Msg        string `json:"Msg"`
	Complete   int    `json:"Complete"`
	MsgIdx     int    `json:"MsgIdx"`
	RetCode    string `json:"RetCode"`
	RetMsg     string `json:"RetMsg"`
	SessionID  string `json:"SessionID"`
	ImagePath1 string `json:"ImagePath1"`
	ImagePath2 string `json:"ImagePath2"`
	ImagePath3 string `json:"ImagePath3"`
}

type SvcPhoneBattery struct {
	PNumber string `json:"PNumber"`
	Status  string `json:"Status"`
	Percent string `json:"Percent"`
	RegDT   string `json:"RegDTString"`
}

type SvcPhoneGPS struct {
	PNumber     string `json:"PNumber"`
	Latitude    string `json:"Latitude"`
	Longitude   string `json:"Longitude"`
	MobileRegDT string `json:"MobileRegDTString"`
}

type SvcPhoneUSIM struct {
	PNumber      string `json:"PNumber"`
	NewPNumber   string `json:"NewPNumber"`
	IMSI         string `json:"IMSI"`
	ICCID        string `json:"ICCID"`
	OperatorName string `json:"OperatorName"`
	OperatorCode string `json:"OperatorCode"`
	IMEI1        string `json:"IMEI1"`
	IMEI2        string `json:"IMEI2"`
	RegDT        string `json:"RegDTString"`
}

type SvcPhoneWIFI struct {
	PNumber   string `json:"PNumber"`
	SSID      string `json:"SSID"`
	Strength  string `json:"Strength"`
	Secret    string `json:"Secret"`
	Frequency string `json:"Frequency"`
	RegDT     string `json:"RegDTString"`
}

type SvcTelecomJoin struct {
	Idx        int    `json:"Idx"`
	Telecom    int    `json:"Telecom"`
	PNumber    string `json:"PNumber"`
	JobType    int    `json:"JobType"`
	PCode      string `json:"PCode"`
	MemberType int    `json:"MemberType"`
	UnRegType  int    `json:"UnRegType"`
	Complete   int    `json:"Complete"`
	PushToken  string `json:"PushToken"`
	RetCode    string `json:"RetCode"`
	RetMsg     string `json:"RetMsg"`
	RegDT      string `json:"RegDT"`
	SendDT     string `json:"SendDT"`
	RecvDT     string `json:"RecvDT"`
}

type SvcSnapShot struct {
	Idx           int    `json:"Idx"`
	PNumber       string `json:"PNumber"`
	LocalFilePath string `json:"LocalFilePath"`
	FileName      string `json:"FileName"`
	FileSize      int    `json:"FileSize"`
	RegDTString   string `json:"RegDTString"`
}

type SvcLockStatus struct {
	Idx             int    `json:"Idx"`
	PNumber         string `json:"PNumber"`
	DeviceStatus    int    `json:"DeviceStatus"`
	DeviceStatusMsg string `json:"DeviceStatusMsg"`
	RegDTString     string `json:"RegDTString"`
}

type SvcFAQ struct {
	Idx           int    `json:"Idx"`
	ContentsType  int    `json:"ContentsType"`
	ContentsGroup string `json:"ContentsGroup"`
	Title         string `json:"Title"`
	Contents      string `json:"Contents"`
	RegDT         string `json:"RegDT"`
}

type CbaMember struct {
	Telecom    int    `json:"Telecom"`
	PNumber    string `json:"PNumber"`
	PushToken  string `json:"PushToken"`
	CompleteDT string `json:"CompleteDT"`
	Complete   int    `json:"Complete"`
}

type CountMember struct {
	CountMember int `json:"CountMember"`
}

type CountCbaCode struct {
	CountCbaCode int `json:"CountCbaCode"`
}

type AdmCSMember struct {
	UserID      string `json:"UserID"`
	Password    string `json:"Password"`
	UserName    string `json:"UserName"`
	GroupName   string `json:"GroupName"`
	PNumber     string `json:"PNumber"`
	AuthType    string `json:"AuthType"`
	StatAuth    string `json:"StatAuth"`
	Email       string `json:"Email"`
	ShowList    string `json:"ShowList"`
	RegDT       string `json:"RegDT"`
	LastLoginDT string `json:"LastLoginDT"`
	LoginErrCnt int    `json:"LoginErrCnt"`
}

type SvcPauseMember struct {
	Idx         int    `json:"Idx"`
	Telecom     int    `json:"Telecom"`
	PNumber     string `json:"PNumber"`
	Pin         string `json:"Pin"`
	RegDT       string `json:"RegDT"`
	CbaCode     string `json:"CbaCode"`
	Des         string `json:"Des"`
	AppName     string `json:"AppName"`
	LastCheckDT string `json:"LastCheckDT"`
	Status      int    `json:"Status"`
}

type SvcMemberTerms struct {
	Idx         int    `json:"Idx"`
	TermsID     int    `json:"TermsID"`
	PNumber     string `json:"PNumber"`
	Terms1Agree int    `json:"Terms1Agree"`
	Terms2Agree int    `json:"Terms2Agree"`
	Terms3Agree int    `json:"Terms3Agree"`
	Terms4Agree int    `json:"Terms4Agree"`
	Terms5Agree int    `json:"Terms5Agree"`
	Terms6Agree int    `json:"Terms6Agree"`
	Terms7Agree int    `json:"Terms7Agree"`
	Terms8Agree int    `json:"Terms8Agree"`
	Terms9Agree int    `json:"Terms9Agree"`
	RegDT       string `json:"RegDT"`
}

type SvcRewardMember struct {
	Idx        int    `json:"Idx"`
	Telecom    int    `json:"Telecom"`
	PNumber    string `json:"PNumber"`
	AuthKey    string `json:"AuthKey"`
	PCODE      string `json:"PCODE"`
	GetFlag    int    `json:"GetFlag"`
	RetCode    string `json:"RetCode"`
	CompleteDT string `json:"CompleteDT"`
	RegDT      string `json:"RegDT"`
}

type RewardMembers struct {
	Idx            int    `json:"Idx"`
	Telecom        int    `json:"Telecom"`
	PNumber        string `json:"PNumber"`
	AuthKey        string `json:"AuthKey"`
	PCODE          string `json:"PCODE"`
	GetFlag        int    `json:"GetFlag"`
	RetCode        string `json:"RetCode"`
	CompleteDT     string `json:"CompleteDT"`
	RegDT          string `json:"RegDT"`
	RewardAPI      string `json:"RewardAPI"`
	RewardURI      string `json:"RewardURI"`
	DelayRewardMin int    `json:"DelayRewardMin"`
}

type SvcUseTransMsg struct {
	Idx      int    `json:"Idx"`
	PCode    string `json:"PCode"`
	Telecom  int    `json:"Telecom"`
	PNumber  string `json:"PNumber"`
	MsgIdx   int    `json:"MsgIdx"`
	Complete int    `json:"Complete"`
	RetCode  string `json:"RetCode"`
	RetMsg   string `json:"RetMsg"`
}

type SvcUseTransMsgContent struct {
	Idx     int    `json:"Idx"`
	Title   string `json:"Title"`
	Msg     string `json:"Msg"`
	RegDT   string `json:"RegDT"`
	MsgType int    `json:"MsgType"`
	UseFlag int    `json:"UseFlag"`
}

type SvcUseMember struct {
	Idx        int    `json:"Idx"`
	Telecom    int    `json:"Telecom"`
	PNumber    string `json:"PNumber"`
	Complete   int    `json:"Complete"`
	RegApp     string `json:"RegApp"`
	RegAction  string `json:"RegAction"`
	RegDT      string `json:"RegDT"`
	SendDT     string `json:"SendDT"`
	RecvDT     string `json:"RecvDT"`
	Trid       string `json:"Trid"`
	ResultCode string `json:"ResultCode"`
}

type EventMembersV2 struct {
	Idx        int    `json:"Idx"`
	Telecom    int    `json:"Telecom"`
	PNumber    string `json:"PNumber"`
	PCODE      string `json:"PCODE"`
	CouponCode string `json:"CouponCode"`
	GetFlag    int    `json:"GetFlag"`
	RetCode    string `json:"RetCode"`
	RetMsg     string `json:"RetMsg"`
	CompleteDT string `json:"CompleteDT"`
	PCodeGroup int    `json:"PCodeGroup"`
	DelaySec   int    `json:"DelaySec"`
	Title      string `json:"Title"`
	Msg        string `json:"Msg"`
	FailMsg    string `json:"FailMsg"`
}

type EventMembers struct {
	Idx        int    `json:"Idx"`
	Telecom    int    `json:"Telecom"`
	PNumber    string `json:"PNumber"`
	PCODE      string `json:"PCODE"`
	GetFlag    int    `json:"GetFlag"`
	RetCode    string `json:"RetCode"`
	RetMsg     string `json:"RetMsg"`
	CompleteDT string `json:"CompleteDT"`
	PCodeGroup int    `json:"PCodeGroup"`
	DelaySec   int    `json:"DelaySec"`
	Title      string `json:"Title"`
	Msg        string `json:"Msg"`
	FailMsg    string `json:"FailMsg"`
}

type EventCode struct {
	Idx      int    `json:"Idx"`
	ICode    string `json:"ICode"`
	SendFlag int    `json:"SendFlag"`
	Telecom  int    `json:"Telecom"`
	PNumber  string `json:"PNumber"`
	StartDT  string `json:"StartDT"`
	GroupIdx int    `json:"GroupIdx"`
}

// add struct at 2021 11 19
type HomeNews struct {
	Idx       int    `json:"Idx"`
	NewsType  int    `json:"NewsType"`
	NewsTitle string `json:"NewsTitle"`
	LinkPath  string `json:"LinkPath"`
	ImgPath   string `json:"ImgPath"`
	RegDT     string `json:"RegDT"`
}

type HomePopup struct {
	Idx        int    `json:"Idx"`
	PopupTitle string `json:"PopupTitle"`
	LinkPath   string `json:"LinkPath"`
	ImgPath    string `json:"ImgPath"`
	StartDT    string `json:"StartDT"`
	EndDT      string `json:"EndDT"`
	RegDT      string `json:"RegDT"`
}

// end add struct at 2021 11 19

type SvcMessageBatch struct {
	Idx            int    `json:"Idx"`
	Telecom        int    `json:"Telecom"`
	GPNumber       string `json:"GPNumber"`
	ReceivePNumber string `json:"ReceivePNumber"`
	MemberPNumber  string `json:"MemberPNumber"`
	Period         int    `json:"Period"`
	Cnt            int    `json:"Cnt"`
	Complete       int    `json:"Complete"`
	SendDT         string `json:"SendDT"`
	RegDT          string `json:"RegDT"`
}

type SvcMemberGuardian struct {
	Idx      int    `json:"Idx"`
	PNumber  string `json:"PNumber"`
	GPNumber string `json:"GPNumber"`
	RegDT    string `json:"RegDT"`
}

type SvcCallBatch struct {
	Idx           int    `json:"Idx"`
	Telecom       int    `json:"Telecom"`
	GPNumber      string `json:"GPNumber"`
	MemberPNumber string `json:"MemberPNumber"`
	Period        int    `json:"Period"`
	Cnt           int    `json:"Cnt"`
	Complete      int    `json:"Complete"`
	SendDT        string `json:"SendDT"`
	RegDT         string `json:"RegDT"`
}

type SvcCallHistory struct {
	Idx           int    `json:"Idx"`
	GPNumber      string `json:"GPNumber"`
	MemberPNumber string `json:"MemberPNumber"`
	ARSCallStatus int    `json:"ARSCallStatus"`
	GCallStatus   int    `json:"GCallStatus"`
	RegDT         string `json:"RegDT"`
	ARSCallRspDT  string `json:"ARSCallRspDT"`
	GCallRspDT    string `json:"GCallRspDT"`
}

type SvcEventGroup struct {
	Idx        int    `json:"Idx"`
	PCodeGroup int    `json:"PCodeGroup"`
	PCODE      string `json:"PCODE"`
}

type SvcMemberHistory struct {
	Telecom         int    `json:"Telecom"`
	PNumber         string `json:"PNumber"`
	PCode           string `json:"PCode"`
	PushToken       string `json:"PushToken"`
	ReqApp          string `json:"ReqApp"`
	ReqAppUnreg     string `json:"ReqAppUnreg"`
	Model           string `json:"Model"`
	DeviceStatus    int    `json:"DeviceStatus"`
	DeviceStatusMsg string `json:"DeviceStatusMsg"`
	LastLoginDT     string `json:"LastLoginDT"`
	UnRegType       int    `json:"UnRegType"`
	UnRegMsg        string `json:"UnRegMsg"`
	MemberType      int    `json:"MemberType"`
	PhoneType       string `json:"PhoneType"`
	Terms           int    `json:"Terms"`
	Version         string `json:"Version"`
	RegDT           string `json:"RegDT"`
	UnRegDT         string `json:"UnRegDT"`
}

type SvcTermsRefund struct {
	Idx       int    `json:"Idx"`
	SKTTitle  string `json:"SKTTitle"`
	SKTTerms  string `json:"SKTTerms"`
	KTTitle   string `json:"KTTitle"`
	KTTerms   string `json:"KTTerms"`
	LgupTitle string `json:"LgupTitle"`
	LgupTerms string `json:"LgupTerms"`
	RegDT     string `json:"RegDT"`
}

type PbaMember struct {
	Type       int    `json:"Type"`
	Telecom    int    `json:"Telecom"`
	PNumber    string `json:"PNumber"`
	PushToken  string `json:"PushToken"`
	CompleteDT string `json:"CompleteDT"`
	Complete   int    `json:"Complete"`
}

type LgupNoChargeMember struct {
	PNumber string `json:"PNumber"`
	Model   string `json:"Model"`
	SendDT  string `json:"SendDT"`
	RegDT   string `json:"RegDT"`
}

type Cnt struct {
	Cnt int `json:"Cnt"`
}

type SvcIpHistory struct {
	IP      string `json:"IP"`
	PNumber string `json:"PNumber"`
	RegDT   string `json:"RegDT"`
}

type SvcIpBlock struct {
	IP    string `json:"IP"`
	RegDT string `json:"RegDT"`
}
