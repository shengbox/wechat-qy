package api

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

const (
	jsSDKTicketURI      = "https://qyapi.weixin.qq.com/cgi-bin/get_jsapi_ticket"
	jsSDKAgentTicketURI = "https://qyapi.weixin.qq.com/cgi-bin/ticket/get"
)

func (a *API) GetJSSDKSignature(uri, timestamp, noncestr string) (string, error) {
	return a.getJSSDKSignature(jsSDKTicketURI, uri, timestamp, noncestr, "")
}

func (a *API) GetJSSDKAgentSignature(uri, timestamp, noncestr string) (string, error) {
	return a.getJSSDKSignature(jsSDKAgentTicketURI, uri, timestamp, noncestr, "agent_config")
}

// GetJSSDKSignature 方法用于获取 JSSDK 的签名
func (a *API) getJSSDKSignature(api, uri, timestamp, noncestr, typeStr string) (string, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return "", err
	}

	qs := make(url.Values)
	qs.Add("access_token", token)
	if typeStr != "" {
		qs.Add("type", typeStr)
	}

	ticketURL := api + "?" + qs.Encode()

	body, err := a.Client.GetJSON(ticketURL)
	if err != nil {
		return "", err
	}

	ticketInfo := struct {
		Ticket    string `json:"ticket"`
		ExpiresIn int64  `json:"expires_in"`
	}{}

	if err = json.Unmarshal(body, &ticketInfo); err != nil {
		return "", err
	}

	signParams := map[string]string{
		"jsapi_ticket": ticketInfo.Ticket,
		"noncestr":     noncestr,
		"timestamp":    timestamp,
		"url":          uri,
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
