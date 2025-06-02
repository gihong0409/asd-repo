package tcrsapi

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	"git.datau.co.kr/ferrari/ferrari-common/commonutils"
	"git.datau.co.kr/ferrari/ferrari-common/dmrsapi/dmrsclient"
	"git.datau.co.kr/ferrari/ferrari-common/dmrsapi/dmrsformats"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats/ktformats"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats/lgupformats"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats/sktformats"
)

type Loging func(ReqID string, v ...interface{})

func CheckServiceJoinStatus(TCWSURL string, telecom string, pnumber string, callAppName string, requestID string, printCB Loging) map[string]interface{} {
	var retData map[string]interface{}
	var teleName string
	if telecom == "0" || telecom == "1" || telecom == "2" {
		teleName = commonutils.TeleTypeNumToTelecomName(telecom)
	} else {
		teleName = telecom
	}
	tcrsHeader := tcrsformats.ReqHeader{CmdType: "SUBSCRIBERS", RequestID: requestID, CallAppName: callAppName}
	tcrsBody := tcrsformats.ReqBodyPNumber{PNumber: pnumber}

	printCB(requestID, "[SEND CheckServiceJoinStatus]", string(commonutils.MakeJsonData(tcrsHeader, tcrsBody)))

	tcrsRSP := commonutils.RestfulSendData(TCWSURL+teleName, commonutils.MakeJsonData(tcrsHeader, tcrsBody))
	retData = make(map[string]interface{})
	printCB(requestID, "[RECV CheckServiceJoinStatus]", string(tcrsRSP))

	if strings.ToUpper(teleName) == "SKT" {
		var tcrsRspHeader tcrsformats.RspHeader
		var tcrsRspBody sktformats.RspMain
		var tcrsRspBodyDetail sktformats.JoinSearchRsp
		tcrsRspBody.Body = &tcrsRspBodyDetail

		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)
		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody
		retData["BodyInfo"] = tcrsRspBodyDetail

	} else if strings.ToUpper(teleName) == "KT" {
		var tcrsRspHeader tcrsformats.RspHeader
		var tcrsRspBody ktformats.RSPBodyHeader

		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody

	} else if strings.ToUpper(teleName) == "LGUP" {
		var tcrsRspHeader tcrsformats.RspHeader
		var tcrsRspBody lgupformats.RSPUserInfo

		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody
	}

	return retData

}

// GetModel TELECOM Model 정보 확인
func GetModel(TCWSURL string, telecom string, pnumber string, callAppName string, requestID string, printCB Loging) (string, tcrsformats.DURetHeader) {
	var retData string
	var teleName string

	duHeader := tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorSuccess, ErrMsg: "정상"}

	if telecom == "0" || telecom == "1" || telecom == "2" {
		teleName = commonutils.TeleTypeNumToTelecomName(telecom)
	} else {
		teleName = telecom
	}

	if strings.ToUpper(teleName) == "SKT" {

		tcrsHeader := tcrsformats.ReqHeader{CmdType: "MEMBERSEARCH", RequestID: requestID, CallAppName: callAppName}
		tcrsBody := tcrsformats.ReqBodyPNumber{PNumber: pnumber}
		printCB(requestID, "SKT SEND DATA", string(commonutils.MakeJsonData(tcrsHeader, tcrsBody)))

		tcrsRSP := commonutils.RestfulSendData(TCWSURL+"SKT", commonutils.MakeJsonData(tcrsHeader, tcrsBody)) //사용자 조회
		printCB(requestID, "SKT RECV DATA", string(tcrsRSP))

		var tcrsRspHeader tcrsformats.RspHeader
		var tcrsRspBody sktformats.RspMain
		var tcrsRspBodyDetail sktformats.UserInfoRsp
		tcrsRspBody.Body = &tcrsRspBodyDetail
		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

		SVC_MGMT_NUM := tcrsRspBodyDetail.SVC_MGMT_NUM

		if SVC_MGMT_NUM != "" {
			tcrsHeader = tcrsformats.ReqHeader{CmdType: "DEVICEINFO", RequestID: requestID, CallAppName: callAppName}
			tcrsBody2 := tcrsformats.ReqBodyMGMTNUM{SVC_MGMT_NUM: SVC_MGMT_NUM}
			printCB(requestID, "SKT SEND DATA", string(commonutils.MakeJsonData(tcrsHeader, tcrsBody2)))

			tcrsRSP = commonutils.RestfulSendData(TCWSURL+"SKT", commonutils.MakeJsonData(tcrsHeader, tcrsBody2)) //모델조회
			printCB(requestID, "SKT RECV DATA"+string(tcrsRSP))

			var tcrsDvcinfoHeader tcrsformats.RspHeader
			var tcrsDvcinfoBody sktformats.RspMain
			var tcrsDvcinfoBodyDetail sktformats.DeviceInfoRSP
			tcrsDvcinfoBody.Body = &tcrsDvcinfoBodyDetail
			commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsDvcinfoHeader, &tcrsDvcinfoBody)

			retData = tcrsDvcinfoBodyDetail.USG_EQP_MDL_NM

		} else {
			duHeader.ErrCode = tcrsformats.ErrorFailed
			duHeader.ErrMsg = "단말조회 오류"
		}

	} else {

		tcrsHeader := tcrsformats.ReqHeader{CmdType: "MODEL", RequestID: requestID, CallAppName: callAppName}
		tcrsBody := tcrsformats.ReqBodyPNumber{PNumber: pnumber}

		printCB(requestID, "[SEND GetModel]", string(commonutils.MakeJsonData(tcrsHeader, tcrsBody)))

		tcrsRSP := commonutils.RestfulSendData(TCWSURL+teleName, commonutils.MakeJsonData(tcrsHeader, tcrsBody))
		printCB(requestID, "[SEND GetModel]", string(tcrsRSP))

		if strings.ToUpper(teleName) == "KT" {
			var tcrsRspHeader tcrsformats.RspHeader
			var tcrsRspBodyInfo ktformats.RSPModelInfo

			commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBodyInfo)

			retData = tcrsRspBodyInfo.MODEL_NAME

		} else if strings.ToUpper(teleName) == "LGUP" {
			var tcrsRspHeader tcrsformats.RspHeader
			var tcrsRspBody lgupformats.RSPUserInfo

			commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

			re := regexp.MustCompile(`[\x00-\x1F]+`)

			retData = re.ReplaceAllString(tcrsRspBody.UNIT_MDL, "")

		}

	}

	return retData, duHeader
}

// 알뜰폰
func GetSubModel(TCRSURL string, telecom string, pnumber string, logger *logrus.Logger) map[string]interface{} {
	var retData map[string]interface{}
	var teleName string
	if telecom == "0" || telecom == "1" || telecom == "2" || telecom == "10" || telecom == "11" {
		teleName = commonutils.TeleTypeNumToTelecomName(telecom)
	} else {
		teleName = telecom
	}
	tcrsHeader := tcrsformats.ReqHeader{CmdType: "MVNO_CHK"}
	tcrsBody := tcrsformats.ReqBodyPNumber{PNumber: pnumber}

	if logger != nil {
		logger.Print("[TCRS CALL]", commonutils.MakeJsonData(tcrsHeader, tcrsBody))
	}

	tcrsRSP := commonutils.RestfulSendData(TCRSURL+teleName, commonutils.MakeJsonData(tcrsHeader, tcrsBody))
	retData = make(map[string]interface{})
	if logger != nil {
		logger.Print("[TCRS CALL RESULT]", tcrsRSP)
	}

	if strings.ToUpper(teleName) == "SKT" {
		var tcrsRspHeader tcrsformats.RspHeader
		var tcrsRspBody sktformats.RspMain
		var tcrsRspBodyDetail sktformats.UserInfoRsp
		tcrsRspBody.Body = &tcrsRspBodyDetail
		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)
		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody

	} else if strings.ToUpper(teleName) == "KT" {
		var tcrsRspHeader tcrsformats.RspHeader
		var tcrsRspBody ktformats.RSPBodyHeader
		var tcrsRspBodyInfo ktformats.RSPModelInfo

		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)
		commonutils.JsonToBody([]byte(tcrsRSP), &tcrsRspBodyInfo)

		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody
		retData["BodyInfo"] = tcrsRspBodyInfo

	} else if strings.ToUpper(teleName) == "LGUP" {
		var tcrsRspHeader tcrsformats.RspHeader
		var tcrsRspBody lgupformats.RSPUserInfo

		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody
	}

	return retData
}

// GetMemberInfo Telecom Member Info 조회
func GetMemberInfo(TCWSURL string, telecom string, pnumber string, callAppName string, requestID string, printCB Loging) map[string]interface{} {
	var retData map[string]interface{}
	var teleName string

	if telecom == "0" || telecom == "1" || telecom == "2" {
		teleName = commonutils.TeleTypeNumToTelecomName(telecom)
	} else {
		teleName = telecom
	}

	cmdType := "USERINFO"

	if telecom == "1" {
		cmdType = "USERINFOANDKWAYS"
	}

	tcrsHeader := tcrsformats.ReqHeader{CmdType: cmdType, RequestID: requestID, CallAppName: callAppName}
	tcrsBody := tcrsformats.ReqBodyPNumber{PNumber: pnumber}

	printCB(requestID, "[SEND GetMemberInfo]", string(commonutils.MakeJsonData(tcrsHeader, tcrsBody)))

	tcrsRSP := commonutils.RestfulSendData(TCWSURL+teleName, commonutils.MakeJsonData(tcrsHeader, tcrsBody))
	retData = make(map[string]interface{})
	printCB(requestID, "[RECV GetMemberInfo]", string(tcrsRSP))

	if strings.ToUpper(teleName) == "SKT" {
		var tcrsRspHeader tcrsformats.RspHeader
		var tcrsRspBody sktformats.RspMain
		var tcrsRspBodyDetail sktformats.UserInfoRsp
		tcrsRspBody.Body = &tcrsRspBodyDetail

		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)
		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody
		retData["BodyInfo"] = tcrsRspBodyDetail

	} else if strings.ToUpper(teleName) == "KT" {
		var tcrsRspHeader tcrsformats.RspHeader
		var tcrsRspBody ktformats.RSPUserInfoAndKways

		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody

	} else if strings.ToUpper(teleName) == "LGUP" {
		var tcrsRspHeader tcrsformats.RspHeader
		var tcrsRspBody lgupformats.RSPUserInfo

		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody
	}

	return retData

}

func MemberAddCheck(TCWSURL string, Telecom int, PNumber string, callAppName string, RequestID string, DMRSINFO dmrsformats.DMRSInfo, printCB Loging) tcrsformats.DURetHeader {

	var retData tcrsformats.DURetHeader

	//공통 xml파일을 사용하기 위해
	DMRSINFO.APPNAME = "CMN"
	DMRSINFO.XMLNAME = "CMN"

	//BlockMember Check
	blockMembers := []dmrsformats.SvcSpMember{}
	dmrsclient.DBMCall(DMRSINFO, dmrsformats.SELECTQUERY, "CheckSPMember", []interface{}{PNumber}, &blockMembers, RequestID)

	if len(blockMembers) != 0 {
		retData.ErrCode = tcrsformats.ErrorFailed
		retData.ErrMsg = "서비스를 이용하실 수 없습니다. 고객센터(1811-4031)로 문의 해주세요."
		return retData
	}

	switch Telecom {
	case 0:
		retData = sktAddCheck(TCWSURL, PNumber, callAppName, RequestID, DMRSINFO, printCB)
	case 1:
		retData = ktAddCheck(TCWSURL, PNumber, callAppName, RequestID, DMRSINFO, printCB)
	case 2:
		retData = lgupAddCheck(TCWSURL, PNumber, callAppName, RequestID, DMRSINFO, printCB)
	}

	return retData
}

func isBlockModelAll(phonemodelName string) bool {

	SIMOnly := []string{"OMPHONE", "STDPHONE", "OPENMODEL", "PTA", "핸드셋"} // 유심 단독
	for _, chkModel := range SIMOnly {
		if strings.Contains(phonemodelName, chkModel) {
			return true
		}
	}

	FeaturePhone := []string{"LM-Y", "LG-T"} // 피쳐폰
	for _, chkModel := range FeaturePhone {
		if strings.Contains(phonemodelName, chkModel) {
			return true
		}
	}

	return false

}

func isBlockModelAllIPhone(telecom int, RequestID, phonemodelName string, DMRSINFO dmrsformats.DMRSInfo) bool {

	re := regexp.MustCompile(`A\d{4}`)
	if re.MatchString(phonemodelName) {
		return true
	}

	re2 := regexp.MustCompile(`IP1`)
	if re2.MatchString(phonemodelName) {
		return true
	}

	// DB 조회
	model := []tcrsformats.Model{}
	dmrsclient.DBMCall(DMRSINFO, dmrsformats.SELECTQUERY, "CheckIphone", []interface{}{phonemodelName, telecom}, &model, RequestID)
	if len(model) > 0 {
		return true
	}

	return false

}

func isBlockModel(phonemodelName string) bool {
	r, _ := regexp.Compile("A[1-2]{1}[0-9]{3}-") // LGU+ IPHONE
	iphoneName := r.FindStringIndex(phonemodelName)

	if len(iphoneName) > 0 { //LGU+ IPHONE
		return true
	} else if strings.Index(phonemodelName, "IPHONE") > -1 ||
		strings.Index(phonemodelName, "APPLE") > -1 ||
		strings.Index(phonemodelName, "AIP") > -1 {
		return true
	}

	return false
}

// 피쳐폰 조회 추가
func isFeaturePhone(phonemodelName, RequestID, callAppName, telecom, PNumber string, DMRSINFO dmrsformats.DMRSInfo) bool {
	if phonemodelName != "" {
		DMRSINFO.APPNAME = "CMN"
		DMRSINFO.XMLNAME = "CMN"

		cnt := []dmrsformats.Cnt{}
		dmrsclient.DBMCall(DMRSINFO, dmrsformats.SELECTQUERY, "CheckFeaturePhone", []interface{}{phonemodelName}, &cnt, RequestID)

		if len(cnt) > 0 {
			if cnt[0].Cnt > 0 {
				return true
			}
		}
	} else {
		DMRSINFO.APPNAME = "CMN"
		DMRSINFO.XMLNAME = "CMN"

		dmrsclient.DBMCall(DMRSINFO, dmrsformats.CUDQUERY, "InsertNilModel", []interface{}{callAppName, telecom, PNumber}, nil, RequestID)
	}
	return false
}

func sktAddCheck(TCWSURL string, PNumber string, callAppName string, RequestID string, DMRSINFO dmrsformats.DMRSInfo, printCB Loging) tcrsformats.DURetHeader {
	retData := tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorSuccess, ErrMsg: tcrsformats.GetMsg(tcrsformats.ErrorSuccess)}

	tcrsHeader := tcrsformats.ReqHeader{CmdType: "MEMBERSEARCH", RequestID: RequestID, CallAppName: callAppName}
	tcrsBody := tcrsformats.ReqBodyPNumber{PNumber: PNumber}
	printCB(RequestID, "SKT SEND DATA", string(commonutils.MakeJsonData(tcrsHeader, tcrsBody)))

	tcrsRSP := commonutils.RestfulSendData(TCWSURL+"SKT", commonutils.MakeJsonData(tcrsHeader, tcrsBody)) //사용자 조회
	printCB(RequestID, "SKT RECV DATA", string(tcrsRSP))

	var tcrsRspHeader tcrsformats.RspHeader
	var tcrsRspBody sktformats.RspMain
	var tcrsRspBodyDetail sktformats.UserInfoRsp
	tcrsRspBody.Body = &tcrsRspBodyDetail
	commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)
	if tcrsRspBody.Header.ReturnCode == "PCI_DTS_E3162" && tcrsRspBody.Header.Result == "F" {
		return tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorNoTel, ErrMsg: tcrsformats.GetMsg(tcrsformats.ErrorNoTel)}
	} else {

		memberAge := commonutils.GetAge(tcrsRspBodyDetail.SSN_BIRTH_DT)

		if memberAge < 19 {
			printCB(RequestID, "미성년자 가입 불가", PNumber)
			return tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorMinor, ErrMsg: tcrsformats.GetMsg(tcrsformats.ErrorMinor)}
		}

		SVC_MGMT_NUM := tcrsRspBodyDetail.SVC_MGMT_NUM

		if SVC_MGMT_NUM != "" {
			tcrsHeader = tcrsformats.ReqHeader{CmdType: "DEVICEINFO", RequestID: RequestID, CallAppName: callAppName}
			tcrsBody2 := tcrsformats.ReqBodyMGMTNUM{SVC_MGMT_NUM: SVC_MGMT_NUM}
			printCB(RequestID, "SKT SEND DATA", string(commonutils.MakeJsonData(tcrsHeader, tcrsBody2)))

			tcrsRSP = commonutils.RestfulSendData(TCWSURL+"SKT", commonutils.MakeJsonData(tcrsHeader, tcrsBody2)) //모델조회

			printCB(RequestID, "SKT RECV DATA"+string(tcrsRSP))

			var tcrsRspHeader tcrsformats.RspHeader
			var tcrsRspBody sktformats.RspMain
			var tcrsRspBodyDetail sktformats.DeviceInfoRSP
			tcrsRspBody.Body = &tcrsRspBodyDetail
			commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

			phonemodelName := tcrsRspBodyDetail.USG_EQP_MDL_NM

			if isBlockModelAllIPhone(0, RequestID, phonemodelName, DMRSINFO) {
				return tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorFailed, ErrMsg: tcrsformats.GetMsg(tcrsformats.ErrorIphone)} //아이폰
			}
			if isBlockModelAll(phonemodelName) {
				return tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorFailed, ErrMsg: tcrsformats.GetMsg(tcrsformats.ErrorFeaturePhone)} //유심단독
			}
			if isFeaturePhone(phonemodelName, RequestID, callAppName, "0", PNumber, DMRSINFO) {
				return tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorFailed, ErrMsg: tcrsformats.GetMsg(tcrsformats.ErrorFeaturePhone)} //피쳐폰
			}

			tcrsHeader = tcrsformats.ReqHeader{CmdType: "LOSTPHONECHECK", RequestID: RequestID, CallAppName: callAppName}
			tcrsBody3 := tcrsformats.ReqBodyLostPhoneCheck{EQP_MDL_CD: tcrsRspBodyDetail.USG_EQP_MDL_CD, EQP_SER_NUM: tcrsRspBodyDetail.USG_EQP_SER_NUM, SVC_MGMT_NUM: SVC_MGMT_NUM}
			printCB(RequestID, "SKT SEND DATA", string(commonutils.MakeJsonData(tcrsHeader, tcrsBody3)))

			tcrsRSP = commonutils.RestfulSendData(TCWSURL+"SKT", commonutils.MakeJsonData(tcrsHeader, tcrsBody2)) //분실폰 체크
			printCB(RequestID, "SKT RECV DATA"+string(tcrsRSP))

			var tcrsRspHeaderLostPhone tcrsformats.RspHeader
			var tcrsRspBodyLostPhone sktformats.RspMain
			var tcrsRspBodyDetail2 sktformats.LostPhoneCheckRSP_Hub
			tcrsRspBodyLostPhone.Body = &tcrsRspBodyDetail2
			commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeaderLostPhone, &tcrsRspBodyLostPhone)

			if tcrsRspBodyLostPhone.Header.Result == "S" && tcrsRspBodyLostPhone.Header.ReturnCode == "00" {
				if len(tcrsRspBodyDetail2.EqpLossInfoList.EqpLossInfo) > 0 {
					fmt.Println("분실폰", PNumber, tcrsRspBodyDetail2)
					return tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorFailed, ErrMsg: tcrsformats.GetMsg(tcrsformats.ErrorLossPhone)} //분실폰
				}
			}

		} else {
			return tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorFailed, ErrMsg: "단말조회 오류 : 통신사 선택 오류 입니다."}
		}

	}

	return retData

}

func ktAddCheck(TCWSURL string, PNumber string, callAppName string, RequestID string, DMRSINFO dmrsformats.DMRSInfo, printCB Loging) tcrsformats.DURetHeader {
	retData := tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorSuccess, ErrMsg: tcrsformats.GetMsg(tcrsformats.ErrorSuccess)}

	tcrsHeader := tcrsformats.ReqHeader{CmdType: "MODEL", RequestID: RequestID, CallAppName: callAppName}
	tcrsBody := tcrsformats.ReqBodyPNumber{PNumber: PNumber}
	printCB(RequestID, "KT SEND DATA"+string(commonutils.MakeJsonData(tcrsHeader, tcrsBody)))

	tcrsRSP := commonutils.RestfulSendData(TCWSURL+"KT", commonutils.MakeJsonData(tcrsHeader, tcrsBody)) //사용자 조회
	printCB(RequestID, "KT RECV DATA"+string(tcrsRSP))

	var tcrsRspHeader tcrsformats.RspHeader
	var tcrsRspBody ktformats.RSPModelInfo

	commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)
	if tcrsRspBody.MODEL_NAME == "" {
		return tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorNoTel, ErrMsg: tcrsformats.GetMsg(tcrsformats.ErrorNoTel)} // 사용자 조회 안됨
	}

	if isBlockModelAllIPhone(1, RequestID, tcrsRspBody.MODEL_NAME, DMRSINFO) {
		return tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorFailed, ErrMsg: tcrsformats.GetMsg(tcrsformats.ErrorIphone)} //아이폰
	}
	if isBlockModelAll(tcrsRspBody.MODEL_NAME) {
		return tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorFailed, ErrMsg: tcrsformats.GetMsg(tcrsformats.ErrorFeaturePhone)} //유심단독
	}
	if isFeaturePhone(tcrsRspBody.MODEL_NAME, RequestID, callAppName, "1", PNumber, DMRSINFO) {
		return tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorFailed, ErrMsg: tcrsformats.GetMsg(tcrsformats.ErrorFeaturePhone)} //피쳐폰
	}

	tcrsHeader = tcrsformats.ReqHeader{CmdType: "DEVICESTATUS", RequestID: RequestID, CallAppName: callAppName}
	printCB(RequestID, "KT SEND DATA"+string(commonutils.MakeJsonData(tcrsHeader, tcrsBody)))

	tcrsRSP = commonutils.RestfulSendData(TCWSURL+"KT", commonutils.MakeJsonData(tcrsHeader, tcrsBody)) //분실폰 체크
	printCB(RequestID, "KT RECV DATA"+string(tcrsRSP))

	var tcrsRspBody2 ktformats.RSPDeviceInfo
	commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody2)

	if tcrsRspBody2.RSNCD != "" {
		return tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorFailed, ErrMsg: tcrsformats.GetMsg(tcrsformats.ErrorLossPhone)} //분실폰
	}

	return retData
}

func lgupAddCheck(TCWSURL string, PNumber string, callAppName string, RequestID string, DMRSINFO dmrsformats.DMRSInfo, printCB Loging) tcrsformats.DURetHeader {
	retData := tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorSuccess, ErrMsg: tcrsformats.GetMsg(tcrsformats.ErrorSuccess)}

	//공통 xml파일을 사용하기 위해
	DMRSINFO.APPNAME = "CMN"
	DMRSINFO.XMLNAME = "CMN"

	//freeMember Check
	freeMembers := []dmrsformats.SvcFreemembers{}
	dmrsclient.DBMCall(DMRSINFO, dmrsformats.SELECTQUERY, "CheckFreeMember", []interface{}{PNumber}, &freeMembers, RequestID)

	if len(freeMembers) < 1 {
		//국번 8080 차단
		pnumber := PNumber[3:7]
		if pnumber == "8080" || pnumber == "8400" {
			return tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorLGUP8080, ErrMsg: tcrsformats.GetMsg(tcrsformats.ErrorLGUP8080)}
		}
	}

	tcrsHeader := tcrsformats.ReqHeader{CmdType: "MODEL", RequestID: RequestID, CallAppName: callAppName}
	tcrsBody := tcrsformats.ReqBodyPNumber{PNumber: PNumber}
	printCB(RequestID, "LGU SEND DATA"+string(commonutils.MakeJsonData(tcrsHeader, tcrsBody)))

	tcrsRSP := commonutils.RestfulSendData(TCWSURL+"LGUP", commonutils.MakeJsonData(tcrsHeader, tcrsBody)) //사용자 조회

	printCB(RequestID, "LGU RECV DATA"+string(tcrsRSP))

	var tcrsRspHeader tcrsformats.RspHeader
	var tcrsRspBody lgupformats.IFDATARsp

	commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

	if tcrsRspBody.UNIT_MDL == "" {
		return tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorNoTel, ErrMsg: tcrsformats.GetMsg(tcrsformats.ErrorNoTel)} // 사용자 조회 안됨
	}

	if isBlockModelAllIPhone(2, RequestID, tcrsRspBody.UNIT_MDL, DMRSINFO) {
		return tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorFailed, ErrMsg: tcrsformats.GetMsg(tcrsformats.ErrorIphone)} //아이폰
	}
	if isBlockModelAll(tcrsRspBody.UNIT_MDL) {
		return tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorFailed, ErrMsg: tcrsformats.GetMsg(tcrsformats.ErrorFeaturePhone)} //유심 단독
	}
	if isFeaturePhone(tcrsRspBody.UNIT_MDL, RequestID, callAppName, "1", PNumber, DMRSINFO) {
		return tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorFailed, ErrMsg: tcrsformats.GetMsg(tcrsformats.ErrorFeaturePhone)} //피쳐폰
	}

	return retData

}

// GetFreeMemberSearch TELECOM Model 정보 확인
func MemberAdd(TCWSURL string, telecom string, pnumber string, callAppName string, requestID string, DMRSINFO dmrsformats.DMRSInfo, printCB Loging) (map[string]interface{}, tcrsformats.DURetHeader) {

	var retData map[string]interface{}
	var teleName string
	var duHeader tcrsformats.DURetHeader

	//공통 xml파일을 사용하기 위해
	DMRSINFO.APPNAME = "CMN"
	DMRSINFO.XMLNAME = "CMN"

	//freeMember Check
	freeMembers := []dmrsformats.SvcFreemembers{}
	dmrsclient.DBMCall(DMRSINFO, dmrsformats.SELECTQUERY, "CheckFreeMember", []interface{}{pnumber}, &freeMembers, requestID)

	if len(freeMembers) != 0 {
		duHeader.ErrCode = tcrsformats.ErrorSuccess
		duHeader.ErrMsg = tcrsformats.GetMsg(tcrsformats.ErrorSuccess)
		return retData, duHeader
	}

	if telecom == "0" || telecom == "1" || telecom == "2" {
		teleName = commonutils.TeleTypeNumToTelecomName(telecom)
	} else {
		teleName = telecom
	}
	if pnumber == "01010001000" {
		return nil, tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorSuccess, ErrMsg: "정상"}
	}
	retData = make(map[string]interface{})

	if strings.ToUpper(teleName) == "SKT" {
		tcrsHeader := tcrsformats.ReqHeader{CmdType: "JOINMEMBER", RequestID: requestID, CallAppName: callAppName}
		tcrsBody := tcrsformats.ReqBodyPNumber{PNumber: pnumber}

		requestData := commonutils.MakeJsonData(tcrsHeader, tcrsBody)
		printCB(requestID, "SKT SEND DATA"+string(requestData))

		tcrsRSP := commonutils.RestfulSendData(TCWSURL+teleName, requestData)

		printCB(requestID, "SKT RECV DATA"+string(tcrsRSP))

		var tcrsRspHeader tcrsformats.RspHeader
		var tcrsRspBody sktformats.RspMain
		var tcrsRspBodyDetail sktformats.JoinDelRsp
		tcrsRspBody.Body = tcrsRspBodyDetail
		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

		var errData = ""
		if tcrsRspBody.Header.ReturnDesc != "" {
			s := strings.Split(tcrsRspBody.Header.ReturnDesc, "|")
			if len(s) > 0 {
				errData = s[0]
			}
		}

		if tcrsRspBody.Header.Result == "S" || errData == "ZINVE8101" { //이미 가입자
			duHeader.ErrCode = tcrsformats.ErrorSuccess
			duHeader.ErrMsg = tcrsformats.GetMsg(tcrsformats.ErrorSuccess)
		} else {
			duHeader.ErrCode = tcrsformats.ErrorFailed
			duHeader.ErrMsg = tcrsRspBody.Header.ReturnDesc
		}

		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody
		retData["BodyInfo"] = tcrsRspBodyDetail

	} else if strings.ToUpper(teleName) == "KT" {

		//약관
		tcrsHeader := tcrsformats.ReqHeader{CmdType: "TERMS_OK", RequestID: requestID, CallAppName: callAppName}
		tcrsTermsBody := tcrsformats.ReqBodyTerms{PNumber: pnumber, Terms1: "Y", Terms2: "Y", Terms3: "Y", Terms4: "Y", Terms5: "Y"}
		printCB(requestID, "KT SEND DATA"+string(commonutils.MakeJsonData(tcrsHeader, tcrsTermsBody)))

		tcrsRSP := commonutils.RestfulSendData(TCWSURL+teleName, commonutils.MakeJsonData(tcrsHeader, tcrsTermsBody))
		printCB(requestID, "KT RECV DATA"+string(tcrsRSP))

		var tcrsRspHeader tcrsformats.RspHeader
		var tcrsRspBody ktformats.RSPTerms

		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

		if tcrsRspBody.Return_code != "1" {
			duHeader.ErrCode = tcrsformats.ErrorFailed
			duHeader.ErrMsg = tcrsRspBody.Return_desc
		} else { //가입

			tcrsHeader = tcrsformats.ReqHeader{CmdType: "JOINMEMBER", RequestID: requestID, CallAppName: callAppName}
			tcrsBody := tcrsformats.ReqBodyPNumber{PNumber: pnumber}

			printCB(requestID, "KT SEND DATA"+string(commonutils.MakeJsonData(tcrsHeader, tcrsBody)))

			tcrsRSP = commonutils.RestfulSendData(TCWSURL+teleName, commonutils.MakeJsonData(tcrsHeader, tcrsBody))

			printCB(requestID, "KT RECV DATA"+string(tcrsRSP))

			var tcrsRspHeader tcrsformats.RspHeader
			var tcrsRspBody ktformats.RspSHUB

			commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

			if tcrsRspBody.RetCode == "1" || tcrsRspBody.ErrDetail.ErrCode == "SANTA1106" {
				duHeader.ErrCode = tcrsformats.ErrorSuccess
				duHeader.ErrMsg = tcrsformats.GetMsg(tcrsformats.ErrorSuccess)

			} else {
				duHeader.ErrCode = tcrsformats.ErrorFailed
				duHeader.ErrMsg = tcrsRspBody.ErrDetail.ErrMsg
			}

		}

		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody

	} else if strings.ToUpper(teleName) == "LGUP" {

		//국번 8080 차단
		pnmuberChk := pnumber[3:7]
		if pnmuberChk == "8080" {
			return nil, tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorFailed, ErrMsg: "8080 차단"}
		}

		tcrsHeader := tcrsformats.ReqHeader{CmdType: "JOINMEMBER", RequestID: requestID, CallAppName: callAppName}
		tcrsBody := tcrsformats.ReqBodyPNumber{PNumber: pnumber}

		printCB(requestID, "LGU+ SEND DATA"+string(commonutils.MakeJsonData(tcrsHeader, tcrsBody)))

		tcrsRSP := commonutils.RestfulSendData(TCWSURL+teleName, commonutils.MakeJsonData(tcrsHeader, tcrsBody))

		printCB(requestID, "LGU+ RECV DATA"+string(tcrsRSP))

		var tcrsRspHeader tcrsformats.RspHeader

		commonutils.JsonToHader([]byte(tcrsRSP), &tcrsRspHeader)

		if tcrsRspHeader.ErrCode == 0 || tcrsRspHeader.ErrCode == 12 { // 이미 가입자
			duHeader.ErrCode = tcrsformats.ErrorSuccess
			duHeader.ErrMsg = tcrsformats.GetMsg(tcrsformats.ErrorSuccess)
		} else {

			duHeader.ErrCode = tcrsformats.ErrorFailed
			duHeader.ErrMsg = strconv.Itoa(tcrsRspHeader.ErrCode) + ": 통신사 가입 오류"
		}
		retData["Header"] = tcrsRspHeader
	}

	return retData, duHeader
}

// GetFreeMemberSearch TELECOM Model 정보 확인
func MemberDel(TCWSURL string, telecom string, pnumber string, callAppName string, requestID string, printCB Loging) (map[string]interface{}, tcrsformats.DURetHeader) {
	var retData map[string]interface{}
	var teleName string
	var duHeader tcrsformats.DURetHeader

	if telecom == "0" || telecom == "1" || telecom == "2" {
		teleName = commonutils.TeleTypeNumToTelecomName(telecom)
	} else {
		teleName = telecom
	}

	if pnumber == "01010001000" {
		return nil, tcrsformats.DURetHeader{ErrCode: tcrsformats.ErrorSuccess, ErrMsg: "정상"}
	}

	retData = make(map[string]interface{})

	if strings.ToUpper(teleName) == "SKT" {
		tcrsHeader := tcrsformats.ReqHeader{CmdType: "DELMEMBER", RequestID: requestID, CallAppName: callAppName}
		tcrsBody := tcrsformats.ReqBodyPNumber{PNumber: pnumber}

		requestData := commonutils.MakeJsonData(tcrsHeader, tcrsBody)

		printCB(requestID, "SKT SEND DATA", string(requestData))
		tcrsRSP := commonutils.RestfulSendData(TCWSURL+teleName, requestData)
		printCB(requestID, "SKT RECV DATA", string(tcrsRSP))

		var tcrsRspHeader tcrsformats.RspHeader
		var tcrsRspBody sktformats.RspMain
		var tcrsRspBodyDetail sktformats.JoinDelRsp
		tcrsRspBody.Body = tcrsRspBodyDetail
		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

		var errData = ""
		if tcrsRspBody.Header.ReturnDesc != "" {
			s := strings.Split(tcrsRspBody.Header.ReturnDesc, "|")
			if len(s) > 0 {
				errData = s[0]
			}
		}

		if tcrsRspBody.Header.Result == "S" || errData == "ZNGME0068" { //이미 해지자
			duHeader.ErrCode = tcrsformats.ErrorSuccess
			duHeader.ErrMsg = tcrsformats.GetMsg(tcrsformats.ErrorSuccess)
		} else {
			duHeader.ErrCode = tcrsformats.ErrorFailed
			duHeader.ErrMsg = tcrsRspBody.Header.ReturnDesc
		}

		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody
		retData["BodyInfo"] = tcrsRspBodyDetail

	} else if strings.ToUpper(teleName) == "KT" {
		tcrsHeader := tcrsformats.ReqHeader{CmdType: "DELMEMBER", RequestID: requestID, CallAppName: callAppName}
		tcrsBody := tcrsformats.ReqBodyPNumber{PNumber: pnumber}

		printCB(requestID, "KT SEND DATA", commonutils.MakeJsonData(tcrsHeader, tcrsBody))

		tcrsRSP := commonutils.RestfulSendData(TCWSURL+teleName, commonutils.MakeJsonData(tcrsHeader, tcrsBody))

		printCB(requestID, "KT RECV DATA", tcrsRSP)

		var tcrsRspHeader tcrsformats.RspHeader
		var tcrsRspBody ktformats.RspSHUB

		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)
		commonutils.JsonToBody([]byte(tcrsRSP), &tcrsRspBody)

		if tcrsRspBody.RetCode == "1" ||
			tcrsRspBody.ErrDetail.ErrCode == "SANTA1111" ||
			tcrsRspBody.ErrDetail.ErrCode == "SANTA1022" ||
			tcrsRspBody.ErrDetail.ErrCode == "SANTA1024" ||
			tcrsRspBody.ErrDetail.ErrCode == "SANTA1026" ||
			tcrsRspBody.ErrDetail.ErrCode == "SANTA1028" {
			duHeader.ErrCode = tcrsformats.ErrorSuccess
			duHeader.ErrMsg = tcrsformats.GetMsg(tcrsformats.ErrorSuccess)
		} else {
			duHeader.ErrCode = tcrsformats.ErrorFailed
			duHeader.ErrMsg = tcrsRspBody.ErrDetail.ErrMsg
		}

		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody
		retData["BodyInfo"] = tcrsRspBody

	} else if strings.ToUpper(teleName) == "LGUP" {
		tcrsHeader := tcrsformats.ReqHeader{CmdType: "DELMEMBER", RequestID: requestID, CallAppName: callAppName}
		tcrsBody := tcrsformats.ReqBodyPNumber{PNumber: pnumber}

		printCB(requestID, "LGUP SEND DATA", commonutils.MakeJsonData(tcrsHeader, tcrsBody))

		tcrsRSP := commonutils.RestfulSendData(TCWSURL+teleName, commonutils.MakeJsonData(tcrsHeader, tcrsBody))

		printCB(requestID, "LGUP RECV DATA", tcrsRSP)

		var tcrsRspHeader tcrsformats.RspHeader

		commonutils.JsonToHader([]byte(tcrsRSP), &tcrsRspHeader)
		// commonutils.JsonToBody([]byte(tcrsRSP), &tcrsRspBody)

		if tcrsRspHeader.ErrCode == 0 ||
			tcrsRspHeader.ErrCode == 20 ||
			tcrsRspHeader.ErrCode == 42 ||
			tcrsRspHeader.ErrCode == 94 {
			duHeader.ErrCode = tcrsformats.ErrorSuccess
			duHeader.ErrMsg = tcrsformats.GetMsg(tcrsformats.ErrorSuccess)
		} else {
			duHeader.ErrCode = tcrsformats.ErrorFailed
			duHeader.ErrMsg = tcrsRspHeader.ErrMsg
		}

		retData["Header"] = tcrsRspHeader
		// retData["BodyInfo"] = tcrsRspBody
	}

	return retData, duHeader
}

// TermsUploadKTOnly KT로 보내는 약관동의 마케팅수신동의는 빠짐 뭐가 넘어가도 Y 임
func TermsUploadKTOnly(TCWSURL string, pnumber string, callAppName string, requestID string, printCB Loging) map[string]interface{} {
	var retData map[string]interface{}
	retData = make(map[string]interface{})

	tcrsHeader := tcrsformats.ReqHeader{CmdType: "TERMS_OK", RequestID: requestID, CallAppName: callAppName}
	tcrsBody := tcrsformats.ReqBodyTerms{PNumber: pnumber, Terms1: "Y", Terms2: "Y", Terms3: "Y", Terms4: "Y", Terms5: "Y"}

	printCB(requestID, "[KT SEND TermsUploadKTOnly]", commonutils.MakeJsonData(tcrsHeader, tcrsBody))

	tcrsRSP := commonutils.RestfulSendData(TCWSURL+"KT", commonutils.MakeJsonData(tcrsHeader, tcrsBody))

	printCB(requestID, "[KT RECV TermsUploadKTOnly]", tcrsRSP)

	var tcrsRspHeader tcrsformats.RspHeader
	var tcrsRspBody ktformats.RSPTerms

	commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

	retData["Header"] = tcrsRspHeader
	retData["Body"] = tcrsRspBody

	return retData
}

// 사용양이 있을때 SKT로 전송
func ServiceUse(TCWSURL string, pnumber string, callAppName string, requestID string, printCB Loging) map[string]interface{} {
	var retData map[string]interface{}
	retData = make(map[string]interface{})

	tcrsHeader := tcrsformats.ReqHeader{CmdType: "LONGTIMENOTUSE", RequestID: requestID, CallAppName: callAppName}
	tcrsBody := sktformats.UseServiceReq{PNumber: pnumber}

	printCB(requestID, "[SKT SEND ServiceUse]", commonutils.MakeJsonData(tcrsHeader, tcrsBody))

	tcrsRSP := commonutils.RestfulSendData(TCWSURL+"SKT", commonutils.MakeJsonData(tcrsHeader, tcrsBody))

	printCB(requestID, "[SKT RECV ServiceUse]", tcrsRSP)

	var tcrsRspHeader tcrsformats.RspHeader
	var tcrsRspBody sktformats.LongtimeNotUseTrID

	commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

	retData["Header"] = tcrsRspHeader
	retData["Body"] = tcrsRspBody

	return retData
}

// TermsUploadLGUPOnly LGUP로 보내는 약관동의 마케팅수신동의는 빠짐 뭐가 넘어가도 Y 임
// KT와 마찬가지
// func TermsUploadLGUPOnly(TCWSURL string, pnumber string, requestID string, printCB Loging) map[string]interface{} {
// 	var retData map[string]interface{}
// 	retData = make(map[string]interface{})

// 	tcrsHeader := tcrsformats.ReqHeader{CmdType: "TERMS_OK"}
// 	tcrsBody := tcrsformats.ReqBodyTerms{PNumber: pnumber, Terms1: "Y", Terms2: "Y", Terms3: "Y", Terms4: "Y", Terms5: "Y"}

// 	if logger != nil {
// 		logger.Print("[TCWS CALL]", commonutils.MakeJsonData(tcrsHeader, tcrsBody))
// 	}

// 	tcrsRSP := commonutils.RestfulSendData(TCWSURL+"LGUP", commonutils.MakeJsonData(tcrsHeader, tcrsBody))

// 	if logger != nil {
// 		logger.Print("[TCWS CALL MENBER ADD RESULT]", tcrsRSP)
// 	}

// 	// var tcrsRspHeader tcrsformats.RspHeader
// 	// var tcrsRspBody lgupformats.RSPTerms

// 	// commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

// 	// retData["Header"] = tcrsRspHeader
// 	// retData["Body"] = tcrsRspBody

// 	return retData
// }
