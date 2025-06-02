package batch

import (
	"ASD/factory"
	"ASD/utils"
	"fmt"
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
	AsdMember := []dmrsformats.AsdMember{} //CBA에서 사용하던 포멧 활용

	var loop bool = true

	for loop {
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
			loop = false

		}

		for i := range AsdMember {
			data := _self.Fac.GetMemberInfo(_self.Fac.TargetTCRSUrl, "0", AsdMember[i].PNumber)

			time.Sleep(1 * time.Second)
			AsdMember[i].Age = utils.ExtBD(data, AsdMember[i].Telecom)

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

		loop = false

	}

	fmt.Println(AsdMember)

}
