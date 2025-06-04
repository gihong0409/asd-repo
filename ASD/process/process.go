package process

import (
	"ASD/factory"
	"ASD/process/batch"
	"git.datau.co.kr/ferrari/ferrari-common/commonutils"
	"log"
	"sync"
)

type ASDProcess struct {
	// mutex sync.Mutex
	Fac         *factory.Factory
	sktProcess  batch.SKTProcess
	ktProcess   batch.KTProcess
	lgupProcess batch.LGUPProcess
}

func (_self *ASDProcess) Initialize(fac *factory.Factory) {
	_self.Fac = fac
	_self.sktProcess = batch.SKTProcess{Fac: fac}
	_self.ktProcess = batch.KTProcess{Fac: fac}
	_self.lgupProcess = batch.LGUPProcess{Fac: fac}
}

func (_self *ASDProcess) Processing() {
	requestID := commonutils.GenTransID()
	log.Print("나이 조회 프로세스 실행")

	var wg sync.WaitGroup
	wgCount := 0

	if _self.Fac.Propertys().SKTProcess {
		wgCount++
	}

	if _self.Fac.Propertys().KTProcess {
		wgCount++
	}
	if _self.Fac.Propertys().LGUPProcess {
		wgCount++
	}

	wg.Add(wgCount)

	// SKT 나이 가져오기
	go func() {
		defer wg.Done() // 고루틴 완료 시  카운트 감소
		_self.sktProcess.Process(requestID)

	}()
	go func() {
		defer wg.Done() // 고루틴 완료 시  카운트 감소
		_self.ktProcess.Process(requestID)

	}()
	go func() {
		defer wg.Done() // 고루틴 완료 시  카운트 감소
		_self.lgupProcess.Process(requestID)

	}()

	wg.Wait()

}
