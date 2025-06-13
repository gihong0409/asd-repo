package process

import (
	"sync"

	"git.datau.co.kr/earth/earth-asd/factory"
	"git.datau.co.kr/earth/earth-asd/process/batch"
	"github.com/datauniverse-lab/earth-common/utils"
)

type ASDProcess struct {
	Fac         *factory.Factory
	sktProcess  batch.SKTProcess
	ktProcess   batch.KTProcess
	lgupProcess batch.LGUPProcess
}

func (self *ASDProcess) Initialize(fac *factory.Factory) {
	self.Fac = fac
	self.sktProcess = batch.SKTProcess{Fac: fac}
	self.ktProcess = batch.KTProcess{Fac: fac}
	self.lgupProcess = batch.LGUPProcess{Fac: fac}
}

func (self *ASDProcess) Processing() {
	requestID := utils.GenTransID()

	self.Fac.Print("***************** 나이 조회 프로세스 실행 *****************")

	var wg sync.WaitGroup

	if self.Fac.Propertys().SKTProcess {
		wg.Add(1)
		go func() {
			defer wg.Done()
			self.sktProcess.Process(requestID)
		}()
	}

	if self.Fac.Propertys().KTProcess {
		wg.Add(1)
		go func() {
			defer wg.Done()
			self.ktProcess.Process(requestID)
		}()
	}

	if self.Fac.Propertys().LGUPProcess {
		wg.Add(1)
		go func() {
			defer wg.Done()
			self.lgupProcess.Process(requestID)
		}()
	}

	wg.Wait()
}
