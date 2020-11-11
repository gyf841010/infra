package fcmUtil

import (
	"errors"
	"github.com/NaySoftware/go-fcm"
	"github.com/astaxie/beego"
	. "pz_backend/infra/logging"
	"pz_backend/infra/timeUtil"
	"strconv"
	"time"
)

const (
	CONFIG_FCM_API_KEY = "fcm.apiKey"
	EVENT_TYPE_KEY     = "EVENT_TYPE"
	TIME_STAMP_KEY     = "TIME_STAMP"

	MAX_RETRY_TIMES = 3
)

func newFcmClient() *fcm.FcmClient {
	apiKey := beego.AppConfig.String(CONFIG_FCM_API_KEY)
	client := fcm.NewFcmClient(apiKey)
	//client.SetDelayWhileIdle(true)
	client.SetContentAvailable(true)
	client.SetPriorety(fcm.Priority_HIGH)
	return client
}

// Push Event with Custom Data
func PushEventData(receiverTokenList []string, eventType string, data map[string]string) (bool, int, error) {
	newClient := newFcmClient()
	if data == nil {
		data = make(map[string]string, 0)
	}
	data[EVENT_TYPE_KEY] = eventType
	data[TIME_STAMP_KEY] = strconv.Itoa(timeutil.CurrentUnixInt())
	//newClient.NewFcmMsgTo(receiverTokenList, data)
	newClient.NewFcmRegIdsMsg(receiverTokenList, data)
	fcmResponse, err := doPush(newClient, receiverTokenList)
	if err != nil {
		return false, 0, err
	}
	if !isPushSucceed(fcmResponse) {
		return false, 0, nil
	}
	Log.Debug("Succeed to Push Event Data", With("receiverToken", receiverTokenList), With("eventType", eventType), With("data", data))
	return true, fcmResponse.MsgId, nil
}

// Push Message with Notification Data, response
func PushMessageData(receiverTokenList []string, msgTitle string, msgBody string, clickAction string) (bool, int, error) {
	newClient := newFcmClient()
	notification := fcm.NotificationPayload{
		Title:       msgTitle,
		Body:        msgBody,
		ClickAction: clickAction,
		Sound:       "default",
	}
	newClient.AppendDevices(receiverTokenList)
	newClient.SetNotificationPayload(&notification)
	fcmResponse, err := doPush(newClient, receiverTokenList)
	if err != nil {
		return false, 0, err
	}
	if !isPushSucceed(fcmResponse) {
		return false, 0, nil
	}
	Log.Debug("Succeed to Push Message Data", With("receiverToken", receiverTokenList), With("msgTitle", msgTitle), With("msgBody", msgBody))
	return true, fcmResponse.MsgId, nil
}

// Push Event with Message, for FCM, send Notification Payload align with Data Payload
func PushEventWithMessage(receiverTokenList []string, eventType string, data map[string]string, msgTitle string, msgBody string, clickAction string) (bool, int, error) {
	newClient := newFcmClient()
	if data == nil {
		data = make(map[string]string, 0)
	}
	data[EVENT_TYPE_KEY] = eventType
	data[TIME_STAMP_KEY] = strconv.Itoa(timeutil.CurrentUnixInt())
	newClient.NewFcmRegIdsMsg(receiverTokenList, data)
	notification := fcm.NotificationPayload{
		Title:       msgTitle,
		Body:        msgBody,
		ClickAction: clickAction,
		Sound:       "default",
	}
	newClient.SetNotificationPayload(&notification)
	fcmResponse, err := doPush(newClient, receiverTokenList)
	if err != nil {
		return false, 0, err
	}
	if !isPushSucceed(fcmResponse) {
		return false, 0, nil
	}
	Log.Debug("Succeed to Push Event with Message", With("receiverToken", receiverTokenList), With("eventType", eventType), With("data", data))
	return true, fcmResponse.MsgId, nil
}

// Do Push Message, will perform retry in MAX_RETRY_TIMES
func doPush(client *fcm.FcmClient, receiverTokenList []string) (*fcm.FcmResponseStatus, error) {
	retryTimes := 0
	var responseStatus *fcm.FcmResponseStatus
	var err error
	for {
		if retryTimes > MAX_RETRY_TIMES {
			err = errors.New("Failed to Push Notification as Exceed Maximum Retry Times")
			return responseStatus, err
		}
		responseStatus, err = client.Send()
		if err != nil {
			Log.Error("Failed to Push Message Data", With("receiverTokenList", receiverTokenList), WithError(err))
			return responseStatus, err
		}
		if responseStatus.IsTimeout() {
			Log.Warn("responseStatus of FCM is Timeout.", With("responseStatus", responseStatus))
			duration, err := responseStatus.GetRetryAfterTime()
			if err != nil {
				Log.Error("Failed to parse retry after header", With("RetryAfter", responseStatus.RetryAfter))
				duration = time.Duration(2) * time.Second
			}
			select {
			case <-time.After(duration):
				retryTimes = retryTimes + 1
				Log.Info("Retry to Resend Message Data To FCM", With("RetryAfter", responseStatus.RetryAfter))
				continue
			}
		}
		Log.Debug("Submitted to Push Message Data", With("receiverTokenList", receiverTokenList), With("responseStatus", responseStatus))
		return responseStatus, nil
	}
}

func isPushSucceed(fcmResponse *fcm.FcmResponseStatus) bool {
	if fcmResponse == nil {
		return false
	}
	if !fcmResponse.Ok || fcmResponse.StatusCode != 200 {
		return false
	}
	for _, result := range fcmResponse.Results {
		if errMsg, found := result["error"]; found && errMsg != "" {
			Log.Error("fcmResponse is failed", With("error_result", result))
			return false
		}
	}
	if fcmResponse.Fail > 0 {
		Log.Error("fcmResponse is failed more than 0.", With("fail", fcmResponse.Fail), With("success", fcmResponse.Success))
		return false
	}
	return true
}
