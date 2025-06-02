package factory

import (
	"encoding/json"
	"fmt"
	"git.datau.co.kr/ferrari/ferrari-common/commonutils"
	"git.datau.co.kr/ferrari/ferrari-common/dmrsapi/dmrsformats"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats/ktformats"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats/lgupformats"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats/sktformats"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	DmrsInfo     dmrsformats.DMRSInfo `json:"MIDDLECONF"`
	Debug        bool                 `json:"Debug"`
	TCRSUrl      string               `json:"TCRSUrl"`
	AppName      string               `json:"AppName"`
	DelaySecSKT  int                  `json:"DelaySecSKT"`
	DelaySecKT   int                  `json:"DelaySecKT"`
	DelaySecLGUP int                  `json:"DelaySecLGUP"`

	SKTProcess    bool     `json:"SKTProcess"`
	KTProcess     bool     `json:"KTProcess"`
	LGUPProcess   bool     `json:"LGUPProcess"`
	NumGoRoutine  int      `json:"NumGoRoutine"`
	MaxMemberList int      `json:"MaxMemberList"`
	ServiceNames  []string `json:"ServiceNames"` // 서비스 이름 배열 추가

	BentleyDMRSUrl string `json:"BentleyDMRSUrl"`
	BenzDMRSUrl    string `json:"BenzDMRSUrl"`
	FerrariDMRSUrl string `json:"FerrariDMRSUrl"`
	TeslaDMRSUrl   string `json:"TeslaDMRSUrl"`
	MarsDMRSUrl    string `json:"MarsDMRSUrl"`
	SaturnDMRSUrl  string `json:"SaturnDMRSUrl"`

	BentleyTCRSUrl string `json:"BentleyTCRSUrl"`
	BenzTCRSUrl    string `json:"BenzTCRSUrl"`
	FerrariTCRSUrl string `json:"FerrariTCRSUrl"`
	TeslaTCRSUrl   string `json:"TeslaTCRSUrl"`
	MarsTCRSUrl    string `json:"MarsTCRSUrl"`
	SaturnTCRSUrl  string `json:"SaturnTCRSUrl"`
}

type Factory struct {
	logger         *logrus.Logger
	JSONConfigPath string
	property       Config
	ConfigMap      map[string]interface{}
	AppMode        string
	TargetService  int
	TargetTCRSUrl  string
}

func (_self *Factory) loadConfiguration(file string) {

	dat, _ := ioutil.ReadFile(file)
	fmt.Print(string(dat))

	json.Unmarshal(dat, &_self.property)

	_self.ConfigMap = make(map[string]interface{})
	json.Unmarshal(dat, &_self.ConfigMap)

}

// Initialize factory initialize
func (_self *Factory) Initialize() {

	fmt.Println(_self.JSONConfigPath + "config.json")

	_self.loadConfiguration(_self.JSONConfigPath + "config.json")

	fmt.Println(_self.property)

	mw := io.MultiWriter(os.Stdout)

	customformatster := new(logrus.TextFormatter)
	customformatster.TimestampFormat = "2006-01-02 15:04:05"

	_self.logger = logrus.New()

	// _self.logger.formatster = new(logrus.JSONformatster)

	//	_self.logger.formatster = new(logrus.Textformatster) //default
	// log.formatster.(*logrus.Textformatster).DisableTimestamp = true // remove timestamp from test output

	_self.logger.Level = logrus.DebugLevel
	_self.logger.Out = mw
	_self.logger.Formatter = customformatster

	_self.TargetService = _self.loadEnv(_self.property.ServiceNames)

	switch _self.TargetService {

	case 0:
		_self.property.DmrsInfo.DMRSURL = _self.property.BentleyDMRSUrl
		_self.TargetTCRSUrl = _self.property.BentleyTCRSUrl
	//case 1:
	//	_self.property.DmrsInfo.DMRSURL=_self.property.BenzDMRSUrl
	//	_self.TargetTCRSUrl = _self.property.BenzTCRSUrl
	case 2:
		_self.property.DmrsInfo.DMRSURL = _self.property.FerrariDMRSUrl
		_self.TargetTCRSUrl = _self.property.FerrariTCRSUrl
	case 3:
		_self.property.DmrsInfo.DMRSURL = _self.property.TeslaDMRSUrl
		_self.TargetTCRSUrl = _self.property.TeslaTCRSUrl
	//case 4:
	//	_self.property.DmrsInfo.DMRSURL=_self.property.MarsDMRSUrl
	//	_self.TargetTCRSUrl = _self.property.MarsTCRSUrl
	case 5:
		_self.property.DmrsInfo.DMRSURL = _self.property.SaturnDMRSUrl
		_self.TargetTCRSUrl = _self.property.SaturnTCRSUrl
	}

	println("target service: ", _self.TargetService, "dmrs and tcrs URL: ",
		_self.property.DmrsInfo.DMRSURL, _self.TargetTCRSUrl)

}

// Logger returns logger
func (_self *Factory) Logger() *logrus.Logger {
	return _self.logger
}

// Propertys returns config
func (_self *Factory) Propertys() Config {
	return _self.property
}
func (_self *Factory) loadEnv(services []string) int {

	names := services

	if len(names) == 0 {
		return -1
	}
	envService := make([]bool, len(names))
	for i, name := range names {
		envValue := os.Getenv(name)
		fmt.Printf("Env %s: %s\n", name, envValue)
		var err error
		if envService[i], err = strconv.ParseBool(envValue); err != nil {
			return -1
		}
		fmt.Printf("%s: %v\n", name, envService[i])
	}
	trueCount, targetService := 0, -1
	for i, val := range envService {
		if val {
			trueCount++
			targetService = i
		}
	}
	if trueCount == 0 {
		fmt.Println("스크립트에서 나이 추출을 실행할 서비스의 값을 true로 변경해주세요.")

		return -1
	}
	if trueCount > 1 {
		fmt.Println("스크립트에 2개 이상의 서비스가 true로 설정되어 있습니다.")
		return -1
	}
	fmt.Printf("선택된 서비스: %s (index: %d)\n", names[targetService], targetService)
	return targetService
}

func (_self *Factory) GetMemberInfo(TCWSURL string, telecom string, pnumber string) map[string]interface{} {
	var retData map[string]interface{}
	var teleName string

	if telecom == "0" || telecom == "1" || telecom == "2" {
		teleName = commonutils.TeleTypeNumToTelecomName(telecom)
	} else {
		teleName = telecom
	}

	println("telecom: ", telecom)
	println("teleName: ", teleName)
	cmdType := "USERINFO"

	if telecom == "1" {
		cmdType = "USERINFOANDKWAYS"
	}

	tcrsHeader := tcrsformats.ReqHeader{CmdType: cmdType}
	tcrsBody := tcrsformats.ReqBodyPNumber{PNumber: pnumber}

	tcrsRSP := commonutils.RestfulSendData(TCWSURL+teleName, commonutils.MakeJsonData(tcrsHeader, tcrsBody))

	retData = make(map[string]interface{})

	if strings.ToUpper(teleName) == "SKT" {
		var tcrsRspHeader tcrsformats.RspHeader
		var tcrsRspBody sktformats.RspMain
		var tcrsRspBodyDetail sktformats.UserInfoRsp
		tcrsRspBody.Body = &tcrsRspBodyDetail

		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)
		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody
		retData["BodyInfo"] = tcrsRspBodyDetail

	} else if strings.ToUpper(teleName) == "KT" {
		var tcrsRspHeader tcrsformats.RspHeader
		var tcrsRspBody ktformats.RSPUserInfoAndKways

		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody

	} else if strings.ToUpper(teleName) == "LGUP" {
		var tcrsRspHeader tcrsformats.RspHeader

		//원래 매핑하던 포멧에는 Age 포함하지 않는 이슈로 LGUPRSPUserInfo로 변경
		var tcrsRspBody lgupformats.RSPUserInfo

		commonutils.JsonToHaderBody([]byte(tcrsRSP), &tcrsRspHeader, &tcrsRspBody)

		retData["Header"] = tcrsRspHeader
		retData["Body"] = tcrsRspBody
	}

	return retData

}
