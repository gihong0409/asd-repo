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

type LGUPProcess struct {
	Fac *factory.Factory
	wg  *sync.WaitGroup
}

func (_self *LGUPProcess) Process(requestID string) {
	AsdMember := []dmrsapi.AsdMember{} //Common 에 통합 전. 현재 프로젝트에 추가함

	for {
		dmrsclient.DBMCall(
			_self.Fac.Propertys().DmrsInfo,
			dmrsformats.SELECTQUERY,
			"SelectAsdMember",
			[]interface{}{
				2,
				_self.Fac.Propertys().MaxMemberList},
			&AsdMember,
			requestID,
		)

		if len(AsdMember) == 0 {
			return
		} else {

			for i := range AsdMember {
				data := utils.GetMemberInfoTCRS(_self.Fac.TargetTCRSUrl, "2", AsdMember[i].PNumber)

				time.Sleep(1 * time.Second)
				AsdMember[i].Age = utils.ExtBD(data, 2)

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
