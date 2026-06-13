package api

import (
	"encoding/json"
	"net/url"

	"github.com/shengbox/wechat-qy/base"
)

const (
	getNewExternalUseridURI = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_new_external_userid"
)

func (a *API) GetNewExternalUserid(externalUseridList []string) (*NewExternalUseridRes, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	qs.Add("debug", "1")
	apiUrl := getNewExternalUseridURI + "?" + qs.Encode()
	data, err := json.Marshal(map[string]any{
		"external_userid_list": externalUseridList,
	})
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	base.GetLogger().Println(string(body))
	result := &NewExternalUseridRes{}
	err = json.Unmarshal(body, result)
	return result, err
}

const fromServiceExternalUserIDURI = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/from_service_external_userid"

// FromServiceExternalUserIDRes Represents the response for converting service external userid to selfbuilt external userid
type FromServiceExternalUserIDRes struct {
	BaseResp `json:",inline"`
	ExternalUserID string `json:"external_userid"`
}

// FromServiceExternalUserID converts external_userid from third-party/DK app to self-built app
func (a *API) FromServiceExternalUserID(externalUserid string, sourceAgentID int) (*FromServiceExternalUserIDRes, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := fromServiceExternalUserIDURI + "?" + qs.Encode()
	data, err := json.Marshal(map[string]any{
		"external_userid": externalUserid,
		"source_agentid":  sourceAgentID,
	})
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	base.GetLogger().Println(string(body))
	result := &FromServiceExternalUserIDRes{}
	err = json.Unmarshal(body, result)
	return result, err
}


