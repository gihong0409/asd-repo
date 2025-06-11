package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"git.datau.co.kr/earth/earth-asd/formats"
	"git.datau.co.kr/ferrari/ferrari-common/commonutils"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats/ktformats"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats/sktformats"
)

// GetMemberInfoTCRS AGE_OUT을 포함하는 LGUPRSPUserInfo 사용이 불가피함에 따라 ferrari-common에서 현재 프로젝트로 가져와 수정하였습니다.
func GetMemberInfoTCRS(TCRSURL string, telecom string, pnumber string) map[string]interface{} {

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

	tcrsRSP := RestfulSendData(TCRSURL+teleName, commonutils.MakeJsonData(tcrsHeader, tcrsBody))

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

	} else if strings.ToUpper(teleName) == "LGUP" {
		var tcrsRspHeader tcrsformats.RspHeader

		//원래 매핑하던 포멧에는 AGE가 포함하지 않는 이슈로 LGUPRSPUserInfo로 변경
		var tcrsRspBody formats.LGUPRSPUserInfo

		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody

	}

	return retData

}

func RestfulSendData(url string, inData []byte) []byte {
	// TCRS의 KT API 안정성 이슈로 응답이 돌아오지 않을 때가 있어 5초 타임아웃이 있는 HTTP 클라이언트 사용
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	reqData := bytes.NewBuffer(inData)
	// http.Post 대신 client.Post 사용
	resp, err := client.Post(url, "application/json", reqData)
	if err != nil {
		fmt.Println("RestfulSendData error: ", err.Error())
	}
	var f []byte
	if resp != nil {

		f, _ = ioutil.ReadAll(resp.Body)
		resp.Body.Close()
	}

	return f

}
