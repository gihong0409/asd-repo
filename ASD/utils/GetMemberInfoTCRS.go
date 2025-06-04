package utils

import (
	"ASD/dmrsapi"
	"git.datau.co.kr/ferrari/ferrari-common/commonutils"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats/ktformats"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats/sktformats"
	"strings"
)

func GetMemberInfoTCRS(TCRSURL string, telecom string, pnumber string) map[string]interface{} {

	println("TCRSURL: ", TCRSURL, telecom, pnumber)

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

	tcrsHeader := tcrsformats.ReqHeader{CmdType: cmdType}
	tcrsBody := tcrsformats.ReqBodyPNumber{PNumber: pnumber}

	tcrsRSP := commonutils.RestfulSendData(TCRSURL+teleName, commonutils.MakeJsonData(tcrsHeader, tcrsBody))

	retData = make(map[string]interface{})

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
		println("error flag here4: ", teleName, pnumber)

	} else if strings.ToUpper(teleName) == "LGUP" {
		var tcrsRspHeader tcrsformats.RspHeader

		//원래 매핑하던 포멧에는 Age 포함하지 않는 이슈로 LGUPRSPUserInfo로 변경
		var tcrsRspBody dmrsapi.LGUPRSPUserInfo

		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody

	}

	return retData

}
