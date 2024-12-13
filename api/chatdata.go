package api

import (
	"encoding/json"
	"errors"
	"net/url"
)

const (
	syncCallProgramURI    = "https://qyapi.weixin.qq.com/cgi-bin/chatdata/sync_call_program" // 应用同步调用专区程序
	asyncProgramTaskURI   = "https://qyapi.weixin.qq.com/cgi-bin/chatdata/async_program_task"
	asyncProgramResultURI = "https://qyapi.weixin.qq.com/cgi-bin/chatdata/async_program_result"
	setPublicKeyURI       = "https://qyapi.weixin.qq.com/cgi-bin/chatdata/set_public_key"       // 设置公钥
	openDebugModeURI      = "https://qyapi.weixin.qq.com/cgi-bin/chatdata/open_debug_mode"      // 应用开启调试模式
	closeDebugModeURI     = "https://qyapi.weixin.qq.com/cgi-bin/chatdata/close_debug_mode"     // 关闭专区调试模式
	setReceiveCallbackURI = "https://qyapi.weixin.qq.com/cgi-bin/chatdata/set_receive_callback" // 设置专区接收回调事件
)

type SyncCallProgramReq struct {
	ProgramID   string `json:"program_id,omitempty"`
	AbilityID   string `json:"ability_id,omitempty"`
	NotifyID    string `json:"notify_id,omitempty"`
	RequestData string `json:"request_data,omitempty"`
}

type SyncCallProgramResp struct {
	BaseResp     `json:",inline"`
	ResponseData string `json:"response_data,omitempty"`
}

// 应用同步调用专区程序
func (a *API) SyncCallProgram(req *SyncCallProgramReq) (string, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return "", err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	url := syncCallProgramURI + "?" + qs.Encode()

	data, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	body, err := a.Client.PostJSON(url, data)
	if err != nil {
		return "", err
	}
	result := &SyncCallProgramResp{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return "", err
	}
	if result.Errcode != 0 {
		return "", errors.New(result.Errmsg)
	}
	return result.ResponseData, nil
}

// 创建专区程序调用任务
func (a *API) AsyncProgramTask(req *SyncCallProgramReq) (string, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return "", err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	url := asyncProgramTaskURI + "?" + qs.Encode()

	data, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	body, err := a.Client.PostJSON(url, data)
	if err != nil {
		return "", err
	}
	result := &struct {
		BaseResp `json:",inline"`
		Jobid    string `json:"jobid"`
	}{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return "", err
	}
	if result.Errcode != 0 {
		return "", errors.New(result.Errmsg)
	}
	return result.Jobid, nil
}

// 获取专区程序任务结果
func (a *API) AsyncProgramResult(jobid string) (string, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return "", err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	url := asyncProgramResultURI + "?" + qs.Encode()

	data, err := json.Marshal(map[string]string{"jobid": jobid})
	if err != nil {
		return "", err
	}
	body, err := a.Client.PostJSON(url, data)
	if err != nil {
		return "", err
	}
	var result struct {
		BaseResp        `json:",inline"`
		ResponseErrcode int    `json:"response_errcode"`
		ResponseData    string `json:"response_data"`
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}
	if result.Errcode != 0 {
		return "", errors.New(result.Errmsg)
	}
	return result.ResponseData, nil
}

// 设置公钥
func (a *API) SetPublicKey(publicKey string, keyVer int) (*BaseResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	url := setPublicKeyURI + "?" + qs.Encode()
	data, err := json.Marshal(map[string]any{
		"public_key":     publicKey,
		"public_key_ver": keyVer,
	})
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(url, data)
	if err != nil {
		return nil, err
	}
	var result BaseResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.Errcode != 0 {
		return nil, errors.New(result.Errmsg)
	}
	return &result, nil
}

// 应用开启调试模式
func (a *API) OpenDebugMode(programId, debugToken string) (*BaseResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	url := openDebugModeURI + "?" + qs.Encode()
	data, err := json.Marshal(map[string]any{
		"program_id":  programId,
		"debug_token": debugToken,
	})
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(url, data)
	if err != nil {
		return nil, err
	}
	var result BaseResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.Errcode != 0 {
		return nil, errors.New(result.Errmsg)
	}
	return &result, nil
}

// 关闭专区调试模式
func (a *API) CloseDebugMode(programId string) (*BaseResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	url := closeDebugModeURI + "?" + qs.Encode()
	data, err := json.Marshal(map[string]any{"program_id": programId})
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(url, data)
	if err != nil {
		return nil, err
	}
	var result BaseResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.Errcode != 0 {
		return nil, errors.New(result.Errmsg)
	}
	return &result, nil
}

// 设置专区接收回调事件
func (a *API) SetReceiveCallback(programId string) (*BaseResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	url := setReceiveCallbackURI + "?" + qs.Encode()
	data, err := json.Marshal(map[string]any{"program_id": programId})
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(url, data)
	if err != nil {
		return nil, err
	}
	var result BaseResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.Errcode != 0 {
		return nil, errors.New(result.Errmsg)
	}
	return &result, nil
}
