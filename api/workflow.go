package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"
)

const (
	copytemplateURI = "https://qyapi.weixin.qq.com/cgi-bin/oa/approval/copytemplate"
)

func (a *API) Copytemplate(openTemplateId string) (string, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return "", err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := copytemplateURI + "?" + qs.Encode()
	data, err := json.Marshal(map[string]string{"open_template_id": openTemplateId})
	if err != nil {
		return "", err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	log.Println(string(body))
	if err != nil {
		return "", err
	}
	result := &struct {
		Errcode    int    `json:"errcode"`
		Errmsg     string `json:"errmsg"`
		TemplateId string `json:"template_id"`
	}{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return "", err
	}
	if result.Errcode > 0 {
		return "", errors.New(result.Errmsg)
	}
	return result.TemplateId, err
}
