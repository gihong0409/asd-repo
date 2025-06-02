package factory

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"git.datau.co.kr/ferrari/ferrari-common/dmrsapi/dmrsformats"
	"git.datau.co.kr/ferrari/ferrari-common/mq"
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
)

type Slack struct {
	BatchStartCheckFlag    bool   `json:"BatchStartCheckFlag"`
	BatchStartCheckChannel string `json:"BatchStartCheckChannel"`
	KTCmdType              int    `json:"KTCmdType"`
	KTChannel              string `json:"KTChannel"`
	LGUPCmdType            int    `json:"LGUPCmdType"`
	LGUPChannel            string `json:"LGUPChannel"`
	SKTCmdType             int    `json:"SKTCmdType"`
	SKTChannel             string `json:"SKTChannel"`
}

type Config struct {
	DmrsInfo      dmrsformats.DMRSInfo `json:"MIDDLECONF"`
	RedisAddr     string               `json:"RedisAddr"`
	Debug         bool                 `json:"Debug"`
	TCRSUrl       string               `json:"TCRSUrl"`
	AppName       string               `json:"AppName"`
	DelaySecSKT   int                  `json:"DelaySecSKT"`
	DelaySecKT    int                  `json:"DelaySecKT"`
	DelaySecLGUP  int                  `json:"DelaySecLGUP"`
	SKTRunDay     int                  `json:"SKTRunDay"`
	SKTProcess    bool                 `json:"SKTProcess"`
	KTProcess     bool                 `json:"KTProcess"`
	LGUPProcess   bool                 `json:"LGUPProcess"`
	NumGoRoutine  int                  `json:"NumGoRoutine"`
	MaxMemberList int                  `json:"MaxMemberList"`
	Slack         Slack                `json:"Slack"`
}

type Factory struct {
	logger         *logrus.Logger
	JSONConfigPath string
	property       Config
	ConfigMap      map[string]interface{}
	RedisClient    *redis.Client
	AppMode        string
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

	_self.RedisClient = redis.NewClient(&redis.Options{
		Addr:     _self.property.RedisAddr,
		Password: "",
		DB:       2,
	})

}

// Logger returns logger
func (_self *Factory) Logger() *logrus.Logger {
	return _self.logger
}

// Propertys returns config
func (_self *Factory) Propertys() Config {
	return _self.property
}

// Print log print
func (_self *Factory) Print(ReqID string, v ...interface{}) {

	_self.logger.WithFields(logrus.Fields{"RequestID": ReqID}).Println(v)
	hostname, _ := os.Hostname()
	msg := ""
	for _, item := range v {
		msg += fmt.Sprint(item)
	}

	mq.MQSendLogger(_self.RedisClient, mq.MQLogFormats{Sender: _self.property.DmrsInfo.APPNAME, ServerName: hostname, RequestID: ReqID, Log: msg})

}
