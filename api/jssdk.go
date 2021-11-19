package api

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/shengbox/wechat-qy/base"
)

const (
	jsSDKTicketURI      = "https://qyapi.weixin.qq.com/cgi-bin/get_jsapi_ticket"
	jsSDKAgentTicketURI = "https://qyapi.weixin.qq.com/cgi-bin/ticket/get"
)

func (a *API) GetSDKAgentTicket() (*base.Ticket, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	qs.Add("type", "agent_config")
	ticketURL := jsSDKAgentTicketURI + "?" + qs.Encode()

	body, err := a.Client.GetJSON(ticketURL)
	if err != nil {
		return nil, err
	}
	ticketInfo := base.Ticket{}
	if err = json.Unmarshal(body, &ticketInfo); err != nil {
		return nil, err
	}
	return &ticketInfo, nil
}

// func (a *API) GetJSSDKSignature(uri, timestamp, noncestr string) (string, error) {
// 	return a.getJSSDKSignature(jsSDKTicketURI, uri, timestamp, noncestr, "")
// }

// func (a *API) GetJSSDKAgentSignature(uri, timestamp, noncestr string) (string, error) {
// 	return a.getJSSDKSignature(jsSDKAgentTicketURI, uri, timestamp, noncestr, "agent_config")
// }

// getJSSDKSignature 方法用于获取 JSSDK 的签名
// func (a *API) getJSSDKSignature(api, uri, timestamp, noncestr, typeStr string) (string, error) {
// 	ticketInfo, err := a.GetSDKAgentTicket()
// 	if err != nil {
// 		return "", err
// 	}
// 	return JSSDKSignature(ticketInfo.Ticket, noncestr, timestamp, uri)
// }

func JSSDKSignature(ticket string, noncestr string, timestamp string, url string) (string, error) {
	signParams := map[string]string{
		"jsapi_ticket": ticket,
		"noncestr":     noncestr,
		"timestamp":    timestamp,
		"url":          url,
	}
	var keys []string
	for key := range signParams {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var signStrs []string
	for _, key := range keys {
		signStrs = append(signStrs, fmt.Sprintf("%s=%s", key, signParams[key]))
	}
	hashsum := sha1.Sum([]byte(strings.Join(signStrs, "&")))
	return hex.EncodeToString(hashsum[:]), nil
}
