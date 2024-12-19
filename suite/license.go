package suite

import (
	"encoding/json"
	"net/url"
)

const (
	activeAccountURI       = "https://qyapi.weixin.qq.com/cgi-bin/license/active_account"          // 激活账号
	getActiveInfoByUserURI = "https://qyapi.weixin.qq.com/cgi-bin/license/get_active_info_by_user" // 获取成员的激活详情
	listActivedAccountURI  = "https://qyapi.weixin.qq.com/cgi-bin/license/list_actived_account"    // 获取企业的账号列表
	listOrderURI           = "https://qyapi.weixin.qq.com/cgi-bin/license/list_order"              // 获取订单列表
	getOrderURI            = "https://qyapi.weixin.qq.com/cgi-bin/license/get_order"               // 获取订单详情
	listOrderAccountURI    = "https://qyapi.weixin.qq.com/cgi-bin/license/list_order_account"      // 获取订单中的账号列表
	getActiveInfoByCodeURI = "https://qyapi.weixin.qq.com/cgi-bin/license/get_active_info_by_code" // 获取激活码详情
)

// 获取订单列表
func (s *Suite) ListOrder(corpId string) (*OrderListRes, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := url.Values{}
	qs.Add("provider_access_token", token)
	uri := listOrderURI + "?" + qs.Encode()
	buf, _ := json.Marshal(map[string]any{
		"corpid": corpId,
	})
	body, err := s.client.PostJSON(uri, buf)
	if err != nil {
		return nil, err
	}
	var result OrderListRes
	err = json.Unmarshal(body, &result)
	return &result, err
}

// 获取订单详情
func (s *Suite) GetOrder(orderID string) (*GetOrderResp, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := url.Values{}
	qs.Add("provider_access_token", token)
	uri := getOrderURI + "?" + qs.Encode()
	buf, _ := json.Marshal(map[string]any{
		"order_id": orderID,
	})
	body, err := s.client.PostJSON(uri, buf)
	if err != nil {
		return nil, err
	}
	var result GetOrderResp
	err = json.Unmarshal(body, &result)
	return &result, err
}

// 获取订单中的账号列表
func (s *Suite) ListOrderAccount(orderID string) (*OrderAccountRes, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := url.Values{}
	qs.Add("provider_access_token", token)
	uri := listOrderAccountURI + "?" + qs.Encode()
	buf, _ := json.Marshal(map[string]any{
		"order_id": orderID,
	})
	body, err := s.client.PostJSON(uri, buf)
	if err != nil {
		return nil, err
	}
	var result OrderAccountRes
	err = json.Unmarshal(body, &result)
	return &result, err
}

// 获取企业的账号列表
func (s *Suite) ListActivedAccount(corpID string) (*ActivedList, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := url.Values{}
	qs.Add("provider_access_token", token)
	uri := listActivedAccountURI + "?" + qs.Encode()
	buf, _ := json.Marshal(map[string]any{
		"corpid": corpID,
		"limit":  1000,
	})
	body, err := s.client.PostJSON(uri, buf)
	if err != nil {
		return nil, err
	}
	var result ActivedList
	err = json.Unmarshal(body, &result)
	return &result, err
}

// 获取成员的激活详情
func (s *Suite) GetActiveInfoByUser(corpID, userID string) (*ActiveInfoRes, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := url.Values{}
	qs.Add("provider_access_token", token)
	uri := getActiveInfoByUserURI + "?" + qs.Encode()
	buf, _ := json.Marshal(map[string]string{
		"corpid": corpID,
		"userid": userID,
	})
	body, err := s.client.PostJSON(uri, buf)
	if err != nil {
		return nil, err
	}
	result := ActiveInfoRes{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// 获取激活码详情
func (s *Suite) GetActiveInfoByCode(corpID, activeCode string) (*CodeActiveInfoRes, error) {
	token, err := s.tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := url.Values{}
	qs.Add("provider_access_token", token)
	uri := getActiveInfoByCodeURI + "?" + qs.Encode()
	buf, _ := json.Marshal(map[string]string{
		"corpid":      corpID,
		"active_code": activeCode,
	})
	body, err := s.client.PostJSON(uri, buf)
	if err != nil {
		return nil, err
	}
	result := CodeActiveInfoRes{}
	err = json.Unmarshal(body, &result)
	return &result, err
}
