package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

const (
	getAgentURI = "https://qyapi.weixin.qq.com/cgi-bin/agent/get"
)

type RefundList struct {
	OutRefundNo   string `json:"out_refund_no"`
	RefundUserid  string `json:"refund_userid"`
	RefundComment string `json:"refund_comment"`
	RefundReqtime int64  `json:"refund_reqtime"`
	RefundStatus  int64  `json:"refund_status"`
	RefundFee     int64  `json:"refund_fee"`
}

type AgentInfoResp struct {
	BaseResp       `json:",inline"`
	Agentid        int64  `json:"agentid"`
	Name           string `json:"name"`
	SquareLogoURL  string `json:"square_logo_url"`
	Description    string `json:"description"`
	AllowUserinfos struct {
		User []User `json:"user"`
	} `json:"allow_userinfos"`
	AllowPartys struct {
		Partyid []int64 `json:"partyid"`
	} `json:"allow_partys"`
	Close                   int64  `json:"close"`
	RedirectDomain          string `json:"redirect_domain"`
	ReportLocationFlag      int64  `json:"report_location_flag"`
	Isreportenter           int64  `json:"isreportenter"`
	HomeURL                 string `json:"home_url"`
	CustomizedPublishStatus int64  `json:"customized_publish_status"`
}

func (a *API) GetAgent(agentid string) (*AgentInfoResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	qs.Add("agentid", agentid)
	url := getAgentURI + "?" + qs.Encode()
	body, err := a.Client.GetJSON(url)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	result := &AgentInfoResp{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	if result.Errcode != 0 {
		return nil, errors.New(result.Errmsg)
	}
	return result, nil
}
