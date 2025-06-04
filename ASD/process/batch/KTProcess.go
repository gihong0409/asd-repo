package batch

import (
	"ASD/dmrsapi"
	"ASD/factory"
	"ASD/formats"
	"ASD/utils"
	"git.datau.co.kr/ferrari/ferrari-common/dmrsapi/dmrsclient"
	"git.datau.co.kr/ferrari/ferrari-common/dmrsapi/dmrsformats"
	"sync"
	"time"
)

type KTProcess struct {
	Fac *factory.Factory
	wg  *sync.WaitGroup
}

func (_self *KTProcess) Process(requestID string) {
	AsdMember := []dmrsapi.AsdMember{} //Common 에 통합 전. 현재 프로젝트에 추가함

	for {
		dmrsheader := dmrsclient.DBMCall(
			_self.Fac.Propertys().DmrsInfo,
			dmrsformats.SELECTQUERY,
			"SelectAsdMember",
			[]interface{}{
				1,
				_self.Fac.Propertys().MaxMemberList},
			&AsdMember,
			requestID,
		)
		if dmrsheader.ErrCode != formats.Error {
			println("select err: ", dmrsheader.ErrCode, dmrsheader.ErrCode)
			return
		}
		if len(AsdMember) == 0 {
			return
		} else {
			for i, _ := range AsdMember {

				data := utils.GetMemberInfoTCRS(_self.Fac.TargetTCRSUrl, "1", AsdMember[i].PNumber)

				time.Sleep(1 * time.Second)
				AsdMember[i].Age = utils.ExtBD(data, 1)

				dmrsheader = dmrsclient.DBMCall(
					_self.Fac.Propertys().DmrsInfo,
					dmrsformats.CUDQUERY,
					"UpdateAgeCheck",
					[]interface{}{
						AsdMember[i].Age, AsdMember[i].PNumber,
					},
					nil,
					requestID,
				)
				if dmrsheader.ErrCode != formats.Error {
					continue
				}

			}
		}

	}

}
