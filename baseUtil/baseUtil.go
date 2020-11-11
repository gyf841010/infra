package baseUtil

import (
	"io/ioutil"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"reflect"
	"encoding/json"
	"strings"
	"github.com/astaxie/beego/validation"
	"bytes"
	"compress/gzip"
	. "pz_backend/infra/errorUtil"
	. "pz_backend/infra/logging"
	"strconv"
	"net/http"
	"github.com/jinzhu/gorm"
)

const (
	STATUS_DELETE = 0
	STATUS_NORMAL = 1
	STATUS_CLEAR  = 2

	BIND   = 1
	UNBIND = 0

	DEFAULT_VIP_LEVEL       = 1
	DEFAULT_VIP_EXTEND_DAYS = 7                    //7天
	DEFAULT_VIP_DURATION    = 7 * 24 * 3600 * 1000 //7天毫秒

	NO_ACTIVATED_VIP = 0
	ACTIVATED_VIP    = 1

	RES_DEFAULT_ACTIVE = 1

	SN_UPPER_LIMIT_PER_MONTH = 100000 //每月生成sn上限数
)

const (
	RESPONSE_STATUS_FAIL    = 1
	RESPONSE_STATUS_SUCCESS = 0
)

const (
	DANCE_LV_1 = 1
	DANCE_LV_3 = 3
)

/***
http状态只包含200和500
200:app正常解析response
500:按照HError进行解析返回
*/
const (
	HTTP_STATUS_SUCCESS = 200
	HTTP_STATUS_FAIL    = 500
)

/***
HTTP请求头里面存放的sessionId.key和userId.key
*/
const (
	HTTP_HEADER_KEY_USER_ID      = "User-Id"
	HTTP_HEADER_KEY_GOWILD_ID    = "Gowild-Id"
	HTTP_HEADER_KEY_CLIENT_ID    = "Device-Id"
	HTTP_HEADER_KEY_CLIENT_TYPE  = "Client-Type"
	HTTP_HEADER_KEY_ACCESS_TOKEN = "Access-Token"
	HTTP_HEADER_KEY_IP           = "Client-Ip"
	SESSION_SID_PREFIX           = "holoera_"
	DEFAULT_SESSION_LIFETIME     = 3600 // one hour

	HTTP_HEADER_VERIFY = "verify"
)

//所有需要session的req数据结构都可以带上UserId属性，会自动从Header从填充
const (
	SESSION_ATTR_USER_ID      = "UserId"
	SESSION_ATTR_GOWILD_ID    = "UserId"
	SESSION_ATTR_CLIENT_ID    = "DeviceId"
	SESSION_ATTR_CLIENT_TYPE  = "ClientType"
	SESSION_ATTR_ACCESS_TOKEN = "AccessToken"
	SESSION_ATTR_IP           = "Ip"
)

const (
	DEFAULT_PAGE_SIZE = 10
	MAX_PAGE_SIZE     = 100
)

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func (p *Pagination) Fix() *Pagination {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = DEFAULT_PAGE_SIZE
	}
	if p.PageSize > MAX_PAGE_SIZE {
		p.PageSize = MAX_PAGE_SIZE
	}
	return p
}

func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.PageSize
}

func (p *Pagination) TotalPages(total int) int {
	totalPages := (total / p.PageSize)
	if total%p.PageSize != 0 {
		totalPages++
	}
	return totalPages
}

func (p *Pagination) Do(conn *gorm.DB) *gorm.DB {
	p = p.Fix()
	offset := p.Offset()
	conn = conn.Offset(offset).Limit(p.PageSize)
	Log.Debug("Pagination of offset", With("p", p))
	return conn
}

type BaseRequest struct {
	UserId      int    `json:"-" description:"用户id"`
	AccessToken string `json:"-" description:"access token"`
}

type BaseResponse struct {
	ResCode int    `json:"code" required:"true" description:"0成功 >1失败"`
	Message string `json:"message" description:"提示消息"`
}

type SingleCorpus struct {
	Text   string `json:"text" description:"回复语料文本"`
	Action string `json:"action" description:"回复语料动作"`
}

type CorpusData struct {
	Start SingleCorpus `json:"startCorpus" description:"开始语料"`
	Exec  SingleCorpus `json:"execCorpus" description:"执行语料"`
	End   SingleCorpus `json:"endCorpus" description:"结束语料"`
}

type CorpusBaseResponse struct {
	ResCode int        `json:"code" required:"true" description:"0成功 >1失败"`
	Message string     `json:"message" description:"提示消息"`
	Corpus  CorpusData `json:"corpus" description:"语料数据"`
}

type Md5Request struct {
	Md5 string `json:"md5" description:"上次返回数据的md5"`
}

func SuccessBaseResponse() BaseResponse {
	return BaseResponse{ResCode: RESPONSE_STATUS_SUCCESS, Message: ""}
}

func SuccessCorpusBaseResponse(corpus *CorpusData) CorpusBaseResponse {
	return CorpusBaseResponse{ResCode: RESPONSE_STATUS_SUCCESS, Message: "", Corpus: *corpus}
}

func FailResponse(message string) BaseResponse {
	return BaseResponse{ResCode: RESPONSE_STATUS_FAIL, Message: message}
}

func FailCorpusResponse(err error, corpus *CorpusData) CorpusBaseResponse {
	switch errType := err.(type) {
	case *HError:
		return FailCorpusResponseWithCode(errType.ResCode, errType.Message, corpus)
	default:
		return CorpusBaseResponse{ResCode: RESPONSE_STATUS_FAIL, Message: err.Error(), Corpus: *corpus}
	}
}

func FailResponseWithCode(errorCode int, message string) BaseResponse {
	return BaseResponse{ResCode: errorCode, Message: message}
}

func FailCorpusResponseWithCode(errorCode int, message string, corpus *CorpusData) CorpusBaseResponse {
	return CorpusBaseResponse{ResCode: errorCode, Message: message, Corpus: *corpus}
}

type RunHandlerOptions struct {
	// not unmarshall request body if request data is already unmarshalled
	NotUnmarshallRequestBody bool
	// do not auto render response entity
	// controller can do extra process response entity and call RespHandle explictly
	NotAutoRenderResponse bool
}

type ServiceMethod func(request interface{}) (response interface{}, err error)

type BaseController struct {
	beego.Controller
	SessionLifeTime int // in seconds
}

func (this *BaseController) ComposeServerError(err error) {
	switch errType := err.(type) {
	case *HError:
		this.Data["json"] = FailResponseWithCode(errType.ResCode, errType.Message)
	default:
		response := BaseResponse{ResCode: RESPONSE_STATUS_FAIL, Message: err.Error()}
		this.Data["json"] = response
	}
	this.ServeJSON()
}

func (this *BaseController) ComposeCorpusServerError(err error, corpus *CorpusData) {
	switch errType := err.(type) {
	case *HError:
		this.Data["json"] = FailCorpusResponseWithCode(errType.ResCode, errType.Message, corpus)
	default:
		response := CorpusBaseResponse{ResCode: RESPONSE_STATUS_FAIL, Message: err.Error(), Corpus: *corpus}
		this.Data["json"] = response
	}
	this.ServeJSON()
}

func (this *BaseController) ComposeLoginError(err error) {
	response := FailResponseWithCode(http.StatusUnauthorized, err.Error())
	this.Data["json"] = response
	this.ServeJSON()
}

func (c *BaseController) setStringField(v interface{}, fieldName string, fieldValue string) error {
	val := reflect.ValueOf(v)
	s := val.Elem()
	if s.Kind() == reflect.Struct {
		f := s.FieldByName(fieldName)
		if f.IsValid() {
			if f.CanSet() {
				if f.Kind() == reflect.String {
					f.SetString(fieldValue)
				}
			}
		}
	}
	return nil
}

func (c *BaseController) setIntField(v interface{}, fieldName string, fieldValue int) error {
	val := reflect.ValueOf(v)
	s := val.Elem()
	if s.Kind() == reflect.Struct {
		f := s.FieldByName(fieldName)
		if f.IsValid() {
			if f.CanSet() {
				if f.Kind() == reflect.Int {
					f.SetInt(int64(fieldValue))
				}
			}
		}
	}
	return nil
}

/***
验证accessToken
1.获取header中的client_id, access_token
2.根据client_id获取redis中的access_token
3.比对access_token
*/
func (c *BaseController) validateAccessToken(v interface{}) error {
	err := NewHErrorCustom(ERROR_CODE_ACCESS_TOKEN_ERROR)
	header := c.Ctx.Request.Header
	accessToken := header.Get(HTTP_HEADER_KEY_ACCESS_TOKEN)
	if "" == accessToken {
		Log.Debug("Validate Access Token Failed", With("accessToken", accessToken))
		return err
	}
	if v != nil {
		c.setAccessToken(v, accessToken)
	}

	return nil
}

/***
请求处理
转换为请求对象
*/
func (c *BaseController) reqHandle(v interface{}, option *RunHandlerOptions) error {
	if v == nil {
		return nil
	}
	val := reflect.ValueOf(v)
	s := val.Elem()
	if s.Type() == reflect.TypeOf(BaseRequest{}) || (option != nil && option.NotUnmarshallRequestBody) {
		return nil
	}
	//Check if need GZip Handling
	var requestBody = c.Ctx.Input.RequestBody
	if c.Ctx.Request.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(bytes.NewReader(requestBody))
		if err != nil {
			Log.Error("Failed Read Gzip Content of Request", WithError(err))
			return NewHErrorCustom(ERROR_CODE_PARAMETER_FORMAT_INVALID)
		}
		defer reader.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(reader)
		requestBody = buf.Bytes()
		Log.Debug("Succeed to Read Gzip Content of Request")
	}
	err := json.Unmarshal(requestBody, v)

	if err != nil {
		Log.Error("parse request json", With("request", string(c.Ctx.Input.RequestBody)), WithError(err))
		return NewHErrorCustom(ERROR_CODE_PARAMETER_FORMAT_INVALID)
	}
	return nil
}

/***
验证参数格式
*/
func (c *BaseController) formatHandle(v interface{}) error {
	if v == nil {
		return nil
	}

	//跳过基础类型的判断
	if strings.Index(reflect.TypeOf(v).String(), "string") >= 0 {
		return nil
	}

	valid := validation.Validation{}
	b, err := valid.RecursiveValid(v)
	if err != nil {
		Log.Error("formatHandle with error ", WithError(err))
		return NewHErrorCustom(ERROR_CODE_PARAMETER_FORMAT_INVALID)
	}
	if !b {
		var message string
		for _, err := range valid.Errors {
			message = err.Key + " " + err.Message
			Log.Debug("formatHandle valid ", With("err.Key", err.Key), With("err.Message", err.Message))
		}
		return NewHError(ERROR_CODE_PARAMETER_FORMAT_INVALID, message)
	}
	return nil
}

func (c *BaseController) setAccessToken(v interface{}, accessToken string) error {
	return c.setStringField(v, SESSION_ATTR_ACCESS_TOKEN, accessToken)
}

func (c *BaseController) handlerAccessToken(v interface{}) error {
	header := c.Ctx.Request.Header
	accessToken := header.Get(HTTP_HEADER_KEY_ACCESS_TOKEN)
	if strings.TrimSpace(accessToken) == "" {
		Log.Debug("AccessToken is blank")
		c.setAccessToken(v, "")
		return nil
	}
	c.setAccessToken(v, accessToken)

	return nil
}

func (c *BaseController) setIp(v interface{}, ip string) error {
	Log.Debug("Setting ip as ", With("ip", ip))
	return c.setStringField(v, SESSION_ATTR_IP, ip)
}

func (c *BaseController) handlerIp(v interface{}) error {
	header := c.Ctx.Request.Header
	ip := header.Get(HTTP_HEADER_KEY_IP)
	if strings.TrimSpace(ip) == "" {
		Log.Debug("Ip is blank")
		c.setIp(v, "")
		return nil
	}
	c.setIp(v, ip)

	return nil
}

func (c *BaseController) setClientId(v interface{}, clientId int) error {
	Log.Debug("Setting client id as ", With("clientId", clientId))
	return c.setIntField(v, SESSION_ATTR_CLIENT_ID, clientId)
}

func (c *BaseController) handlerClientId(v interface{}) error {
	header := c.Ctx.Request.Header
	clientId := header.Get(HTTP_HEADER_KEY_CLIENT_ID)
	if strings.TrimSpace(clientId) == "" {
		c.setClientId(v, 0)
		return nil
	}
	clientIdInt, err := strconv.Atoi(clientId)
	if err != nil {
		c.setClientId(v, 0)
		return nil
	}
	c.setClientId(v, clientIdInt)

	return nil
}

func (c *BaseController) setUserId(v interface{}, userId int) error {
	return c.setIntField(v, SESSION_ATTR_USER_ID, userId)
}

func (c *BaseController) handlerUserId(v interface{}) error {
	header := c.Ctx.Request.Header
	userId := header.Get(HTTP_HEADER_KEY_USER_ID)
	if strings.TrimSpace(userId) == "" {
		Log.Debug("userId is blank")
		c.setUserId(v, 0)
		return nil
	}
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		Log.Debug("userId is error")
		c.setUserId(v, 0)
		return nil
	}
	c.setUserId(v, userIdInt)

	return nil
}

func (c *BaseController) setGowildId(v interface{}, gowildId int) error {
	Log.Debug("Setting gowild id as ", With("gowildId", gowildId))
	return c.setIntField(v, SESSION_ATTR_GOWILD_ID, gowildId)
}

func (c *BaseController) handlerGowildId(v interface{}) error {
	header := c.Ctx.Request.Header
	gowildId := header.Get(HTTP_HEADER_KEY_GOWILD_ID)
	if strings.TrimSpace(gowildId) == "" {
		Log.Debug("gowildId is blank")
		c.setGowildId(v, 0)
		return nil
	}
	gowildIdInt, err := strconv.Atoi(gowildId)
	if err != nil {
		Log.Debug("gowildId is error")
		c.setGowildId(v, 0)
		return nil
	}
	c.setGowildId(v, gowildIdInt)

	return nil
}

func (c *BaseController) setClientType(v interface{}, clientType int) error {
	Log.Debug("Setting client Type as ", With("clientType", clientType))
	return c.setIntField(v, SESSION_ATTR_CLIENT_TYPE, clientType)
}

func (c *BaseController) handlerClientType(v interface{}) error {
	header := c.Ctx.Request.Header
	clientType := header.Get(HTTP_HEADER_KEY_CLIENT_TYPE)
	if strings.TrimSpace(clientType) == "" {
		Log.Debug("clientType is blank")
		c.setClientType(v, 0)
		return nil
	}
	clientTypeInt, err := strconv.Atoi(clientType)
	if err != nil {
		Log.Debug("clientType is error")
		c.setClientType(v, 0)
		return nil
	}
	c.setClientType(v, clientTypeInt)

	return nil
}

/***
验证Locale相关信息
1.获取header中的 通用公共字段
*/
func (c *BaseController) handleLocaleInfo(v interface{}) error {
	c.handlerUserId(v)
	return nil
}

/***
Controller Common 处理方法
@param need_handle_locate_info 是否需要处理常用头字段
@param need_validate_token 是否需要验证登录权限
@param req 请求参数封装
@param resp 返回结果封装
@param method 服务的Service方法
1.解析参数
2.验证参数格式
3.access token验证
4.Service方法调用
5.统一返回
*/
func (c *BaseController) RequestHandle(need_handle_locate_info, need_validate_token bool, req interface{}, options ...*RunHandlerOptions) (err error) {
	var option *RunHandlerOptions = nil
	Log.Debug("gongyaofei request json", With("request", string(c.Ctx.Input.RequestBody)))

	if len(options) > 0 {
		option = options[0]
	}

	for {
		err = c.reqHandle(req, option)
		if err != nil {
			break
		}

		err = c.formatHandle(req)
		if err != nil {
			break
		}
		if need_handle_locate_info {
			err = c.handleLocaleInfo(req)
			if err != nil {
				Log.Error("Error Handle Locale Info", With("request", req))
			}
		}

		if need_validate_token {
			err = c.validateAccessToken(req)
			if err != nil {
				break
			}
		}
		break
	}
	return
}

/***
响应处理
只返回两中HTTP状态，返回不同的数据结构
*/
func (c *BaseController) ResponseHandle(resp interface{}, err error) {
	if err != nil {
		c.Ctx.Output.Status = HTTP_STATUS_FAIL
		if beego.BConfig.RunMode == "dev" {
			// show detail error message during development
			c.Data["json"] = FailResponse(err.Error())
		} else {
			// not show detail error message in product environment
			c.Data["json"] = FailResponse("internal server error")
		}
		Log.Error("response handling error occurs", With("json", c.Data["json"]), WithError(err))
	} else {
		c.Ctx.Output.Status = HTTP_STATUS_SUCCESS
		c.Data["json"] = resp
		Log.Debug("response for request Handling", With("resp", resp))
	}
	c.ServeJSON()
}

func (c *BaseController) GetBaseUrl() string {
	apiUrl := beego.AppConfig.String("apiUrl")
	Log.Debug("Got YAPI GetBaseUrl ", With("yapiUrl", apiUrl))
	if strings.TrimSpace(apiUrl) == "" {
		Log.Warn("Got Empty YAPI GetBaseUrl")
	}
	return apiUrl
}

/**
判断 request 来源是否正常
*/
func (c *BaseController) VerifyRequest() bool {
	header := c.Ctx.Request.Header
	if header == nil || len(header) == 0 {
		Log.Debug("request header is empty or nil")
		return false
	}
	verify := header.Get(HTTP_HEADER_VERIFY)
	if len(verify) == 0 {
		Log.Debug("verify is not exist or blank")
		return false
	}

	requestBody, _ := ioutil.ReadAll(c.Ctx.Request.Body)
	c.Ctx.Request.Body.Close()
	//body := string(requestBody[:])
	uri := []byte(c.Ctx.Request.URL.RequestURI())
	h := md5.New()
	h.Write(append(requestBody, uri...))
	verifyHash := hex.EncodeToString(h.Sum(nil))
	Log.Debug("VerifyRequest result: ", With("verifyHash", verifyHash), With("verify", verify))

	return verifyHash == verify
}
