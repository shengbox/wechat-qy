package api

import (
	"encoding/json"
	"errors"
	"net/url"
)

const (
	addInterceptRuleURI     = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_intercept_rule"      // 新建敏感词规则
	getInterceptRuleListURI = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_intercept_rule_list" // 获取敏感词规则列表
	getInterceptRuleURI     = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_intercept_rule"      // 获取敏感词规则详情
	delInterceptRuleURI     = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_intercept_rule"      // 删除敏感词规则
	editInterceptRuleURI    = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/edit_intercept_rule"     // 编辑敏感词规则
)

// 新建敏感词规则
func (a *API) AddInterceptRule(req *InterceptRule) (any, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := addInterceptRuleURI + "?" + qs.Encode()
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	var result struct {
		BaseResp `json:",inline"`
		RuleID   string `json:"rule_id"`
	}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// 编辑敏感词规则
func (a *API) EditInterceptRule(ruleID string, rule *InterceptRule) (any, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := editInterceptRuleURI + "?" + qs.Encode()
	req := struct {
		*InterceptRule `json:",inline"`
		RuleID         string `json:"rule_id"`
	}{
		RuleID:        ruleID,
		InterceptRule: rule,
	}
	data, err := json.Marshal(req)
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

// 获取敏感词规则列表
func (a *API) GetInterceptRuleList() ([]RuleList, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := getInterceptRuleListURI + "?" + qs.Encode()
	body, err := a.Client.GetJSON(apiUrl)
	if err != nil {
		return nil, err
	}
	var result RuleListRes
	err = json.Unmarshal(body, &result)
	if result.Errcode != 0 {
		return nil, errors.New(result.Errmsg)
	}
	return result.RuleList, err
}

// 获取敏感词规则详情
func (a *API) GetInterceptRule(ruleID string) (*InterceptRule, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := getInterceptRuleURI + "?" + qs.Encode()
	data, err := json.Marshal(map[string]any{"rule_id": ruleID})
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	var result struct {
		BaseResp `json:",inline"`
		Rule     InterceptRule `json:"rule"`
	}
	err = json.Unmarshal(body, &result)
	if result.Errcode != 0 {
		return nil, errors.New(result.Errmsg)
	}
	return &result.Rule, err
}

// 删除敏感词规则
func (a *API) DelInterceptRule(ruleID string) error {
	token, err := a.Tokener.Token()
	if err != nil {
		return err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := delInterceptRuleURI + "?" + qs.Encode()
	data, err := json.Marshal(map[string]any{"rule_id": ruleID})
	if err != nil {
		return err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return err
	}
	var result BaseResp
	err = json.Unmarshal(body, &result)
	if result.Errcode != 0 {
		return errors.New(result.Errmsg)
	}
	return err
}
