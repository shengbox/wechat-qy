package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

const (
	getBillListURI = "https://qyapi.weixin.qq.com/cgi-bin/externalpay/get_bill_list"
	getAgentURI    = "https://qyapi.weixin.qq.com/cgi-bin/agent/get"
)

type GetBillListReq struct {
	BeginTime   int64  `json:"begin_time"`
	EndTime     int64  `json:"end_time"`
	PayeeUserid string `json:"payee_userid,omitempty"`
	Cursor      string `json:"cursor,omitempty"`
	Limit       int64  `json:"limit,omitempty"`
}

type GetBillListResp struct {
	BaseResp   `json:",inline"`
	NextCursor string     `json:"next_cursor"`
	BillList   []BillList `json:"bill_list"`
}

type BillList struct {
	TransactionID   string          `json:"transaction_id"`
	BillType        int64           `json:"bill_type"`
	TradeState      int64           `json:"trade_state"`
	PayTime         int64           `json:"pay_time"`
	OutTradeNo      string          `json:"out_trade_no"`
	OutRefundNo     string          `json:"out_refund_no"`
	ExternalUserid  string          `json:"external_userid"`
	TotalFee        int64           `json:"total_fee"`
	PayeeUserid     string          `json:"payee_userid"`
	PaymentType     int64           `json:"payment_type"`
	MchID           string          `json:"mch_id"`
	Remark          string          `json:"remark"`
	CommodityList   []CommodityList `json:"commodity_list"`
	TotalRefundFee  int64           `json:"total_refund_fee"`
	RefundList      []RefundList    `json:"refund_list"`
	ContactInfo     ContactInfo     `json:"contact_info"`
	MiniprogramInfo MiniprogramInfo `json:"miniprogram_info"`
}

type CommodityList struct {
	Description string `json:"description"`
	Amount      int64  `json:"amount"`
}

type ContactInfo struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type MiniprogramInfo struct {
	Appid string `json:"appid"`
	Name  string `json:"name"`
}

type RefundList struct {
	OutRefundNo   string `json:"out_refund_no"`
	RefundUserid  string `json:"refund_userid"`
	RefundComment string `json:"refund_comment"`
	RefundReqtime int64  `json:"refund_reqtime"`
	RefundStatus  int64  `json:"refund_status"`
	RefundFee     int64  `json:"refund_fee"`
}

// 获取对外收款记录
func (a *API) GetBillListURI(req *GetBillListReq) (*GetBillListResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	url := getBillListURI + "?" + qs.Encode()

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(url, data)
	if err != nil {
		return nil, err
	}
	result := &GetBillListResp{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	if result.Errcode != 0 {
		return nil, errors.New(result.Errmsg)
	}
	return result, nil
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
