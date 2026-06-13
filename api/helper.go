package api

import (
	"encoding/json"
	"net/url"
)

// PostJSON 发送 POST 请求并自动填充 access_token，且支持可选的参数和响应结构解析
func (a *API) PostJSON(uri string, query url.Values, body, result interface{}) error {
	token, err := a.Tokener.Token()
	if err != nil {
		return err
	}

	if query == nil {
		query = make(url.Values)
	}
	query.Set("access_token", token)
	reqURL := uri + "?" + query.Encode()

	var data []byte
	if body != nil {
		data, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}

	resp, err := a.Client.PostJSON(reqURL, data)
	if err != nil {
		return err
	}

	if result != nil {
		return json.Unmarshal(resp, result)
	}
	return nil
}

// GetJSON 发送 GET 请求并自动填充 access_token，并解析响应结果
func (a *API) GetJSON(uri string, query url.Values, result interface{}) error {
	token, err := a.Tokener.Token()
	if err != nil {
		return err
	}

	if query == nil {
		query = make(url.Values)
	}
	query.Set("access_token", token)
	reqURL := uri + "?" + query.Encode()

	resp, err := a.Client.GetJSON(reqURL)
	if err != nil {
		return err
	}

	if result != nil {
		return json.Unmarshal(resp, result)
	}
	return nil
}
