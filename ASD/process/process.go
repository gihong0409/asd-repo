package process

import (
	"ASD/factory"
	"ASD/process/batch"
	"log"
	"sync"
)

type ASDProcess struct {
	// mutex sync.Mutex
	Fac        *factory.Factory
	sktProcess batch.SKTProcess
	//ktProcess   batch.KTProcess
	//lgupProcess batch.LGUPProcess
}

func (_self *ASDProcess) Initialize(fac *factory.Factory) {
	_self.Fac = fac
	_self.sktProcess = batch.SKTProcess{Fac: fac}
	//_self.ktProcess = batch.KTProcess{Fac: fac}
	//_self.lgupProcess = batch.LGUPProcess{Fac: fac}
}

func (_self *ASDProcess) Processing() {

	log.Print("나이 조회 프로세스 실행")

	var wg sync.WaitGroup

	// 3개의 고루틴을 기다리도록 설정
	wg.Add(1)

	// SKT 나이 가져오기
	go func() {
		defer wg.Done() // 고루틴 완료 시  카운트 감소
	}()

	// KT 나이 가져오기
	//go func() {
	//	defer wg.Done()
	//	KtGetAgeAll(kt, TCRSurl)
	//}()
	//
	//// LG U+ 나이 가져오기
	//go func() {
	//	defer wg.Done()
	//	LGUPGetAgeAll(lgup, TCRSurl)
	//}()

	// 모든 고루틴이 완료될 때까지 대기
	wg.Wait()

}
