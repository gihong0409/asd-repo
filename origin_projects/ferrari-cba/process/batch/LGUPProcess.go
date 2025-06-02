package batch

import (
	"encoding/base64"
	"fmt"
	"git.datau.co.kr/ferrari/ferrari-common/notifyapi/notifyClient"
	"git.datau.co.kr/ferrari/ferrari-common/notifyapi/notifyformats"
	"strconv"
	"sync"
	"time"

	"git.datau.co.kr/ferrari/ferrari-cba/factory"
	"git.datau.co.kr/ferrari/ferrari-common/commonutils"
	"git.datau.co.kr/ferrari/ferrari-common/dmrsapi/dmrsclient"
	"git.datau.co.kr/ferrari/ferrari-common/dmrsapi/dmrsformats"
	"git.datau.co.kr/ferrari/ferrari-common/mq"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats/lgupformats"
)

type LGUPProcess struct {
	Fac      *factory.Factory
	CheckCnt int
	DelCnt   int
	ErrCnt   int
	wg       *sync.WaitGroup
}

func (_self *LGUPProcess) modProcess(divisor int, remainder int, requestID string, memberCount int, mutex *sync.Mutex) {

	defer _self.wg.Done()

	loop := true

	for loop {
		searchList := []dmrsformats.CbaMember{}
		dmrsclient.DBMCall(
			_self.Fac.Propertys().DmrsInfo,
			dmrsformats.SELECTQUERY,
			"GetSearchMember_lgup_mod",
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

func (_self *LGUPProcess) checkMember(searchList []dmrsformats.CbaMember, requestID string, memberCount int, mutex *sync.Mutex) {

	for _, memberInfo := range searchList {
		mutex.Lock()
		_self.CheckCnt++
		mutex.Unlock()
		_self.Fac.Print(requestID, "LGUP Del Batch Search START : ", _self.CheckCnt, "["+memberInfo.PNumber+"]")
		dmrsclient.DBMCall(
			_self.Fac.Propertys().DmrsInfo,
			dmrsformats.CUDQUERY,
			"UpdateStatus_lgup",
			[]interface{}{1, memberInfo.PNumber},
			nil,
			requestID,
		)

		tcrsRSP := tcrsapi.CheckServiceJoinStatus(
			_self.Fac.Propertys().TCRSUrl,
			strconv.Itoa(memberInfo.Telecom),
			memberInfo.PNumber,
			"LGUP CBA",
			requestID,
			_self.Fac.Print,
		)
		rspBody := tcrsRSP["Body"].(lgupformats.RSPUserInfo)

		CbaCodeCount := []dmrsformats.CountCbaCode{}
		dmrsclient.DBMCall(
			_self.Fac.Propertys().DmrsInfo,
			dmrsformats.SELECTQUERY,
			"CheckCancelCode",
			[]interface{}{memberInfo.Telecom, rspBody.RESPCODE},
			&CbaCodeCount,
			requestID,
		)

		if CbaCodeCount[0].CountCbaCode == 0 {
			dmrsclient.DBMCall(
				_self.Fac.Propertys().DmrsInfo,
				dmrsformats.CUDQUERY,
				"InsertCancelCode",
				[]interface{}{rspBody.RESPCODE, memberInfo.Telecom, rspBody.RESPMSG, _self.Fac.Propertys().AppName},
				nil,
				requestID,
			)
		}

		if rspBody.RESPCODE == 0 { //통신사 LG, 부가서비스 조회 가능
			if rspBody.SVC_AUTH_DT == "0" { //서비스 해지 됨
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

			}
		} else if rspBody.RESPCODE == 70 ||
			rspBody.RESPCODE == 71 ||
			rspBody.RESPCODE == 76 { //71 : skt, 76 : kt 번호이동
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
			mutex.Lock()
			_self.ErrCnt++
			mutex.Unlock()
		}
		dmrsclient.DBMCall(
			_self.Fac.Propertys().DmrsInfo,
			dmrsformats.CUDQUERY,
			"UpdateStatus_lgup",
			[]interface{}{2, memberInfo.PNumber},
			nil,
			requestID,
		)
		_self.Fac.Print(
			requestID,
			"LGUP Del Batch Search END : [Check:", _self.CheckCnt, "]/ [DEL:", _self.DelCnt, "]/[TOT:", memberCount, "]["+memberInfo.PNumber+"]",
		)
		time.Sleep(time.Duration(_self.Fac.Propertys().DelaySecLGUP) * time.Millisecond)
	}

}

// Process main Process
func (_self *LGUPProcess) Process(requestID string) {

	_self.Fac.Print(requestID, "LGUP Del Batch Process Started")

	//lgup 전체 가입자 수 카운트
	memberCount := []dmrsformats.CountMember{}
	dmrsclient.DBMCall(
		_self.Fac.Propertys().DmrsInfo,
		dmrsformats.SELECTQUERY,
		"GetMemberCount_lgup",
		nil,
		&memberCount,
		requestID,
	)

	if len(memberCount) > 0 {
		startTime := time.Now()

		_self.Fac.Print(requestID, "LGUP TOTAL COUNT: ", memberCount[0].CountMember)
		msg := fmt.Sprintf("[휴대폰분실보호] [LGUP 해지자 확인 시작] 대상 [ %d ]", memberCount[0].CountMember)

		_ = notifyClient.SendEarthNotify(notifyformats.NotifySlackParams{
			CmdType:      _self.Fac.Propertys().Slack.LGUPCmdType,
			ChannelID:    _self.Fac.Propertys().Slack.LGUPChannel,
			Msg:          msg,
			ServiceName:  notifyformats.Mfinder,
			ServerName:   "ferrari-cba",
			BatchType:    notifyformats.BatchTypeStart,
			PlatformName: "lgup",
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
			"[휴대폰분실보호] [LGUP 해지자 확인 완료] [에러 : %d] [해지자: %d] [총: %d] [소요시간:%s]",
			_self.ErrCnt,
			_self.DelCnt,
			memberCount[0].CountMember,
			timedu.String(),
		)

		_self.Fac.Print(requestID, msg)
		_ = notifyClient.SendEarthNotify(notifyformats.NotifySlackParams{
			CmdType:      _self.Fac.Propertys().Slack.LGUPCmdType,
			ChannelID:    _self.Fac.Propertys().Slack.LGUPChannel,
			Msg:          msg,
			ServiceName:  notifyformats.Mfinder,
			ServerName:   "ferrari-cba",
			BatchType:    notifyformats.BatchTypeFinsh,
			PlatformName: "lgup",
			ConfigSet:    _self.Fac.AppMode,
		})
	} else {
		_self.Fac.Print(requestID, "[휴대폰분실보호] [LGUP 해지자 배치] dmrs error")
	}

}
