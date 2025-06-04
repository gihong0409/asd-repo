package factory

import (
	"encoding/json"
	"fmt"
	"git.datau.co.kr/ferrari/ferrari-common/dmrsapi/dmrsformats"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"strconv"
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

	_self.logger.Level = logrus.DebugLevel
	_self.logger.Out = mw
	_self.logger.Formatter = customformatster

	_self.TargetService = _self.loadEnv(_self.property.ServiceNames)

	switch _self.TargetService {

	case 0:
		_self.property.DmrsInfo.DMRSURL = _self.property.BentleyDMRSUrl
		_self.TargetTCRSUrl = _self.property.BentleyTCRSUrl
	case 1:
		//_self.property.DmrsInfo.DMRSURL = _self.property.BenzDMRSUrl
		//_self.TargetTCRSUrl = _self.property.BenzTCRSUrl
		_self.Logger().Error("아직 추가되지 않은 서비스입니다.")

	case 2:
		_self.property.DmrsInfo.DMRSURL = _self.property.FerrariDMRSUrl
		_self.TargetTCRSUrl = _self.property.FerrariTCRSUrl
	case 3:
		_self.property.DmrsInfo.DMRSURL = _self.property.TeslaDMRSUrl
		_self.TargetTCRSUrl = _self.property.TeslaTCRSUrl
	case 4:
		//_self.property.DmrsInfo.DMRSURL = _self.property.MarsDMRSUrl
		//_self.TargetTCRSUrl = _self.property.MarsTCRSUrl
		_self.Logger().Error("아직 추가되지 않은 서비스입니다.")

	case 5:
		_self.property.DmrsInfo.DMRSURL = _self.property.SaturnDMRSUrl
		_self.TargetTCRSUrl = _self.property.SaturnTCRSUrl
	}

	println("target service: ", _self.TargetService, "\nDMRS URL: ", _self.property.DmrsInfo.DMRSURL,
		"\nTCRS URL: ", _self.TargetTCRSUrl)

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
		_self.Logger().Error("환경변수에 나이 추출을 실행할 서비스의 값을 true로 변경해주세요.")

		return -1
	}
	if trueCount > 1 {
		_self.Logger().Error("환경변수에 2개 이상의 서비스가 true로 설정되어 있습니다.")

		return -1
	}
	_self.Logger().Error("선택된 서비스: %s (index: %d)\n", names[targetService], targetService)

	return targetService
}
