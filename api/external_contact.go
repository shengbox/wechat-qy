package api

import (
	"encoding/json"
	"net/url"
)

const (
	getExternalContactURI  = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get"
	listExternalContactURI = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list"
)

type ExternalContact struct {
	Avatar         string `json:"avatar"`
	ExternalUserid string `json:"external_userid"`
	Gender         int    `json:"gender"`
	Name           string `json:"name"`
	Type           int    `json:"type"`
}

type FollowUser struct {
	AddWay         int      `json:"add_way"`
	Createtime     int      `json:"createtime"`
	Description    string   `json:"description"`
	OperUserid     string   `json:"oper_userid"`
	Remark         string   `json:"remark"`
	RemarkCorpName string   `json:"remark_corp_name"`
	RemarkMobiles  []string `json:"remark_mobiles"`
	Tags           []struct {
		GroupName string `json:"group_name"`
		TagId     string `json:"tag_id"`
		TagName   string `json:"tag_name"`
		Type      int    `json:"type"`
	} `json:"tags"`
	Userid string `json:"userid"`
}

type ExternalContactResp struct {
	ExternalContact ExternalContact `json:"external_contact"`
	FollowUser      []FollowUser    `json:"follow_user"`
}

type ExternalContactListResp struct {
	Errcode        int      `json:"errcode"`
	Errmsg         string   `json:"errmsg"`
	ExternalUserid []string `json:"external_userid"`
}

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
