package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"
)

const (
	syncMsgURI = "https://qyapi.weixin.qq.com/cgi-bin/kf/sync_msg"
	sendMsgURI = "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg"
)

type SyncMsgReq struct {
	Cursor      string `json:"cursor"`
	Token       string `json:"token"`
	Limit       int64  `json:"limit"`
	VoiceFormat int64  `json:"voice_format"`
	OpenKfid    string `json:"open_kfid"`
}

type SyncMsgResp struct {
	BaseResp   `json:",inline"`
	NextCursor string    `json:"next_cursor"`
	HasMore    int64     `json:"has_more"`
	MsgList    []MsgList `json:"msg_list"`
}

type MsgList struct {
	Msgid          string  `json:"msgid"`
	OpenKfid       string  `json:"open_kfid"`
	ExternalUserid string  `json:"external_userid"`
	SendTime       int64   `json:"send_time"`
	Origin         int64   `json:"origin"`
	ServicerUserid string  `json:"servicer_userid"`
	Msgtype        string  `json:"msgtype"`
	Event          KfEvent `json:"event"`
	Text           struct {
		Content string `json:"content"`
		MenuID  string `json:"menu_id"`
	} `json:"text"`
	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
	Video struct {
		MediaId string `json:"media_id"`
	} `json:"video"`
	File struct {
		MediaId string `json:"media_id"`
	} `json:"file"`
}

type KfEvent struct {
	EventType      string `json:"event_type"`
	Scene          string `json:"scene"`
	OpenKfid       string `json:"open_kfid"`
	ExternalUserid string `json:"external_userid"`
	WelcomeCode    string `json:"welcome_code"`
}

func (a *API) SyncMsg(req SyncMsgReq) (*SyncMsgResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := syncMsgURI + "?" + qs.Encode()
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	result := &SyncMsgResp{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	if result.Errcode != 0 {
		return nil, errors.New(result.Errmsg)
	}
	return result, err
}

type SendReq struct {
	Touser   string `json:"touser"`
	OpenKfid string `json:"open_kfid"`
	Msgid    string `json:"msgid,omitempty"`
	Msgtype  string `json:"msgtype"`
	Text     Text   `json:"text,omitempty"`
}

type SendResp struct {
	BaseResp `json:",inline"`
	Msgid    string `json:"msgid"`
}

func (a *API) SendMsg(req SendReq) (*SendResp, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := sendMsgURI + "?" + qs.Encode()
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	log.Println(string(body))
	result := &SendResp{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	if result.Errcode != 0 {
		return nil, errors.New(result.Errmsg)
	}
	return result, err
}
