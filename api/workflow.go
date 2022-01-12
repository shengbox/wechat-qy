package api

import (
	"errors"
	"log"

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
	result := &struct {
		BaseResp   `json:",inline"`
		TemplateId string `json:"template_id"`
	}{}
	_, err = resty.New().R().SetResult(result).SetQueryParam("access_token", token).SetBody(map[string]string{
		"open_template_id": openTemplateId,
	}).Post(copytemplateURI)
	if err != nil {
		return "", err
	}
	if result.Errcode > 0 {
		return "", errors.New(result.Errmsg)
	}
	return result.TemplateId, err
}

// 获取模板详情
func (a *API) GetTemplateDetail(templateId string) (*TemplateDetailObj, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	var result TemplateDetailObj
	resp, err := resty.New().R().
		SetQueryParam("access_token", token).
		SetBody(map[string]string{"template_id": templateId}).
		SetResult(&result).Post(gettemplatedetailURI)
	log.Println(resp.String(), err)
	return &result, err
}

type TemplateDetailObj struct {
	BaseResp      `json:",inline"`
	TemplateNames []struct {
		Text string `json:"text"`
	} `json:"template_names"`
	TemplateContent TemplateContent `json:"template_content"`
}

type TemplateContent struct {
	Controls []struct {
		Property Property `json:"property"`
	} `json:"controls"`
}

type Property struct {
	Control string `json:"control"`
	ID      string `json:"id"`
	Title   []struct {
		Text string `json:"text"`
	} `json:"title"`
	Placeholder []struct {
		Text string `json:"text"`
	} `json:"placeholder"`
	Require int64 `json:"require"`
	UnPrint int64 `json:"un_print"`
}

// 提交申请
func (a *API) Applyevent(body map[string]interface{}) (*map[string]interface{}, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	resp, err := resty.New().R().SetQueryParam("access_token", token).SetBody(body).SetResult(&result).Post(applyeventURI)
	log.Println(resp.String(), err)
	return &result, err
}
