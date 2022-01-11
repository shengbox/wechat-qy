package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"

	"github.com/go-resty/resty/v2"
)

const (
	copytemplateURI      = "https://qyapi.weixin.qq.com/cgi-bin/oa/approval/copytemplate"
	gettemplatedetailURI = "https://qyapi.weixin.qq.com/cgi-bin/oa/gettemplatedetail" // 获取审批模板详情
	applyeventURI        = "https://qyapi.weixin.qq.com/cgi-bin/oa/applyevent"        // 提交审批申请
	getapprovalinfoURI   = "https://qyapi.weixin.qq.com/cgi-bin/oa/getapprovalinfo"   // 批量获取审批单号
	getapprovaldetailURI = "https://qyapi.weixin.qq.com/cgi-bin/oa/getapprovaldetail" // 获取审批申请详情
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

func (a *API) GetTemplateDetail(templateId string) (map[string]interface{}, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	resp, err := resty.New().R().
		SetQueryParam("access_token", token).
		SetBody(map[string]string{"template_id": templateId}).
		SetResult(&result).Post(gettemplatedetailURI)
	log.Println(resp.String(), err)
	return result, err
}
