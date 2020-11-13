package pushUtil

import (
	"pz-infra/baseUtil"
	"pz-infra/errorUtil"
	"pz-infra/httpUtil"
	. "pz-infra/logging"
)

const (
	POST = "POST"
	GET  = "GET"

	DEVICE_CLIENT    = 1
	APP_CLIENT       = 2
	SEMANTICS_CLIENT = 3

	//push url list
	BIND_RESULT_URI               = "/bind_result"
	UNBIND_RESULT_URI             = "/unbind_result"
	ACTION_URI                    = "/action"
	CLOTHES_URI                   = "/clothes"
	DEVICE_URI                    = "/change_device_info"
	DANCE_URL                     = "/dance"
	REPORT_SERVICE_URI            = "/report_service"
	REPORT_CLOTHES_URI            = "/report_clothes"
	REPORT_STATE_URI              = "/report_state"
	REPORT_ACTION_URI             = "/report_action"
	USER_FACE_URI                 = "/user_face"
	REPORT_VOLUME                 = "/md"
	BLUETOOTH_STATE_URI           = "/bluetooth_state"
	AGENT_REGISTER_URI            = "/agent_register"
	USER_INFO_URI                 = "/user-info"
	DELETE_USER_FACE_URI          = "/delete_user_face"
	WIFI_NAME_URI                 = "/wifi_name"
	USER_VIP_URI                  = "/user_vip"
	SMART_FURNITURE               = "/smart_furniture"
	UPGRADE_URI                   = "/upgrade"
	REPORT_FIRST_WAKE_UP_CALL_URI = "/report_first_wake_up_call"
	JD_ACCESS_TOKEN               = "/jd_access_token"
)

// Push Result To Client
type PushResultRequest struct {
	UserId     int         `json:"userId" description:"用户id"`
	ClientType int         `json:"clientType" description:"客户端设备类型，要向该用户的哪个设备上推送。1=机器，2=移动设备App"`
	RequestUri string      `json:"requestUri" description:"PUSH协议uri"`
	Body       interface{} `json:"body" description:"PUSH协议的BODY部分"`
	Expire     int         `json:"expire" description:"推送消息的过期时间"`
	Callback   string      `json:"callback" description:"回调地址"`
}

// Push Message To Clients
type PushMessageRequest struct {
	GowildIds  []int  `json:"gowildIds" description:"统一账户中心id列表"`
	ClientType int    `json:"clientType" description:"客户端设备类型，要向该用户的哪个设备上推送。1=机器，2=移动设备App"`
	Desc       string `json:"desc" description:"消息描述"`
	Message    string `json:"message" description:"消息内容"`
}

func ConvertClientTypeToStr(clientType int) string {
	if clientType == APP_CLIENT {
		return "app"
	} else {
		return "voice"
	}
}

func PushToClient(body interface{}, url string) error {
	respContent, err := httpUtil.PostJson(url, nil, body)
	if err != nil {
		Log.Error("Failed to Read response body from Access layer", WithError(err))
		return err
	}

	Log.Debug("Push to Client Success", With("resp", string(respContent)))
	return nil
}

func PushErrorToClient(pushResultRequest *PushResultRequest, err error, url string, corpus ...*baseUtil.CorpusData) error {
	switch errType := err.(type) {
	case *errorUtil.HError:
		if len(corpus) > 0 && corpus[0] != nil {
			pushResultRequest.Body = baseUtil.FailCorpusResponseWithCode(errType.ResCode, errType.Message, corpus[0])
		} else {
			pushResultRequest.Body = baseUtil.FailResponseWithCode(errType.ResCode, errType.Message)
		}
	default:
		if len(corpus) > 0 && corpus[0] != nil {
			pushResultRequest.Body = baseUtil.CorpusBaseResponse{ResCode: baseUtil.RESPONSE_STATUS_FAIL, Message: err.Error(), Corpus: *corpus[0]}
		} else {
			pushResultRequest.Body = baseUtil.BaseResponse{ResCode: baseUtil.RESPONSE_STATUS_FAIL, Message: err.Error()}
		}
	}

	if err := PushToClient(pushResultRequest, url); err != nil {
		return err
	}

	return nil
}

func PushResultToClient(pushResult *PushResultRequest, err error, url string, corpus ...*baseUtil.CorpusData) {
	if err != nil {
		if len(corpus) > 0 && corpus[0] != nil {
			if err1 := PushErrorToClient(pushResult, err, url, corpus[0]); err1 != nil {
				Log.Error("Push Error to Client while error occurs", With("RequestUri", pushResult.RequestUri), WithError(err1))
			}
		} else {
			if err1 := PushErrorToClient(pushResult, err, url); err1 != nil {
				Log.Error("Push Error to Client while error occurs", With("RequestUri", pushResult.RequestUri), WithError(err1))
			}
		}
	} else {
		if err1 := PushToClient(pushResult, url); err1 != nil {
			Log.Error("Push to Client while error occurs", With("RequestUri", pushResult.RequestUri), WithError(err1))
		}
	}
}
