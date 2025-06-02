package mq

import "time"

type MQLogFormats struct {
	Sender     string `json:"Sender"`
	ServerName string `json:"ServerName"`
	RequestID  string `json:"RequestID"`
	Log        string `json:"Log"`
}

type MQTransFormats struct {
	Sender       string    `json:"Sender"`
	Receiver     string    `json:"Receiver"`
	CallAppName  string    `json:"CallAppName"`
	RequestID    string    `json:"RequestID"`
	TransType    int       `json:"TransType"`
	RequestTime  time.Time `json:"RequestTime"`
	RequiredTime int64     `json:"RequiredTime"`
}

type MQMMAFormats struct {
	Sender      string `json:"Sender"`
	RequestID   string `json:"RequestID"`
	HandlerName string `json:"HandlerName"`
}

type MQJDBAFormats struct {
	Sender    string `json:"Sender"`
	JobType   int    `json:"JobType"` // 1 Join 2 cancel
	RequestID string `json:"RequestID"`
	Telecom   string `json:"Telecom"`
}

type MQSDBAFormats struct {
	Sender      string `json:"Sender"`
	RequestID   string `json:"RequestID"`
	HandlerName string `json:"HandlerName"`
}

type MQRBAFormats struct {
	Sender    string `json:"Sender"`
	RequestID string `json:"RequestID"`
}

type MQNBAFormats struct {
	Sender      string `json:"Sender"`
	RequestID   string `json:"RequestID"`
	HandlerName string `json:"HandlerName"`
	PNumber     string `json:"PNumber"`
}