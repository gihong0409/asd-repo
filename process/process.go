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

func (_self *ASDProcess) Initialize(fac *factory.Factory) {
	_self.Fac = fac
	_self.sktProcess = batch.SKTProcess{Fac: fac}
	_self.ktProcess = batch.KTProcess{Fac: fac}
	_self.lgupProcess = batch.LGUPProcess{Fac: fac}
}

func (_self *ASDProcess) Processing() {
	requestID := utils.GenTransID()

	var service string

	switch {
	case _self.Fac.Property.BenzProcess:
		service = "Benz"
	case _self.Fac.Property.BentleyProcess:
		service = "Bentley"
	case _self.Fac.Property.SaturnProcess:
		service = "Saturn"
	case _self.Fac.Property.FerrariProcess:
		service = "Ferrari"
	case _self.Fac.Property.TeslaProcess:
		service = "Tesla"
	default:
		_self.Fac.Print("모든 프로세스가 비활성화 상태")
		select {} // 파드가 죽지않도록 대기
	}
	

	_self.Fac.Print("[",service,"] 나이 조회 프로세스 실행")

	var wg sync.WaitGroup

	if _self.Fac.Propertys().SKTProcess {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_self.sktProcess.Process(requestID)
		}()
	}

	if _self.Fac.Propertys().KTProcess {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_self.ktProcess.Process(requestID)
		}()
	}

	if _self.Fac.Propertys().LGUPProcess {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_self.lgupProcess.Process(requestID)
		}()
	}

	wg.Wait()
}
