package dmrsformats

var messages map[int]string

//this const init error code
const (
	ErrorSuccess = 0 // return
	ErrorSel     = 1001
	ErrorExec    = 1002
)

const (
	SELECTQUERY = 1
	CUDQUERY    = 2
)

func init() {
	messages = make(map[int]string)
	messages[ErrorSuccess] = "정상"
	messages[ErrorSel] = "Select 오류"
	messages[ErrorExec] = "Exec CUD 오류"

}

// GetMsg return errMsg
func GetMsg(errCode int) string {
	return messages[errCode]
}
