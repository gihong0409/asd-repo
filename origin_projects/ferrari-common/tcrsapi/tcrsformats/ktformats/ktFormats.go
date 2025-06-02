package ktformats

type RspKT struct {
	RT            string `json:"RT"`
	RT_MSG        string `json:"RT_MSG"`
	SEQUENCENO    string `json:"SEQUENCENO"`
	TRANSACTIONID string `json:"TRANSACTIONID"`
	ERROR_CODE    string `json:"ERROR_CODE"`
}

type ERRORDETAIL struct {
	ErrCode string `json:"errorcode"`
	ErrMsg  string `json:"errordescription"`
}

type RspSHUB struct {
	TRANSACTIONID string      `json:"TRANSACTIONID"`
	SEQUENCENO    string      `json:"SEQUENCENO"`
	ErrDetail     ERRORDETAIL `json:"ERRORDETAIL"`
	RetCode       string      `json:"returnCode"`
	RetMsg        string      `json:"returnDesc"`
}

type Errordetail struct {
	Errorcode        string `xml:"errorcode,omitempty" json:"errorcode,omitempty" yaml:"errorcode,omitempty"`
	Errordescription string `xml:"errordescription,omitempty" json:"errordescription,omitempty" yaml:"errordescription,omitempty"`
}

type RSPBodyHeader struct {
	ReturnCode    string      `json:"returnCode"`
	ReturnDesc    string      `json:"returnDesc,omitempty"`
	ERRORDETAIL   Errordetail `json:"ERRORDETAIL,omitempty"`
	TRANSACTIONID string      `json:"TRANSACTIONID,omitempty"`
	SEQUENCENO    string      `json:"SEQUENCENO,omitempty"`
	MARKET_GUBUN  string      `json:"MARKET_GUBUN"`
}

type RSPModelInfo struct {
	ROUTE_GBN  string `json:"ROUTE_GBN"`
	MODEL_NAME string `json:"MODEL_NAME"`
	BELLTYPE   string `json:"BELLTYPE"`
	SOUND      string `json:"SOUND"`
	COLOR      string `json:"COLOR"`
	CDMA       string `json:"CDMA"`
	MUSIC      string `json:"MUSIC"`
}

type RSPPREPAIDSEARCH struct {
	ROUTE_GBN  string `json:"ROUTE_GBN"`
	MODEL_NAME string `json:"MODEL_NAME"`
	BELLTYPE   string `json:"BELLTYPE"`
	SOUND      string `json:"SOUND"`
	COLOR      string `json:"COLOR"`
	CDMA       string `json:"CDMA"`
	MUSIC      string `json:"MUSIC"`
}

type RSPUserInfoAndKways struct {
	TRANSACTIONID   string `json:"TRANSACTIONID"`
	SEQUENCENO      string `json:"SEQUENCENO"`
	returnCode      string `json:"returnCode"`
	returnDesc      string `json:"returnDesc"`
	DISPLAINAME     string `json:"DISPLAINAME"`
	NS_CONTRACT_NUM string `json:"NS_CONTRACT_NUM"`
	NS_CUSTOMER_ID  string `json:"NS_CUSTOMER_ID"`
	MODEL_NAME      string `json:"MODEL_NAME"`
	USER_SSN_FRONT  string `json:"USER_SSN_FRONT"`
}

type RSPUserInfo struct {
	SEX     string `json:"SEX"`
	ZIPCODE string `json:"ZIPCODE"`
	AGE     string `json:"AGE"`
}

type RSPTerms struct {
	Transaction_id string `xml:"transaction_id" json:"transaction_id" yaml:"transaction_id"`
	Sequence_no    string `xml:"sequence_no" json:"sequence_no" yaml:"sequence_no"`
	Return_code    string `xml:"return_code" json:"return_code" yaml:"return_code"`
	Return_desc    string `xml:"return_desc,omitempty" json:"return_desc,omitempty" yaml:"return_desc,omitempty"`
}

type RSPDeviceInfo struct {
	ReturnCode string `xml:"errorcode,omitempty" json:"returnCode,omitempty"`
	ReturnDesc string `xml:"ReturnDesc,omitempty" json:"returnDesc,omitempty"`
	RSNCD      string `xml:"RSNCD,omitempty" json:"RSNCD,omitempty"`
}
