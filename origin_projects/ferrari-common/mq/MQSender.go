package mq

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v7"
)

func callPublish(redis *redis.Client, channel string, message string) bool {
	if err := redis.Publish(channel, message).Err(); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func MQSendLogger(redis *redis.Client, jsonSendData MQLogFormats) bool {

	sendData, err := json.Marshal(jsonSendData)
	if err != nil {
		return false
	}

	return callPublish(redis, "Logger", string(sendData))
}

func MQSendTrans(redis *redis.Client, jsonSendData MQTransFormats) bool {

	sendData, err := json.Marshal(jsonSendData)
	if err != nil {
		return false
	}

	return callPublish(redis, "Trans", string(sendData))
}

func MQSendMMA(redis *redis.Client, jsonSendData MQMMAFormats) bool {

	sendData, err := json.Marshal(jsonSendData)
	if err != nil {
		return false
	}
	return callPublish(redis, "MMA", string(sendData))
}

func MQSendJDBA(redis *redis.Client, jsonSendData MQJDBAFormats) bool {

	sendData, err := json.Marshal(jsonSendData)
	if err != nil {
		return false
	}
	return callPublish(redis, "JDBA", string(sendData))
}

func MQSendRBA(redis *redis.Client, jsonSendData MQRBAFormats) bool {

	sendData, err := json.Marshal(jsonSendData)
	if err != nil {
		return false
	}
	return callPublish(redis, "RBA", string(sendData))
}

func MQSendNBA(redis *redis.Client, jsonSendData MQNBAFormats) bool {

	sendData, err := json.Marshal(jsonSendData)
	if err != nil {
		return false
	}
	return callPublish(redis, "NBA", string(sendData))
}