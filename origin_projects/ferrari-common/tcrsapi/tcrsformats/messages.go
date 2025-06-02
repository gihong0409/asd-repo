package tcrsformats

var messages map[int]string

const ErrorSuccess = 0
const ErrorFailed = 1

const ErrorTelError = 100001
const ErrorNoSupportTel = 100002

const ErrorNoTel = 100003
const ErrorCanNotUseApi = 100004

const ErrorIphone = 200001
const ErrorLossPhone = 200002
const ErrorMinor = 200003
const ErrorLGUP8080 = 200004
const ErrorFeaturePhone = 200005
const ErrorPhoneModel = 200006

const ErrorCmdType = 900002
const ErrorParameter = 900003
const ErrorNonePacket = 900004

func init() {
	messages = make(map[int]string)
	messages[ErrorSuccess] = "완료"
	messages[ErrorParameter] = "파라메터 오류입니다."
	messages[ErrorNonePacket] = "알 수 없는 명령입니다."
	messages[ErrorNoTel] = "통신사 조회 오류입니다. 통신사를 확인해주세요."
	messages[ErrorCanNotUseApi] = "지원하지 않는 기능 입니다."

	messages[ErrorTelError] = "통신사 오류 입니다. Body를 확인 하세요."
	messages[ErrorFailed] = "FCM 발송 오류"
	messages[ErrorNoSupportTel] = "지원하지 않는 통신사 입니다."
	messages[ErrorCmdType] = "지원하지 않는 CmdType 입니다."
	messages[ErrorIphone] = "죄송합니다.아이폰은 서비스를 지원하지 않아 서비스에 가입하실 수 없습니다."
	messages[ErrorLossPhone] = "휴대폰분실 신고된 번호는 서비스 가입이 불가합니다."
	messages[ErrorMinor] = "미성년자는 서비스 가입이 불가합니다."
	messages[ErrorLGUP8080] = "서비스를 이용하실 수 없습니다. 고객센터(1811-4031)로 문의 해주세요."
	messages[ErrorFeaturePhone] = "서비스 이용이 불가능한 단말기입니다."
	messages[ErrorPhoneModel] = "가입 불가 단말기입니다. 이용 관련 문의는 고객센터(1811-4031)로 연락 부탁 드립니다."
}

func GetMsg(errCode int) string {
	return messages[errCode]
}
