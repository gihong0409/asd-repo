package batch

import (
	"ASD/factory"
	"encoding/base64"
	"fmt"
	"git.datau.co.kr/ferrari/ferrari-common/notifyapi/notifyClient"
	"git.datau.co.kr/ferrari/ferrari-common/notifyapi/notifyformats"
	"strconv"
	"strings"
	"sync"
	"time"

	"git.datau.co.kr/ferrari/ferrari-common/commonutils"
	"git.datau.co.kr/ferrari/ferrari-common/dmrsapi/dmrsclient"
	"git.datau.co.kr/ferrari/ferrari-common/dmrsapi/dmrsformats"
	"git.datau.co.kr/ferrari/ferrari-common/mq"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats/sktformats"
)

type SKTProcess struct {
	Fac      *factory.Factory
	CheckCnt int
	DelCnt   int
	ErrCnt   int
	wg       *sync.WaitGroup
}

func (_self *SKTProcess) modProcess(divisor int, remainder int, requestID string, memberCount int, mutex *sync.Mutex) {

	defer _self.wg.Done()

	loop := true

	for loop {

		searchList := []dmrsformats.CbaMember{}

		dmrsclient.DBMCall(
			_self.Fac.Propertys().DmrsInfo,
			dmrsformats.SELECTQUERY,
			"GetSearchMember_skt_mod",
			[]interface{}{divisor, remainder, _self.Fac.Propertys().MaxMemberList},
			&searchList,
			requestID,
		)

		if len(searchList) > 0 {
			_self.checkMember(searchList, requestID, memberCount, mutex)
		} else {
			loop = false
		}
	}

}

func (_self *SKTProcess) checkMember(searchList []dmrsformats.CbaMember, requestID string, memberCount int, mutex *sync.Mutex) {

	for _, memberInfo := range searchList {
		mutex.Lock()
		_self.CheckCnt++
		mutex.Unlock()

		_self.Fac.Print(requestID, "SKT Del Batch Search START : ", _self.CheckCnt, "["+memberInfo.PNumber+"]")

		dmrsclient.DBMCall(
			_self.Fac.Propertys().DmrsInfo,
			dmrsformats.CUDQUERY,
			"UpdateStatus_skt",
			[]interface{}{1, memberInfo.PNumber},
			nil,
			requestID,
		)

		tcrsRSP := tcrsapi.CheckServiceJoinStatus(
			_self.Fac.Propertys().TCRSUrl,
			strconv.Itoa(memberInfo.Telecom),
			memberInfo.PNumber,
			"SKT CBA",
			requestID,
			_self.Fac.Print,
		)

		rspBody := tcrsRSP["Body"].(sktformats.RspMain)
		//RspBodyInfo := rspBody.Body.(*sktformats.JoinSearchRsp)

		if rspBody.Header.Result == "F" { // 비가입자 or 조회불가
			tcrsHeader := tcrsformats.ReqHeader{CmdType: "MEMBERSEARCH", RequestID: requestID, CallAppName: _self.Fac.Propertys().AppName}
			tcrsBody := tcrsformats.ReqBodyPNumber{PNumber: memberInfo.PNumber}
			_self.Fac.Print(requestID, "SKT SEND DATA", string(commonutils.MakeJsonData(tcrsHeader, tcrsBody)))

			tcrsRSP := commonutils.RestfulSendData(_self.Fac.Propertys().TCRSUrl+"SKT", commonutils.MakeJsonData(tcrsHeader, tcrsBody)) //사용자 조회
			_self.Fac.Print(requestID, "SKT RECV DATA", string(tcrsRSP))

			var tcrsRspHeader tcrsformats.RspHeader
			var tcrsRspBodyMemberSearch sktformats.RspMain
			var tcrsRspBodyMemberSearchDetail sktformats.UserInfoRsp
			tcrsRspBodyMemberSearch.Body = &tcrsRspBodyMemberSearchDetail
			commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBodyMemberSearch)

			if tcrsRspBodyMemberSearch.Header.Result == "S" && tcrsRspBodyMemberSearch.Header.ReturnCode == "00" { //비가입자
				_self.Fac.Print(requestID, "서비스 해지", memberInfo.PNumber)

				dmrsclient.DBMCall(
					_self.Fac.Propertys().DmrsInfo,
					dmrsformats.CUDQUERY,
					"Member_delete",
					[]interface{}{3000, memberInfo.PNumber},
					nil,
					requestID,
				)
				mutex.Lock()
				_self.DelCnt++
				mutex.Unlock()

				//PUSH
				Member := []dmrsformats.SvcMember{}
				dmrsclient.DBMCall(
					_self.Fac.Propertys().DmrsInfo,
					dmrsformats.SELECTQUERY,
					"GetDelMember",
					[]interface{}{memberInfo.PNumber, memberInfo.Telecom},
					&Member,
					requestID,
				)

				if len(Member) > 0 {
					if Member[0].PushToken != "" {
						PushData := dmrsformats.PushData{PushType: "9999", TransID: requestID}
						pushenc := base64.StdEncoding.EncodeToString(commonutils.ObjectToJson(PushData))

						//insert trans msg
						header := dmrsclient.DBMCall(
							_self.Fac.Propertys().DmrsInfo,
							dmrsformats.CUDQUERY,
							"TransMsg_Insert",
							[]interface{}{1000, memberInfo.Telecom, memberInfo.PNumber, memberInfo.PushToken, pushenc, "해지푸쉬", "해지완료", 0},
							nil,
							requestID,
						)
						if header.ErrCode == dmrsformats.ErrorSuccess {
							mq.MQSendMMA(_self.Fac.RedisClient, mq.MQMMAFormats{Sender: "CBA", HandlerName: "FCM", RequestID: requestID})
							_self.Fac.Print(requestID, "PushMsg Insert", memberInfo.Telecom, memberInfo.PNumber)
						} else {
							_self.Fac.Print(requestID, "fail PushMsg Insert", memberInfo.Telecom, memberInfo.PNumber, header.ErrCode, header.ErrMsg)
						}
					}
				}
			} else if tcrsRspBodyMemberSearch.Header.Result == "F" && tcrsRspBodyMemberSearch.Header.ReturnCode == "PCI_DTS_E3162" { // 통신사 조회 불가
				_self.Fac.Print(requestID, "서비스 해지 통신사 변경", memberInfo.PNumber)

				dmrsclient.DBMCall(
					_self.Fac.Propertys().DmrsInfo,
					dmrsformats.CUDQUERY,
					"Member_delete",
					[]interface{}{3001, memberInfo.PNumber},
					nil,
					requestID,
				)
				mutex.Lock()
				_self.DelCnt++
				mutex.Unlock()
			} else {
				idx := strings.Index(rspBody.Header.ReturnDesc, "|")
				eBodyCode := rspBody.Header.ReturnCode
				eBodyMsg := rspBody.Header.ReturnDesc
				if idx > -1 {
					eBodyCode += rspBody.Header.ReturnDesc[:]
					eBodyMsg = rspBody.Header.ReturnDesc[idx:]
				}
				CbaCodeCount := []dmrsformats.CountCbaCode{}

				dmrsclient.DBMCall(
					_self.Fac.Propertys().DmrsInfo,
					dmrsformats.SELECTQUERY,
					"CheckCancelCode",
					[]interface{}{memberInfo.Telecom, eBodyCode},
					&CbaCodeCount,
					requestID,
				)

				if CbaCodeCount[0].CountCbaCode == 0 {
					dmrsclient.DBMCall(
						_self.Fac.Propertys().DmrsInfo,
						dmrsformats.CUDQUERY,
						"InsertCancelCode",
						[]interface{}{eBodyCode, memberInfo.Telecom, eBodyMsg, _self.Fac.Propertys().AppName},
						nil,
						requestID,
					)
				}

				mutex.Lock()
				_self.ErrCnt++
				mutex.Unlock()
			}

		}

		dmrsclient.DBMCall(
			_self.Fac.Propertys().DmrsInfo,
			dmrsformats.CUDQUERY,
			"UpdateStatus_skt",
			[]interface{}{2, memberInfo.PNumber},
			nil,
			requestID,
		)

		_self.Fac.Print(
			requestID,
			"SKT Del Batch Search END : [Check:", _self.CheckCnt, "]/ [DEL:", _self.DelCnt, "]/[TOT:", memberCount, "]["+memberInfo.PNumber+"]",
		)
		time.Sleep(time.Duration(_self.Fac.Propertys().DelaySecSKT) * time.Millisecond)
	}
}

func (_self *SKTProcess) Process(requestID string) {

	//skt 전체 가입자 수 카운트
	memberCount := []dmrsformats.CountMember{}

	dmrsclient.DBMCall(
		_self.Fac.Propertys().DmrsInfo,
		dmrsformats.SELECTQUERY,
		"GetMemberCount_skt",
		nil,
		&memberCount,
		requestID,
	)

	if len(memberCount) > 0 {
		startTime := time.Now()

		_self.Fac.Print(requestID, "SKT TOTAL COUNT: ", memberCount[0].CountMember)
		msg := fmt.Sprintf("[휴대폰분실보호] [SKT 해지자 확인 시작] 대상 [ %d ]", memberCount[0].CountMember)

		_ = notifyClient.SendEarthNotify(notifyformats.NotifySlackParams{
			CmdType:      _self.Fac.Propertys().Slack.SKTCmdType,
			ChannelID:    _self.Fac.Propertys().Slack.SKTChannel,
			Msg:          msg,
			ServiceName:  notifyformats.Mfinder,
			ServerName:   "ferrari-cba",
			BatchType:    notifyformats.BatchTypeStart,
			PlatformName: "skt",
			ConfigSet:    _self.Fac.AppMode,
		})

		NumGoRoutine := _self.Fac.Propertys().NumGoRoutine

		_self.wg = &sync.WaitGroup{}
		_self.wg.Add(NumGoRoutine)

		var mutex = &sync.Mutex{}

		for remainder := 0; remainder < NumGoRoutine; remainder++ {
			go _self.modProcess(NumGoRoutine, remainder, requestID, memberCount[0].CountMember, mutex)
		}

		_self.wg.Wait()

		timedu := time.Since(startTime)

		msg = fmt.Sprintf(
			"[휴대폰분실보호] [SKT 해지자 확인 완료] [에러 : %d] [해지자: %d] [총: %d] [소요시간:%s]",
			_self.ErrCnt,
			_self.DelCnt,
			memberCount[0].CountMember,
			timedu.String(),
		)

		_self.Fac.Print(requestID, msg)
		_ = notifyClient.SendEarthNotify(notifyformats.NotifySlackParams{
			CmdType:      _self.Fac.Propertys().Slack.SKTCmdType,
			ChannelID:    _self.Fac.Propertys().Slack.SKTChannel,
			Msg:          msg,
			ServiceName:  notifyformats.Mfinder,
			ServerName:   "ferrari-cba",
			BatchType:    notifyformats.BatchTypeFinsh,
			PlatformName: "skt",
			ConfigSet:    _self.Fac.AppMode,
		})
	} else {
		_self.Fac.Print(requestID, "[휴대폰분실보호] [SKT 해지자 배치] dmrs error")
	}

}
