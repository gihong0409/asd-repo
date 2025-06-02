package process

import (
	"git.datau.co.kr/ferrari/ferrari-common/notifyapi/notifyClient"
	"git.datau.co.kr/ferrari/ferrari-common/notifyapi/notifyformats"
	"sync"
	"time"

	"git.datau.co.kr/ferrari/ferrari-cba/factory"
	"git.datau.co.kr/ferrari/ferrari-cba/process/batch"
	"git.datau.co.kr/ferrari/ferrari-common/commonutils"
)

type CBAProcess struct {
	// mutex sync.Mutex
	Fac         *factory.Factory
	sktProcess  batch.SKTProcess
	ktProcess   batch.KTProcess
	lgupProcess batch.LGUPProcess
}

func (_self *CBAProcess) Initialize(fac *factory.Factory) {
	_self.Fac = fac
	_self.sktProcess = batch.SKTProcess{Fac: fac}
	_self.ktProcess = batch.KTProcess{Fac: fac}
	_self.lgupProcess = batch.LGUPProcess{Fac: fac}
}

func (_self *CBAProcess) Processing() {
	weekDay := int(time.Now().Weekday())
	requestID := commonutils.GenTransID()

	if _self.Fac.Propertys().Slack.BatchStartCheckFlag {
		_ = notifyClient.SendEarthNotify(notifyformats.NotifySlackParams{
			CmdType:   notifyformats.SLACK_MSG_BYPASS,
			ChannelID: _self.Fac.Propertys().Slack.BatchStartCheckChannel,
			Msg:       "해지자 확인 배치",
			Attachments: []notifyformats.Attachments{
				{
					Color: "#36a64f",
					Blocks: []notifyformats.Blocks{
						{
							Type: "section",
							Text: &notifyformats.Text{Type: "mrkdwn", Text: "정상작동 "},
						},
						{
							Type: "section",
							Text: &notifyformats.Text{Type: "mrkdwn", Text: "ferrari-cba "},
						},
					},
				},
			},
			ConfigSet: _self.Fac.AppMode,
		})
	}

	var wg sync.WaitGroup
	wgCount := 0
	if _self.Fac.Propertys().SKTRunDay == weekDay {
		if _self.Fac.Propertys().SKTProcess {
			wgCount++
		}
	}
	if _self.Fac.Propertys().KTProcess {
		wgCount++
	}
	if _self.Fac.Propertys().LGUPProcess {
		wgCount++
	}

	wg.Add(wgCount)
	if _self.Fac.Propertys().SKTProcess && _self.Fac.Propertys().SKTRunDay == weekDay {
		go func() {
			defer wg.Done()
			_self.sktProcess.Process(requestID)
		}()
	}
	if _self.Fac.Propertys().KTProcess {
		go func() {
			defer wg.Done()
			_self.ktProcess.Process(requestID)
		}()
	}
	if _self.Fac.Propertys().LGUPProcess {
		go func() {
			defer wg.Done()
			_self.lgupProcess.Process(requestID)
		}()
	}
	wg.Wait()

}
