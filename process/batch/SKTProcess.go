package batch

import (
	"time"
		"strconv"

	"git.datau.co.kr/earth/earth-asd/factory"
	"git.datau.co.kr/earth/earth-asd/formats"
	"git.datau.co.kr/earth/earth-asd/utils"
	"github.com/datauniverse-lab/earth-common/dmrsapi/dmrsclient"
	"github.com/datauniverse-lab/earth-common/dmrsapi/dmrsformats"
	"github.com/datauniverse-lab/tesla-common/dmsapi/dmsclient"
		"github.com/datauniverse-lab/tesla-common/dmsapi/dmsformats"
)

type SKTProcess struct {
	Fac *factory.Factory
}

func (_self *SKTProcess) Process(requestID string) {

	for {
		asdMember := []formats.AsdMember{}

		var maxMemberList = _self.Fac.Propertys().MaxMemberList
        var dmrsInfo  dmrsformats.DMRSInfo  
		var tcrsUrl string

		if _self.Fac.Property.BenzProcess {
			dmrsInfo = _self.Fac.Propertys().BENZ.DMRSINFO
			tcrsUrl = _self.Fac.Propertys().BENZ.TcrsURL

		} else if _self.Fac.Property.BentleyProcess {
			dmrsInfo = _self.Fac.Propertys().BENTLEY.DMRSINFO
			tcrsUrl = _self.Fac.Propertys().BENTLEY.TcrsURL

		} else if  _self.Fac.Property.SaturnProcess {
			dmrsInfo = _self.Fac.Propertys().SATURN.DMRSINFO
			tcrsUrl = _self.Fac.Propertys().SATURN.TcrsURL

		} else if _self.Fac.Property.FerrariProcess {
			dmrsInfo = _self.Fac.Propertys().FERRARI.DMRSINFO
			tcrsUrl = _self.Fac.Propertys().FERRARI.TcrsURL

		} else if _self.Fac.Property.TeslaProcess {
			dmrsInfo = _self.Fac.Propertys().TESLA.DMSINFO
			tcrsUrl = _self.Fac.Propertys().TESLA.TcrsURL

		}



		if _self.Fac.Property.BenzProcess {
			utils.ReturnBenzAsdMembers(requestID, dmrsInfo, 0, &asdMember,maxMemberList)

		} else if _self.Fac.Property.BentleyProcess ||  _self.Fac.Property.SaturnProcess ||  _self.Fac.Property.FerrariProcess  {
			dmrsclient.DBMCall(
				dmrsInfo,
				dmrsformats.SELECTQUERY,
				"SelectAsdMember",
				[]interface{}{
					0,
					maxMemberList,
				},
				&asdMember,
				requestID,
			)

		}  else if _self.Fac.Property.TeslaProcess {
			dmsclient.DMSCall(
				_self.Fac.Dmscall,
				dmrsInfo,
				dmsformats.SELECTQUERY,
				"SelectAsdMember",
				[]string{
					"0",
					strconv.Itoa(maxMemberList),
				},
				nil,
				&asdMember,
				requestID,
			)

		}

		if len(asdMember) == 0 {
			_self.Fac.Print("***************** SKT 조회 종료 *****************")

			return
		}

		for i := range asdMember {
			time.Sleep(time.Duration(_self.Fac.Propertys().DelaySecSKT) * time.Millisecond)

			data := utils.GetMemberInfoTCRS(tcrsUrl, "0", asdMember[i].PNumber)

			asdMember[i].Age = utils.ExtractAge(data, 0)

			if _self.Fac.Property.BenzProcess {
				utils.UpdateAge(requestID, dmrsInfo, asdMember[i].PNumber, asdMember[i].Age)
			} else if  _self.Fac.Property.BentleyProcess ||  _self.Fac.Property.SaturnProcess ||  _self.Fac.Property.FerrariProcess {
				dmrsclient.DBMCall(
					dmrsInfo,
					dmrsformats.CUDQUERY,
					"UpdateAgeCheck",
					[]interface{}{
						asdMember[i].Age,
						asdMember[i].PNumber,
					},
					nil,
					requestID,
				)
			} else if _self.Fac.Property.TeslaProcess {
				dmrsclient.DMSCall(
					dmrsInfo,
					dmrsformats.CUDQUERY,
					"UpdateAgeCheck",
					[]string{
						strconv.Itoa(asdMember[i].Age),
						asdMember[i].PNumber,
					},
					nil,
					requestID,
				)
			}
		}
	}
}
