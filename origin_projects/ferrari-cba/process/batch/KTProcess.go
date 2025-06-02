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
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats/ktformats"
)

type KTProcess struct {
	Fac      *factory.Factory
	CheckCnt int
	DelCnt   int
	ErrCnt   int
	wg       *sync.WaitGroup
}

func (_self *KTProcess) modProcess(divisor int, remainder int, requestID string, memberCount int, mutex *sync.Mutex) {

	defer _self.wg.Done()

	loop := true

	for loop {

		searchList := []dmrsformats.CbaMember{}

		dmrsclient.DBMCall(
			_self.Fac.Propertys().DmrsInfo,
			dmrsformats.SELECTQUERY,
			"GetSearchMember_kt_mod",
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

func (_self *KTProcess) checkMember(searchList []dmrsformats.CbaMember, requestID string, memberCount int, mutex *sync.Mutex) {

	for _, memberInfo := range searchList {
		mutex.Lock()
		_self.CheckCnt++
		mutex.Unlock()

		_self.Fac.Print(requestID, "KT Del Batch Search START : ", _self.CheckCnt, "["+memberInfo.PNumber+"]")

		dmrsclient.DBMCall(
			_self.Fac.Propertys().DmrsInfo,
			dmrsformats.CUDQUERY,
			"UpdateStatus_kt",
			[]interface{}{1, memberInfo.PNumber},
			nil,
			requestID,
		)

		tcrsRSP := tcrsapi.CheckServiceJoinStatus(
			_self.Fac.Propertys().TCRSUrl,
			strconv.Itoa(memberInfo.Telecom),
			memberInfo.PNumber,
			"KT CBA",
			requestID,
			_self.Fac.Print,
		)
		rspBody := tcrsRSP["Body"].(ktformats.RSPBodyHeader)

		CbaCodeCount := []dmrsformats.CountCbaCode{}
		dmrsclient.DBMCall(
			_self.Fac.Propertys().DmrsInfo,
			dmrsformats.SELECTQUERY,
			"CheckCancelCode",
			[]interface{}{memberInfo.Telecom, rspBody.ERRORDETAIL.Errorcode},
			&CbaCodeCount,
			requestID,
		)
		if CbaCodeCount[0].CountCbaCode == 0 {

			dmrsclient.DBMCall(
				_self.Fac.Propertys().DmrsInfo,
				dmrsformats.CUDQUERY,
				"InsertCancelCode",
				[]interface{}{
					rspBody.ERRORDETAIL.Errorcode,
					memberInfo.Telecom,
					rspBody.ERRORDETAIL.Errordescription,
					_self.Fac.Propertys().AppName,
				},
				nil,
				requestID,
			)

		}

		if rspBody.ReturnCode == "0" {

			if rspBody.ERRORDETAIL.Errorcode == "SANTA1022" || //정액제(SID) 서비스미가입고객(만 12 세이하)
				rspBody.ERRORDETAIL.Errorcode == "SANTA1024" || //정액제(SID) 서비스미가입고객(만 15 세이하)
				rspBody.ERRORDETAIL.Errorcode == "SANTA1026" || //정액제(SID) 서비스미가입고객(만 19 세이하)
				rspBody.ERRORDETAIL.Errorcode == "SANTA1028" || //정액제(SID) 서비스미가입고객(만 20 세이상)
				rspBody.ERRORDETAIL.Errorcode == "SANTA1051" || //일시정지
				rspBody.ERRORDETAIL.Errorcode == "SANTA1053" || //법인 사업자
				rspBody.ERRORDETAIL.Errorcode == "SANTA1057" || //선불고객
				rspBody.ERRORDETAIL.Errorcode == "SANTA1065" || //ssn 오류
				rspBody.ERRORDETAIL.Errorcode == "SANTA1077" || // 유심단독상태 또는 타 통신사
				rspBody.ERRORDETAIL.Errorcode == "SANTA6075" { //필수 데이터 없음

				if rspBody.ERRORDETAIL.Errorcode == "SANTA1051" {
					member := []dmrsformats.SvcMember{}
					dmrsclient.DBMCall(
						_self.Fac.Propertys().DmrsInfo,
						dmrsformats.SELECTQUERY,
						"SelectMember",
						[]interface{}{memberInfo.PNumber},
						&member,
						requestID,
					)

					var pin string
					if len(member) > 0 {
						pin = member[0].Pin
					}

					dmrsclient.DBMCall(
						_self.Fac.Propertys().DmrsInfo,
						dmrsformats.CUDQUERY,
						"InsertPauseMember",
						[]interface{}{
							memberInfo.Telecom,
							memberInfo.PNumber,
							rspBody.ERRORDETAIL.Errorcode,
							rspBody.ERRORDETAIL.Errordescription,
							_self.Fac.Propertys().AppName,
							pin,
						},
						nil,
						requestID,
					)
				}

				_self.Fac.Print(requestID, "서비스 해지", memberInfo.PNumber)

				dmrsclient.DBMCall(
					_self.Fac.Propertys().DmrsInfo,
					dmrsformats.CUDQUERY,
					"Member_delete_kt",
					[]interface{}{3000, rspBody.ERRORDETAIL.Errordescription, memberInfo.PNumber},
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
							requestID)
						if header.ErrCode == dmrsformats.ErrorSuccess {
							mq.MQSendMMA(_self.Fac.RedisClient, mq.MQMMAFormats{Sender: "CBA", HandlerName: "FCM", RequestID: requestID})
							_self.Fac.Print(requestID, "PushMsg Insert", memberInfo.Telecom, memberInfo.PNumber)
						} else {
							_self.Fac.Print(requestID, "fail PushMsg Insert", memberInfo.Telecom, memberInfo.PNumber, header.ErrCode, header.ErrMsg)
						}
					}
				}

			} else {

				mutex.Lock()
				_self.ErrCnt++
				mutex.Unlock()
			}

		}

		dmrsclient.DBMCall(
			_self.Fac.Propertys().DmrsInfo,
			dmrsformats.CUDQUERY,
			"UpdateStatus_kt",
			[]interface{}{2, memberInfo.PNumber},
			nil,
			requestID,
		)
		_self.Fac.Print(
			requestID,
			"KT Del Batch Search END : [Check:", _self.CheckCnt, "]/ [DEL:", _self.DelCnt, "]/[TOT:", memberCount, "]["+memberInfo.PNumber+"]",
		)
		time.Sleep(time.Duration(_self.Fac.Propertys().DelaySecKT) * time.Millisecond)

	}

}

// Process ktProcess
func (_self *KTProcess) Process(requestID string) {

	//kt 전체 가입자 수 카운트
	memberCount := []dmrsformats.CountMember{}
	dmrsclient.DBMCall(
		_self.Fac.Propertys().DmrsInfo,
		dmrsformats.SELECTQUERY,
		"GetMemberCount_kt",
		nil,
		&memberCount,
		requestID,
	)

	if len(memberCount) > 0 {
		startTime := time.Now()

		_self.Fac.Print(requestID, "KT TOTAL COUNT: ", memberCount[0].CountMember)
		msg := fmt.Sprintf("[휴대폰분실보호] [KT 해지자 확인 시작] 대상 [ %d ]", memberCount[0].CountMember)

		_ = notifyClient.SendEarthNotify(notifyformats.NotifySlackParams{
			CmdType:      _self.Fac.Propertys().Slack.KTCmdType,
			ChannelID:    _self.Fac.Propertys().Slack.KTChannel,
			Msg:          msg,
			ServiceName:  notifyformats.Mfinder,
			ServerName:   "ferrari-cba",
			BatchType:    notifyformats.BatchTypeStart,
			PlatformName: "kt",
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
			"[휴대폰분실보호] [KT 해지자 확인 완료] [에러 : %d] [해지자 : %d] [총 : %d] [소요시간 : %s]",
			_self.ErrCnt,
			_self.DelCnt,
			memberCount[0].CountMember,
			timedu.String(),
		)

		_self.Fac.Print(requestID, msg)
		_ = notifyClient.SendEarthNotify(notifyformats.NotifySlackParams{
			CmdType:      _self.Fac.Propertys().Slack.KTCmdType,
			ChannelID:    _self.Fac.Propertys().Slack.KTChannel,
			Msg:          msg,
			ServiceName:  notifyformats.Mfinder,
			ServerName:   "ferrari-cba",
			BatchType:    notifyformats.BatchTypeFinsh,
			PlatformName: "kt",
			ConfigSet:    _self.Fac.AppMode,
		})

	} else {
		_self.Fac.Print(requestID, "[휴대폰분실보호] [KT 해지자 배치] dmrs error")
	}
}
