package batch

import (
	"time"

	"git.datau.co.kr/earth/earth-asd/factory"
	"git.datau.co.kr/earth/earth-asd/formats"
	"git.datau.co.kr/earth/earth-asd/utils"
	"github.com/datauniverse-lab/earth-common/dmrsapi/dmrsclient"
	"github.com/datauniverse-lab/earth-common/dmrsapi/dmrsformats"
)

type KTProcess struct {
	Fac *factory.Factory
}

func (self *KTProcess) Process(requestID string) {

	for {
		asdMember := []formats.AsdMember{}

		if self.Fac.TargetService == 1 {
			utils.ReturnBenzAsdMembers(requestID, self.Fac.Propertys().DmrsInfo.DMRSURL, 1, &asdMember,
				self.Fac.Propertys().MaxMemberList)
		} else {
			dmrsclient.DBMCall(
				self.Fac.Propertys().DmrsInfo,
				dmrsformats.SELECTQUERY,
				"SelectAsdMember",
				[]interface{}{
					1,
					self.Fac.Propertys().MaxMemberList,
				},
				&asdMember,
				requestID,
			)
		}

		if len(asdMember) == 0 {
			self.Fac.Logger().Info("***************** KT 조회 종료 *****************")
			return
		}

		for i := range asdMember {
			time.Sleep(time.Duration(self.Fac.Propertys().DelaySecKT) * time.Millisecond)

			data := utils.GetMemberInfoTCRS(self.Fac.TargetTCRSUrl, "1", asdMember[i].PNumber)

			asdMember[i].Age = utils.ExtractAge(data, 1)

			if self.Fac.TargetService == 1 {
				utils.UpdateAge(requestID, self.Fac.Propertys().DmrsInfo.DMRSURL, asdMember[i].PNumber, asdMember[i].Age)
			} else {
				dmrsclient.DBMCall(
					self.Fac.Propertys().DmrsInfo,
					dmrsformats.CUDQUERY,
					"UpdateAgeCheck",
					[]interface{}{
						asdMember[i].Age,
						asdMember[i].PNumber,
					},
					nil,
					requestID,
				)
			}
		}
	}
}
