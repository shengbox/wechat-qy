package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

const (
	getExternalContactURI    = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get"
	batchExternalContactURI  = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/batch/get_by_user"
	listExternalContactURI   = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list"
	addContactWayURI         = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_contact_way"
	getUserBehaviorDataURI   = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_user_behavior_data"
	groupChatStatisticURI    = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/statistic"
	externalContactRemarkURI = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/remark"
	externalGroupChatListURI = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/list"
	externalGroupChatGetURI  = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/get"
	getCorpTagListURI        = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_corp_tag_list"
	addCorpTagURI            = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_corp_tag"
	markTagURI               = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/mark_tag"
	getMomentListURI         = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_list"
	getGroupmsgListURI       = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_list_v2"
	sendWelcomeMsgURI        = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/send_welcome_msg"
	addMsgTemplateURI        = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_msg_template"
	remindGroupmsgSendURI    = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/remind_groupmsg_send"     //提醒成员群发
	getGroupmsgSendResultURI = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_send_result" // 群发结果
	getGroupmsgTask          = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_task"        // 获取群发成员发送任务列表

	listContactWayURI = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list_contact_way"
	getContactWayURI  = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_contact_way"
	delContactWayURI  = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_contact_way"

	createLinkURI          = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_acquisition/create_link" // 创建获客链接
	addMomentTaskURI       = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_moment_task"                  // 创建发表任务
	getMomentTaskResultURI = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_task_result"           // 获取任务创建结果

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

// BatchExternalContact 批量获取客户详情
func (a *API) BatchExternalContact(req *BatchExternalContactReq) (*BatchExternalContactResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := batchExternalContactURI + "?" + qs.Encode()
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	result := &BatchExternalContactResp{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	if result.Errcode != 0 {
		return nil, errors.New(result.Errmsg)
	}
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
	if err != nil {
		return nil, err
	}
	if result.Errcode != 0 {
		return nil, errors.New(result.Errmsg)
	}
	return result.TagGroup, err
}

func (a *API) AddCorpTag(req *AddTagReq) (*TagGroup, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := addCorpTagURI + "?" + qs.Encode()
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	result := &AddTagResp{}
	err = json.Unmarshal(body, result)
	return &result.TagGroup, err
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

// 获取企业已配置的「联系我」列表
func (a *API) ListContactWay(limit int) (*ContactWayRes, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := listContactWayURI + "?" + qs.Encode()
	data, err := json.Marshal(map[string]any{"limit": limit})
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	result := &ContactWayRes{}
	err = json.Unmarshal(body, result)
	return result, err
}

// 获取企业已配置的「联系我」方式
func (a *API) GetContactWay(configID string) (*ContactWay, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := getContactWayURI + "?" + qs.Encode()
	data, err := json.Marshal(map[string]any{
		"config_id": configID,
	})
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}

	result := &ContactWayDetailRes{}
	err = json.Unmarshal(body, result)
	return &result.ContactWay, err
}

func (a *API) DelContactWay(configID string) (*BaseResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := delContactWayURI + "?" + qs.Encode()
	data, err := json.Marshal(map[string]any{
		"config_id": configID,
	})
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

func (a *API) AddMsgTemplate(template *MsgTemplate) (*MsgTemplateRes, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := addMsgTemplateURI + "?" + qs.Encode()
	data, err := json.Marshal(template)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	var result MsgTemplateRes
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (a *API) RemindGroupmsgSend(msgid string) (*BaseResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := remindGroupmsgSendURI + "?" + qs.Encode()
	data, err := json.Marshal(map[string]string{"msgid": msgid})
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	var result BaseResp
	err = json.Unmarshal(body, &result)
	return &result, err
}

// 创建获客链接
func (a *API) CreateLink(req *CreateLinkReq) (*CreateLinkResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := createLinkURI + "?" + qs.Encode()
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	var result CreateLinkResp
	err = json.Unmarshal(body, &result)
	return &result, err
}

// 创建发表任务
func (a *API) AddMomentTask(req *MomentTask) (*MomentTaskResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := addMomentTaskURI + "?" + qs.Encode()
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	var result MomentTaskResp
	err = json.Unmarshal(body, &result)
	return &result, err
}

// 获取任务创建结果
func (a *API) GetMomentTaskResult(jobid string) (*GetMomentTaskResultResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	qs.Add("jobid", jobid)
	apiUrl := getMomentTaskResultURI + "?" + qs.Encode()
	body, err := a.Client.GetJSON(apiUrl)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	var result GetMomentTaskResultResp
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (a *API) GetGroupmsgTask(msgid, cursor string) (*GetGroupmsgTaskResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := getGroupmsgTask + "?" + qs.Encode()
	data, err := json.Marshal(map[string]any{"msgid": msgid, "cursor": cursor})
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	var result GetGroupmsgTaskResp
	err = json.Unmarshal(body, &result)
	return &result, err
}

// 获取企业群发成员执行结果
func (a *API) GetGroupmsgSendResult(req *GroupmsgSendResultReq) (*GroupmsgSendResultResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := getGroupmsgSendResultURI + "?" + qs.Encode()
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	var result GroupmsgSendResultResp
	err = json.Unmarshal(body, &result)
	return &result, err
}
