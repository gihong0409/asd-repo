package utils

import (
	"ASD/dmrsapi"
	"bytes"
	"fmt"
	"git.datau.co.kr/ferrari/ferrari-common/commonutils"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats/ktformats"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats/sktformats"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetMemberInfoTCRS(TCRSURL string, telecom string, pnumber string) map[string]interface{} {

	println("TCRSURL: ", TCRSURL, telecom, pnumber)

	var retData map[string]interface{}
	var teleName string

	if telecom == "0" || telecom == "1" || telecom == "2" {
		teleName = commonutils.TeleTypeNumToTelecomName(telecom)
		println("error flag here1: ", teleName, pnumber)
	} else {
		teleName = telecom

	}

	cmdType := "USERINFO"

	if telecom == "1" {
		cmdType = "USERINFOANDKWAYS"
		println("error flag here2: ", teleName, pnumber)

	}

	tcrsHeader := tcrsformats.ReqHeader{CmdType: cmdType}
	tcrsBody := tcrsformats.ReqBodyPNumber{PNumber: pnumber}
	println("tcrsHeader: ", tcrsHeader.CmdType, "tcrsBody: ", tcrsBody.PNumber)

	tcrsRSP := RestfulSendData(TCRSURL+"KT", commonutils.MakeJsonData(tcrsHeader, tcrsBody))
	println("tcrsRSP: ", tcrsRSP)

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
		println("error flag here3: ", teleName, pnumber)

		var tcrsRspHeader tcrsformats.RspHeader
		var tcrsRspBody ktformats.RSPUserInfoAndKways

		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody
		println("error flag here4: ", teleName, pnumber)

	} else if strings.ToUpper(teleName) == "LGUP" {
		var tcrsRspHeader tcrsformats.RspHeader
		println("error flag here 5: ", teleName, pnumber)

		//원래 매핑하던 포멧에는 Age 포함하지 않는 이슈로 LGUPRSPUserInfo로 변경
		var tcrsRspBody dmrsapi.LGUPRSPUserInfo

		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody
		println("error flag here 6: ", teleName, pnumber)

	}

	return retData

}

func RestfulSendData(url string, inData []byte) []byte {
	reqData := bytes.NewBuffer(inData)
	resp, err := http.Post(url, "application/json", reqData)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("resp:", resp)
	var f []byte
	// var strJson strng
	if resp != nil {
		f, _ = ioutil.ReadAll(resp.Body)
		resp.Body.Close()
	}
	fmt.Println("f:", f)

	return f
}
