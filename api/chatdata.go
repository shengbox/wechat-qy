package api

import (
	"encoding/json"
	"errors"
	"net/url"
)

const (
	syncCallProgramURI    = "https://qyapi.weixin.qq.com/cgi-bin/chatdata/sync_call_program"
	asyncProgramTaskURI   = "https://qyapi.weixin.qq.com/cgi-bin/chatdata/async_program_task"
	asyncProgramResultURI = "https://qyapi.weixin.qq.com/cgi-bin/chatdata/async_program_result"
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
