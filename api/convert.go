package api

import (
	"encoding/json"
	"log"
	"net/url"
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
	log.Println(string(body))
	result := &NewExternalUseridRes{}
	err = json.Unmarshal(body, result)
	return result, err
}
