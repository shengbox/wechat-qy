package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

const (
	getBillListURI = "https://qyapi.weixin.qq.com/cgi-bin/externalpay/get_bill_list"
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
	fmt.Println(string(data))
	body, err := a.Client.PostJSON(url, data)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
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
