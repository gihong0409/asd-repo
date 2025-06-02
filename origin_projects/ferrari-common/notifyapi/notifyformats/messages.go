package notifyformats

var messages map[int]string

const (
	DB_INSERT        = 1000
	SLACK_SEND       = 2000
	SLACK_MSG_ONLY   = 3000
	SLACK_MSG_BYPASS = 4000
)

const (
	ErrorSuccess = 0
	ErrorSlack   = 9998
	ErrorUnkown  = 9999
)

const (
	BatchTypeStart        = 0 // 배치 시작
	BatchTypeFinsh        = 1 // 배치 완료
	BatchTypeNotification = 2 // 단순 배치 알림
	BatchTypeNeedCheck    = 3 // 확인 필요
)

const (
	Antiscam      = "스마트피싱보호"
	Autocall      = "오토콜"
	Mfinder       = "휴대폰분실보호"
	Familycare    = "휴대폰가족보호"
	GlobalMfinder = "휴대폰분실보호글로벌"
	CouponWallet  = "휴대폰쿠폰지갑"
)

func init() {
	messages = make(map[int]string)
	messages[ErrorSuccess] = "정상"
	messages[ErrorSlack] = "Slack 전송 오류"
	messages[ErrorUnkown] = "알 수 없는 오류"
}

func GetMsg(errCode int) string {
	return messages[errCode]
}
