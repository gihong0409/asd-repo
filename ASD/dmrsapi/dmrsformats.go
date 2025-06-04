package dmrsapi

// 나중에 Common에 통합할 것

type AsdMember struct {
	PNumber  string `json:"PNumber"`
	Telecom  int    `json:"Telecom"`
	PCode    string `json:"PCCode"`
	Age      int    `json:"Age"`
	RegDT    string `json:"RegDT"`
	Complete int    `json:"Complete"`
}
