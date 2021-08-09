package api

import (
	"encoding/json"
	"net/url"
)

const (
	getExternalContactURI  = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get"
	listExternalContactURI = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list"
	addContactWayURI       = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_contact_way"
	getUserBehaviorDataURI = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_user_behavior_data"
	groupChatStatisticURI  = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/statistic"
)

// GetExternalContact 获取客户详情
func (a *API) GetExternalContact(externalUserId string) (*ExternalContactResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}

	qs := make(url.Values)
	qs.Add("access_token", token)
	qs.Add("external_userid", externalUserId)

	apiUrl := getExternalContactURI + "?" + qs.Encode()

	body, err := a.Client.GetJSON(apiUrl)
	if err != nil {
		return nil, err
	}

	result := &ExternalContactResp{}
	err = json.Unmarshal(body, result)

	return result, err
}

// ListExternalContact 获取客户列表
func (a *API) ListExternalContact(userid string) ([]string, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}

	qs := make(url.Values)
	qs.Add("access_token", token)
	qs.Add("userid", userid)

	apiUrl := listExternalContactURI + "?" + qs.Encode()

	body, err := a.Client.GetJSON(apiUrl)
	if err != nil {
		return nil, err
	}

	result := &ExternalContactListResp{}
	err = json.Unmarshal(body, result)

	return result.ExternalUserid, err
}

//AddContactWay 配置客户联系「联系我」方式
func (a *API) AddContactWay(way *AddContactWayReq) (*AddContactWayResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := addContactWayURI + "?" + qs.Encode()
	data, err := json.Marshal(way)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	result := &AddContactWayResp{}
	err = json.Unmarshal(body, result)
	return result, err
}

// GetUserBehaviorData 联系客户统计
func (a *API) GetUserBehaviorData(req *UserBehaviorDataReq) ([]BehaviorData, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := getUserBehaviorDataURI + "?" + qs.Encode()
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	result := &UserBehaviorDataResp{}
	err = json.Unmarshal(body, result)
	return result.BehaviorData, err
}

// GetGroupChatStatistic 群聊数据统计
func (a *API) GetGroupChatStatistic(req *GroupChatStatisticReq) (*GroupChatStatisticResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := groupChatStatisticURI + "?" + qs.Encode()
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	result := &GroupChatStatisticResp{}
	err = json.Unmarshal(body, result)
	return result, err
}
