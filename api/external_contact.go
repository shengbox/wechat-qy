package api

import (
	"encoding/json"
	"errors"
	"net/url"
)

const (
	getExternalContactURI    = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get"
	listExternalContactURI   = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list"
	addContactWayURI         = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_contact_way"
	getUserBehaviorDataURI   = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_user_behavior_data"
	groupChatStatisticURI    = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/statistic"
	externalContactRemarkURI = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/remark"
	externalGroupChatListURI = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/list"
	externalGroupChatGetURI  = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/get"
	getCorpTagListURI        = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_corp_tag_list"
	markTagURI               = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/mark_tag"
	getMomentListURI         = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_list"
	getGroupmsgListURI       = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_list_v2"
	sendWelcomeMsgURI        = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/send_welcome_msg"
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

// AddContactWay 配置客户联系「联系我」方式
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

// ExternalContactRemark 修改客户备注信息
func (a *API) ExternalContactRemark(req *ExternalContactRemark) error {
	token, err := a.Tokener.Token()
	if err != nil {
		return err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := externalContactRemarkURI + "?" + qs.Encode()
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	_, err = a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return err
	}
	return err
}

// ExternalGroupChatList 获取客户群列表
func (a *API) ExternalGroupChatList(req *GroupChatReq) (*GroupChatResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := externalGroupChatListURI + "?" + qs.Encode()
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	result := &GroupChatResp{}
	err = json.Unmarshal(body, result)
	return result, err
}

// GroupChatGetUGet 获取客户群详情
func (a *API) GroupChatGetUGet(req *GroupChatGetReq) (*GroupChat, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := externalGroupChatGetURI + "?" + qs.Encode()
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	result := &struct {
		BaseResp  `json:",inline"`
		GroupChat GroupChat `json:"group_chat"`
	}{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	if result.Errcode != 0 {
		return nil, errors.New(result.Errmsg)
	}
	return &result.GroupChat, err
}

// GetMomentList 获取客户朋友圈发表记录
func (a *API) GetMomentList(req *MomentListReq) ([]Moment, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := getMomentListURI + "?" + qs.Encode()
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	result := &MomentListResp{}
	err = json.Unmarshal(body, result)
	return result.MomentList, err
}

// GetCorpTagList 获取企业标签库
func (a *API) GetCorpTagList(req interface{}) ([]TagGroup, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := getCorpTagListURI + "?" + qs.Encode()
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	result := &CorpTagListResp{}
	err = json.Unmarshal(body, result)
	return result.TagGroup, err
}

// MarkTag 编辑客户企业标签
func (a *API) MarkTag(req *MakeTagReq) (*BaseResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := markTagURI + "?" + qs.Encode()
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	result := &BaseResp{}
	err = json.Unmarshal(body, result)
	return result, err
}

// GetGroupmsgList 获取群发记录列表
func (a *API) GetGroupmsgList(req *GroupmsgListReq) (*GroupMsgListResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := getGroupmsgListURI + "?" + qs.Encode()
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	result := &GroupMsgListResp{}
	err = json.Unmarshal(body, result)
	return result, err
}

func (a *API) SendWelcomeMsg(req *WelcomeMsg) (*BaseResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := sendWelcomeMsgURI + "?" + qs.Encode()
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	result := &BaseResp{}
	err = json.Unmarshal(body, result)
	return result, err
}
