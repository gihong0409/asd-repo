package notifyClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git.datau.co.kr/ferrari/ferrari-common/notifyapi/notifyformats"
	"io"
	"net/http"
	"time"
)

func SlackOAuthSendData(url string, bearer string, inData []byte) (map[string]interface{}, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(inData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+bearer)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(bodyBytes))

	var result map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func SendEarthNotify(params notifyformats.NotifySlackParams) error {
	var notifyURL string

	if params.ConfigSet == "LIVE" {
		notifyURL = "https://notify.datau.co.kr/"
	} else {
		notifyURL = "https://dev-notify.datau.co.kr/"
	}

	// 공통 바디
	reqHeader := notifyformats.ReqHeader{CmdType: params.CmdType}
	reqBodyCommon := notifyformats.ReqBodySMCommon{
		ServiceName: params.ServiceName,
		ServerName:  params.ServerName,
		Platform:    params.PlatformName,
		Msg:         params.Msg,
	}

	// CmdType에 따라 분기
	var reqBody interface{}
	switch params.CmdType {
	case notifyformats.DB_INSERT:
		reqBody = notifyformats.ReqBodyDBInsert{
			ReqBodySMCommon: reqBodyCommon,
			BatchType:       params.BatchType,
		}
	case notifyformats.SLACK_SEND:
		reqBody = notifyformats.ReqBodyWebhookNotification{
			ReqBodySMCommon: reqBodyCommon,
		}
	case notifyformats.SLACK_MSG_BYPASS:
		reqBody = notifyformats.ReqBodySlackMsgByPass{
			Msg:         params.Msg,
			ChannelID:   params.ChannelID,
			Attachments: params.Attachments,
		}
	default:
		return fmt.Errorf("unknown CmdType %d", params.CmdType)
	}

	combinedData := map[string]interface{}{
		"Header": reqHeader,
		"Body":   reqBody,
	}

	notifyData, err := json.Marshal(combinedData)
	if err != nil {
		return fmt.Errorf("error marshalling data: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, notifyURL, bytes.NewBuffer(notifyData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
