package dmrsapi

// 나중에 Common에 통합할 것
type AsdMember struct {
	PNumber  string `json:"PNumber"`
	Telecom  int    `json:"Telecom"`
	PCode    string `json:"PCCode"`
	Age      string `json:"Age"`
	RegDT    string `json:"RegDT"`
	Complete bool   `json:"Complete"`
}
