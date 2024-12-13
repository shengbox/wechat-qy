package suite

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/url"

	"github.com/go-resty/resty/v2"
	crypter "github.com/heroicyang/wechat-crypter"
	"github.com/shengbox/wechat-qy/base"
)

// 应用套件相关操作的 API 地址
const (
	suiteTokenURI    = "https://qyapi.weixin.qq.com/cgi-bin/service/get_suite_token"
	preAuthCodeURI   = "https://qyapi.weixin.qq.com/cgi-bin/service/get_pre_auth_code"
	authURI          = "https://qy.weixin.qq.com/cgi-bin/loginpage"
	permanentCodeURI = "https://qyapi.weixin.qq.com/cgi-bin/service/get_permanent_code"
	authInfoURI      = "https://qyapi.weixin.qq.com/cgi-bin/service/get_auth_info"
	getAgentURI      = "https://qyapi.weixin.qq.com/cgi-bin/service/get_agent"
	setAgentURI      = "https://qyapi.weixin.qq.com/cgi-bin/service/set_agent"
	corpTokenURI     = "https://qyapi.weixin.qq.com/cgi-bin/service/get_corp_token"
	adminListURI     = "https://qyapi.weixin.qq.com/cgi-bin/service/get_admin_list"
	installURI       = "https://open.work.weixin.qq.com/3rdapp/install"
	registerCode     = "https://qyapi.weixin.qq.com/cgi-bin/service/get_register_code"
	registerURI      = "https://open.work.weixin.qq.com/3rdservice/wework/register"

	contactSyncSuccessURI = "https://qyapi.weixin.qq.com/cgi-bin/sync/contact_sync_success"
	getuserinfo3rdURI     = "https://qyapi.weixin.qq.com/cgi-bin/service/getuserinfo3rd"
	getuserinfo3rdAuthURI = "https://qyapi.weixin.qq.com/cgi-bin/service/auth/getuserinfo3rd"
	getuserdetail3rdURI   = "https://qyapi.weixin.qq.com/cgi-bin/service/getuserdetail3rd"
	getLoginInfoURI       = "https://qyapi.weixin.qq.com/cgi-bin/service/get_login_info"
	uploadURI             = "https://qyapi.weixin.qq.com/cgi-bin/service/media/upload"
	idTranslateURI        = "https://qyapi.weixin.qq.com/cgi-bin/service/contact/id_translate"
	getJobResultURI       = "https://qyapi.weixin.qq.com/cgi-bin/service/batch/getresult"
	setSessionInfoURI     = "https://qyapi.weixin.qq.com/cgi-bin/service/set_session_info"

	activeAccountURI       = "https://qyapi.weixin.qq.com/cgi-bin/license/active_account"
	getActiveInfoByUserURI = "https://qyapi.weixin.qq.com/cgi-bin/license/get_active_info_by_user"
	listActivedAccountURI  = "https://qyapi.weixin.qq.com/cgi-bin/license/list_actived_account"
	listOrderURI           = "https://qyapi.weixin.qq.com/cgi-bin/license/list_order"
	getOrderURI            = "https://qyapi.weixin.qq.com/cgi-bin/license/get_order"
	listOrderAccountURI    = "https://qyapi.weixin.qq.com/cgi-bin/license/list_order_account"
)

// Suite 结构体包含了应用套件的相关操作
type Suite struct {
	id             string
	secret         string
	ticket         string
	token          string
	encodingAESKey string
	msgCrypter     crypter.MessageCrypter
	tokener        *base.Tokener
	client         *base.Client
}

// New 方法用于创建 Suite 实例
func New(suiteID, suiteSecret, suiteToken, suiteEncodingAESKey string) *Suite {
	msgCrypter, _ := crypter.NewMessageCrypter(suiteToken, suiteEncodingAESKey, suiteID)

	suite := &Suite{
		id:             suiteID,
		secret:         suiteSecret,
		token:          suiteToken,
		encodingAESKey: suiteEncodingAESKey,
		msgCrypter:     msgCrypter,
	}

	suite.client = base.NewClient(suite)
	suite.tokener = base.NewTokener(suite)

	return suite
}

// Retriable 方法实现了套件在发起请求遇到 token 错误时，先刷新 token 然后再次发起请求的逻辑
func (s *Suite) Retriable(reqURL string, body []byte) (bool, string, error) {
	u, err := url.Parse(reqURL)
	if err != nil {
		return false, "", nil
	}

	q := u.Query()
	if q.Get("suite_access_token") == "" {
		return false, "", nil
	}

	result := &base.Error{}
	if err := json.Unmarshal(body, result); err != nil {
		return false, "", err
	}

	switch result.ErrCode {
	case base.ErrCodeOk:
		return false, "", nil
	case base.ErrCodeSuiteTokenInvalid, base.ErrCodeSuiteTokenTimeout, base.ErrCodeSuiteTokenFailure:
		if err := s.tokener.RefreshToken(); err != nil {
			return false, "", err
		}

		token, err := s.tokener.Token()
		if err != nil {
			return false, "", err
		}

		q.Set("suite_access_token", token)
		u.RawQuery = q.Encode()
		return true, u.String(), nil
	default:
		return false, "", result
	}
}

// Parse 方法用于解析应用套件的消息回调
func (s *Suite) Parse(body []byte, signature, timestamp, nonce string) (interface{}, error) {
	var err error

	reqBody := &base.RecvHTTPReqBody{}
	if err = xml.Unmarshal(body, reqBody); err != nil {
		return nil, err
	}

	if signature != s.msgCrypter.GetSignature(timestamp, nonce, reqBody.Encrypt) {
		return nil, fmt.Errorf("validate signature error")
	}

	origData, suiteID, err := s.msgCrypter.Decrypt(reqBody.Encrypt)
	if err != nil {
		return nil, err
	}

	probeData := &struct {
		InfoType string
		Event    string
	}{}

	if err = xml.Unmarshal(origData, probeData); err != nil {
		return nil, err
	}

	if suiteID != s.id {
		log.Printf("the request is from suite[%s], not from suite[%s] event=%s", suiteID, s.id, probeData.Event)
		// return nil, fmt.Errorf("the request is from suite[%s], not from suite[%s]", suiteID, s.id)
	}

	var data interface{}
	switch probeData.InfoType {
	case "suite_ticket":
		data = &RecvSuiteTicket{}
	case "change_auth", "cancel_auth":
		data = &RecvSuiteAuth{}
	case "create_auth":
		data = &RecvCreateAuth{}
	case "register_corp":
		data = &RecRegisterCorp{}
	case "change_contact":
		data = &RecvChangeContactEvent{}
	case "change_external_tag":
		data = &RecvChangeExternalTagEvent{}
	case "change_external_contact":
		data = &RecvChangeExternalContactEvent{}
	case "change_external_chat":
		data = &RecvChangeExternalChatEvent{}
	case "auto_activate":
		data = &RecvAutoActivateEvent{}
	case "license_pay_success":
		data = &LicensePaySuccess{}
	default:
		switch probeData.Event {
		case "change_app_admin", "subscribe", "enter_agent", "unsubscribe":
			data = &RecvChangeEvent{}
		case "sys_approval_change":
			data = &SysApprovalChangeEvent{}
		case "customer_acquisition":
			data = &RecvCustomerAcquisitionEvent{}
		case "program_notify":
			data = &ProgramNotifyEvent{}
		default:
			return nil, fmt.Errorf("unknown message type: %s origData: %s", probeData.InfoType, origData)
		}
	}

	if err = xml.Unmarshal(origData, data); err != nil {
		return nil, err
	}

	return data, nil
}

// Response 方法用于生成应用套件的被动响应消息
func (s *Suite) Response(message []byte) ([]byte, error) {
	msgEncrypt, err := s.msgCrypter.Encrypt(string(message))
	if err != nil {
		return nil, err
	}

	nonce := base.GenerateNonce()
	timestamp := base.GenerateTimestamp()
	signature := s.msgCrypter.GetSignature(fmt.Sprintf("%d", timestamp), nonce, msgEncrypt)

	resp := &base.RecvHTTPRespBody{
		Encrypt:      base.StringToCDATA(msgEncrypt),
		MsgSignature: base.StringToCDATA(signature),
		TimeStamp:    timestamp,
		Nonce:        base.StringToCDATA(nonce),
	}

	return xml.MarshalIndent(resp, " ", "  ")
}

// SetTicket 方法用于设置套件的 ticket 信息
func (s *Suite) SetTicket(suiteTicket string) {
	s.ticket = suiteTicket
}

// FetchToken 方法用于向 API 服务器获取套件的令牌信息
func (s *Suite) FetchToken() (token string, expiresIn int64, err error) {
	buf, _ := json.Marshal(map[string]string{
		"suite_id":     s.id,
		"suite_secret": s.secret,
		"suite_ticket": s.ticket,
	})

	body, err := s.client.PostJSON(suiteTokenURI, buf)
	if err != nil {
		return
	}

	tokenInfo := &struct {
		Token     string `json:"suite_access_token"`
		ExpiresIn int64  `json:"expires_in"`
	}{}

	if err = json.Unmarshal(body, tokenInfo); err != nil {
		return
	}

	token = tokenInfo.Token
	expiresIn = tokenInfo.ExpiresIn

	if token == "" {
		log.Println(string(body))
	}

	return
}

func (s *Suite) getPreAuthCode(appIDs []int) (*preAuthCodeInfo, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}

	qs := url.Values{}
	qs.Add("suite_access_token", token)
	uri := preAuthCodeURI + "?" + qs.Encode()

	buf, _ := json.Marshal(map[string]interface{}{
		"suite_id": s.id,
		"appid":    appIDs,
	})

	body, err := s.client.PostJSON(uri, buf)
	if err != nil {
		return nil, err
	}

	result := &preAuthCodeInfo{}
	err = json.Unmarshal(body, result)

	return result, err
}

// GetAuthURI 方法用于获取应用套件的授权地址
func (s *Suite) GetAuthURI(appIDs []int, redirectURI, state string) (string, error) {
	preAuthCodeInfo, err := s.getPreAuthCode(appIDs)
	if err != nil {
		return "", err
	}

	qs := url.Values{}
	qs.Add("suite_id", s.id)
	qs.Add("pre_auth_code", preAuthCodeInfo.Code)
	qs.Add("redirect_uri", redirectURI)
	qs.Add("state", state)

	return authURI + "?" + qs.Encode(), nil
}

func (s *Suite) GetPreAuthCode() (*preAuthCodeInfo, error) {
	result := &preAuthCodeInfo{}
	err := s.GetJSON(preAuthCodeURI, nil, result)
	return result, err
}

// SetSessionInfo 设置授权配置
func (s *Suite) SetSessionInfo(PreAuthCode string) error {
	body := map[string]interface{}{
		"pre_auth_code": PreAuthCode,
		"session_info":  map[string]interface{}{"auth_type": 1},
	}
	var result BaseResp
	return s.PostJSON(setSessionInfoURI, nil, body, &result)
}

// GetInstallURI 方法用于获取应用套件的授权地址
func (s *Suite) GetInstallURI(redirectURI, state string, isTest bool) (string, error) {
	preAuthCodeInfo, err := s.GetPreAuthCode()
	if err != nil {
		return "", err
	}
	if isTest {
		s.SetSessionInfo(preAuthCodeInfo.Code)
	}

	qs := url.Values{}
	qs.Add("suite_id", s.id)
	qs.Add("pre_auth_code", preAuthCodeInfo.Code)
	qs.Add("redirect_uri", redirectURI)
	qs.Add("state", state)

	return installURI + "?" + qs.Encode(), nil
}

// GetPermanentCode 方法用于获取企业的永久授权码
func (s *Suite) GetPermanentCode(authCode string) (PermanentCodeInfo, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return PermanentCodeInfo{}, err
	}

	qs := url.Values{}
	qs.Add("suite_access_token", token)
	uri := permanentCodeURI + "?" + qs.Encode()

	buf, _ := json.Marshal(map[string]interface{}{
		"suite_id":  s.id,
		"auth_code": authCode,
	})

	body, err := s.client.PostJSON(uri, buf)
	if err != nil {
		return PermanentCodeInfo{}, err
	}

	result := PermanentCodeInfo{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("GetPermanentCode", string(body))
	}

	return result, err
}

// GetCorpAuthInfo 方法用于获取已授权当前套件的企业号的授权信息
func (s *Suite) GetCorpAuthInfo(corpID, permanentCode string) (CorpAuthInfo, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return CorpAuthInfo{}, err
	}

	qs := url.Values{}
	qs.Add("suite_access_token", token)
	uri := authInfoURI + "?" + qs.Encode()

	buf, _ := json.Marshal(map[string]string{
		"suite_id":       s.id,
		"auth_corpid":    corpID,
		"permanent_code": permanentCode,
	})

	body, err := s.client.PostJSON(uri, buf)
	if err != nil {
		return CorpAuthInfo{}, err
	}

	result := CorpAuthInfo{}
	err = json.Unmarshal(body, &result)

	return result, err
}

// GetCropAgent 方法用于获取已授权当前套件的企业号的某个应用信息
func (s *Suite) GetCropAgent(corpID, permanentCode, agentID string) (CorpAgent, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return CorpAgent{}, err
	}

	qs := url.Values{}
	qs.Add("suite_access_token", token)
	uri := getAgentURI + "?" + qs.Encode()

	buf, _ := json.Marshal(map[string]string{
		"suite_id":       s.id,
		"auth_corpid":    corpID,
		"permanent_code": permanentCode,
		"agentid":        agentID,
	})

	body, err := s.client.PostJSON(uri, buf)
	if err != nil {
		return CorpAgent{}, err
	}

	result := CorpAgent{}
	err = json.Unmarshal(body, &result)

	return result, err
}

// UpdateCorpAgent 方法用于设置已授权当前套件的企业号的某个应用信息
func (s *Suite) UpdateCorpAgent(corpID, permanentCode string, agent AgentEditInfo) error {
	token, err := s.tokener.Token()
	if err != nil {
		return err
	}

	qs := url.Values{}
	qs.Add("suite_access_token", token)
	uri := setAgentURI + "?" + qs.Encode()

	data := struct {
		SuiteID       string        `json:"suite_id"`
		AuthCorpID    string        `json:"auth_corpid"`
		PermanentCode string        `json:"permanent_code"`
		Agent         AgentEditInfo `json:"agent"`
	}{
		s.id,
		corpID,
		permanentCode,
		agent,
	}

	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = s.client.PostJSON(uri, buf)
	return err
}

func (s *Suite) fetchCorpToken(corpID, permanentCode string) (*corpTokenInfo, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}

	qs := url.Values{}
	qs.Add("suite_access_token", token)
	uri := corpTokenURI + "?" + qs.Encode()

	buf, _ := json.Marshal(map[string]string{
		"suite_id":       s.id,
		"auth_corpid":    corpID,
		"permanent_code": permanentCode,
	})

	body, err := s.client.PostJSON(uri, buf)
	if err != nil {
		return nil, err
	}

	result := &corpTokenInfo{}
	err = json.Unmarshal(body, result)

	return result, err
}

func (s *Suite) ActiveAccount(corpID, activeCode, userid string) (*BaseResp, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}

	qs := url.Values{}
	qs.Add("provider_access_token", token)
	uri := activeAccountURI + "?" + qs.Encode()

	buf, _ := json.Marshal(map[string]string{
		"active_code": activeCode,
		"corpid":      corpID,
		"userid":      userid,
	})

	body, err := s.client.PostJSON(uri, buf)
	if err != nil {
		return nil, err
	}

	result := &BaseResp{}
	err = json.Unmarshal(body, result)

	return result, err
}

// GetAdminList 获取应用的管理员列表
func (s *Suite) GetAdminList(corpID, agentId string) ([]*Admin, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}

	qs := url.Values{}
	qs.Add("suite_access_token", token)
	uri := adminListURI + "?" + qs.Encode()

	buf, _ := json.Marshal(map[string]string{
		"auth_corpid": corpID,
		"agentid":     agentId,
	})

	body, err := s.client.PostJSON(uri, buf)
	if err != nil {
		return nil, err
	}

	result := &struct {
		Admins []*Admin `json:"admin"`
	}{}

	err = json.Unmarshal(body, result)

	return result.Admins, err
}

// GetRegisterCode 获取注册码
func (s *Suite) GetRegisterCode(templateId string) (*RegisterCodeInfo, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}

	qs := url.Values{}
	qs.Add("provider_access_token", token)
	uri := registerCode + "?" + qs.Encode()

	buf, _ := json.Marshal(map[string]string{
		"template_id": templateId,
	})

	body, err := s.client.PostJSON(uri, buf)
	if err != nil {
		return nil, err
	}

	result := &RegisterCodeInfo{}
	err = json.Unmarshal(body, result)

	return result, err
}

// GetRegisterURI 方法用于获取应用套件的授权地址
func (s *Suite) GetRegisterURI(templateId string) (string, error) {
	registerCodeInfo, err := s.GetRegisterCode(templateId)
	if err != nil {
		log.Println(err)
		return "", err
	}

	qs := url.Values{}
	qs.Add("register_code", registerCodeInfo.Code)

	return registerURI + "?" + qs.Encode(), nil
}

// ContactSyncSuccess 设置通讯录同步完成
func (s *Suite) ContactSyncSuccess(accessToken string) error {
	qs := url.Values{}
	qs.Add("access_token", accessToken)
	uri := contactSyncSuccessURI + "?" + qs.Encode()

	body, err := s.client.GetJSON(uri)
	if err != nil {
		return err
	}

	result := &struct {
		ErrMsg string `json:"errmsg"`
	}{}
	err = json.Unmarshal(body, result)
	return err
}

// Getuserinfo3rd 获取访问用户身份
func (s *Suite) Getuserinfo3rd(code string) (*UserInfo3RD, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}
	var result UserInfo3RD
	_, err = resty.New().R().SetResult(&result).SetQueryParams(map[string]string{
		"suite_access_token": token,
		"code":               code,
	}).Get(getuserinfo3rdURI)

	if result.Errcode > 0 {
		return nil, errors.New(result.Errmsg)
	}
	return &result, err
}

func (s *Suite) Getuserinfo3rdAuth(code string) (*UserInfo3RDAuth, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}
	var result UserInfo3RDAuth
	resp, err := resty.New().R().SetResult(&result).SetQueryParams(map[string]string{
		"suite_access_token": token,
		"code":               code,
	}).Get(getuserinfo3rdAuthURI)
	fmt.Println(getuserinfo3rdAuthURI, resp.String())

	if result.Errcode > 0 {
		return nil, errors.New(result.Errmsg)
	}
	return &result, err
}

// Getuserdetail3rd 获取访问用户敏感信息
func (s *Suite) Getuserdetail3rd(userTicket string) (*UserInfoDetail3RD, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}
	var result UserInfoDetail3RD
	_, err = resty.New().R().SetResult(&result).SetQueryParam("suite_access_token", token).SetBody(map[string]string{
		"user_ticket": userTicket,
	}).Post(getuserdetail3rdURI)

	if result.Errcode > 0 {
		return nil, errors.New(result.Errmsg)
	}
	return &result, err
}

// GetLoginInfo 获取登录用户信息
func (s *Suite) GetLoginInfo(authCode string) (*LoginInfo, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}
	var result LoginInfo
	_, err = resty.New().R().SetResult(&result).SetQueryParam("access_token", token).SetBody(map[string]string{
		"auth_code": authCode,
	}).Post(getLoginInfoURI)

	if result.Errcode > 0 {
		return nil, errors.New(result.Errmsg)
	}
	return &result, err
}

// ContactIdTranslate 获取转义文件url
func (s *Suite) ContactIdTranslate(corpid string, fileByte []byte) (string, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return "", err
	}
	client := resty.New()
	var media MediaInfo
	_, err = client.R().SetResult(&media).
		SetQueryParam("provider_access_token", token).SetQueryParam("type", "file").
		SetFileReader("media", "员工列表.csv", bytes.NewReader(fileByte)).
		Post(uploadURI)
	if err != nil {
		return "", err
	}
	if media.Errcode > 0 {
		return "", errors.New(media.Errmsg)
	}

	var job JobInfo
	_, err = client.R().SetResult(&job).SetQueryParam("provider_access_token", token).
		SetBody(map[string]interface{}{
			"auth_corpid":   corpid,
			"media_id_list": []string{media.MediaID},
			// "output_file_name": "员工列表",
			// "output_file_format": "pdf",
		}).Post(idTranslateURI)
	if err != nil {
		return "", err
	}
	if job.Errcode > 0 {
		return "", errors.New(job.Errmsg)
	}
	return job.Jobid, nil
}

// GetJobResultURI 获取任务结果
func (s *Suite) GetJobResultURI(jobID string) (*JobResult, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}
	var result JobResult
	_, err = resty.New().R().SetResult(&result).
		SetQueryParam("provider_access_token", token).
		SetQueryParam("jobid", jobID).Get(getJobResultURI)
	if err != nil {
		return nil, err
	}
	if result.Errcode > 0 {
		return nil, errors.New(result.Errmsg)
	}
	return &result, err
}

func (s *Suite) PostJSON(uri string, param url.Values, body, result interface{}) error {
	token, err := s.tokener.Token()
	if err != nil {
		return err
	}
	if param == nil {
		param = url.Values{}
	}
	param.Add("suite_access_token", token)
	uri = uri + "?" + param.Encode()

	buf, _ := json.Marshal(body)
	_body, err := s.client.PostJSON(uri, buf)
	if err != nil {
		return err
	}
	return json.Unmarshal(_body, result)
}

func (s *Suite) GetJSON(uri string, qs url.Values, result interface{}) error {
	token, err := s.tokener.Token()
	if err != nil {
		return err
	}
	if qs == nil {
		qs = url.Values{}
	}
	qs.Add("suite_access_token", token)
	body, err := s.client.GetJSON(uri + "?" + qs.Encode())
	if err != nil {
		return err
	}
	return json.Unmarshal(body, result)
}

// 获取成员的激活详情
func (s *Suite) GetActiveInfoByUser(corpID, userID string) (*ActiveInfo, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := url.Values{}
	qs.Add("provider_access_token", token)
	uri := getActiveInfoByUserURI + "?" + qs.Encode()
	buf, _ := json.Marshal(map[string]string{
		"corpid": corpID,
		"userid": userID,
	})
	body, err := s.client.PostJSON(uri, buf)
	if err != nil {
		return nil, err
	}
	result := ActiveInfo{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// 获取企业的账号列表
func (s *Suite) ListActivedAccount(corpID string) (*ActivedList, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := url.Values{}
	qs.Add("provider_access_token", token)
	uri := listActivedAccountURI + "?" + qs.Encode()
	buf, _ := json.Marshal(map[string]any{
		"corpid": corpID,
		"limit":  1000,
	})
	body, err := s.client.PostJSON(uri, buf)
	if err != nil {
		return nil, err
	}
	var result ActivedList
	err = json.Unmarshal(body, &result)
	return &result, err
}

// 获取订单列表
func (s *Suite) ListOrder(corpId string) (*OrderListRes, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := url.Values{}
	qs.Add("provider_access_token", token)
	uri := listOrderURI + "?" + qs.Encode()
	buf, _ := json.Marshal(map[string]any{
		"corpid": corpId,
	})
	body, err := s.client.PostJSON(uri, buf)
	if err != nil {
		return nil, err
	}
	var result OrderListRes
	err = json.Unmarshal(body, &result)
	return &result, err
}

// 获取订单详情
func (s *Suite) GetOrder(orderID string) (*GetOrderResp, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := url.Values{}
	qs.Add("provider_access_token", token)
	uri := getOrderURI + "?" + qs.Encode()
	buf, _ := json.Marshal(map[string]any{
		"order_id": orderID,
	})
	body, err := s.client.PostJSON(uri, buf)
	if err != nil {
		return nil, err
	}
	var result GetOrderResp
	err = json.Unmarshal(body, &result)
	return &result, err
}

// 获取订单中的账号列表
func (s *Suite) ListOrderAccount(orderID string) (*OrderAccountRes, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := url.Values{}
	qs.Add("provider_access_token", token)
	uri := listOrderAccountURI + "?" + qs.Encode()
	buf, _ := json.Marshal(map[string]any{
		"order_id": orderID,
	})
	body, err := s.client.PostJSON(uri, buf)
	if err != nil {
		return nil, err
	}
	var result OrderAccountRes
	err = json.Unmarshal(body, &result)
	return &result, err
}
