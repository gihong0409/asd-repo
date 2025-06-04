package batch

import (
	"ASD/dmrsapi"
	"ASD/factory"
	"ASD/utils"
	"git.datau.co.kr/ferrari/ferrari-common/dmrsapi/dmrsclient"
	"git.datau.co.kr/ferrari/ferrari-common/dmrsapi/dmrsformats"
	"sync"
	"time"
)

type SKTProcess struct {
	Fac *factory.Factory
	wg  *sync.WaitGroup
}

func (_self *SKTProcess) Process(requestID string) {

	println("requestID: ", requestID)
	for {

		AsdMember := []dmrsapi.AsdMember{} //Common 에 통합 전. 현재 프로젝트에 추가함

		println("flag: SKT")
		dmrsclient.DBMCall(
			_self.Fac.Propertys().DmrsInfo,
			dmrsformats.SELECTQUERY,
			"SelectAsdMember",
			[]interface{}{
				0,
				_self.Fac.Propertys().MaxMemberList},
			&AsdMember,
			requestID,
		)

		if len(AsdMember) == 0 {
			println("SKT   :  len(AsdMember) == 0")

			return
		} else {
			println("SKT   :", len(AsdMember))

			for i := range AsdMember {
				time.Sleep(1 * time.Second)
				println("Member Age Get Start: ", AsdMember[i].PNumber, "Telecom", AsdMember[i].Telecom)

				data := utils.GetMemberInfoTCRS(_self.Fac.TargetTCRSUrl, "0", AsdMember[i].PNumber)

				AsdMember[i].Age = utils.ExtBD(data, 0)

				println("SKT PNumber", AsdMember[i].PNumber, " age: ", AsdMember[i].Age)

				dmrsclient.DBMCall(
					_self.Fac.Propertys().DmrsInfo,
					dmrsformats.CUDQUERY,
					"UpdateAgeCheck",
					[]interface{}{
						AsdMember[i].Age, AsdMember[i].PNumber,
					},
					nil,
					requestID,
				)

			}

		}

	}
}
