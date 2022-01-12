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

type ApplyObj struct {
	CreatorUserid       string        `json:"creator_userid"`
	TemplateID          string        `json:"template_id"`
	UseTemplateApprover int64         `json:"use_template_approver"`
	Approver            []Approver    `json:"approver"`
	Notifyer            []string      `json:"notifyer"`
	NotifyType          int64         `json:"notify_type"`
	ApplyData           ApplyData     `json:"apply_data"`
	SummaryList         []SummaryList `json:"summary_list"`
}

type ApplyData struct {
	Contents []Content `json:"contents"`
}

type Content struct {
	Control string        `json:"control"`
	ID      string        `json:"id"`
	Title   []SummaryInfo `json:"title"`
	Value   Value         `json:"value"`
}

type SummaryInfo struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
}

type Value struct {
	Text string `json:"text"`
}

type Approver struct {
	Attr   int64    `json:"attr"`
	Userid []string `json:"userid"`
}

type SummaryList struct {
	SummaryInfo []SummaryInfo `json:"summary_info"`
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

// 提交申请
func (a *API) Applyevent(body ApplyObj) (string, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return "", err
	}
	result := &struct {
		BaseResp `json:",inline"`
		SpNO     string `json:"sp_no"`
	}{}
	resp, err := resty.New().R().SetQueryParam("access_token", token).SetBody(body).SetResult(&result).Post(applyeventURI)
	log.Println(resp.String(), err)
	if err != nil {
		return "", err
	}
	if result.Errcode > 0 {
		return "", errors.New(result.Errmsg)
	}
	return result.SpNO, err
}

func (a *API) GetApprovalInfo(body ApprovalReq) ([]string, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	result := &struct {
		BaseResp `json:",inline"`
		SpNOList []string `json:"sp_no_list"`
	}{}
	_, _ = resty.New().R().SetResult(&result).SetQueryParam("access_token", token).SetBody(body).Post(getapprovalinfoURI)
	return result.SpNOList, err
}

type ApprovalReq struct {
	Starttime int64    `json:"starttime,omitempty"`
	Endtime   int64    `json:"endtime,omitempty"`
	Cursor    int64    `json:"cursor,omitempty"`
	Size      int64    `json:"size,omitempty"`
	Filters   []Filter `json:"filters,omitempty"`
}

type Filter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (a *API) GetApprovalDetail(spNO string) (*map[string]interface{}, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	_, _ = resty.New().R().SetResult(&result).
		SetQueryParam("access_token", token).
		SetBody(map[string]string{"sp_no": spNO}).
		Post(getapprovaldetailURI)
	return &result, err
}
