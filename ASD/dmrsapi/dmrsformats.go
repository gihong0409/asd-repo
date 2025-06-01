package dmrsapi

type AsdMember struct {
	Telecom  int    `json:"Telecom"`
	PCode    string `json:"PCCode"`
	PNumber  string `json:"PNumber"`
	Age      int    `json:"Age"`
	RegDT    string `json:"RegDT"`
	Complete int    `json:"Complete"`
}
