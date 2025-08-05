package api

import (
	"encoding/json"
	"errors"
	"net/url"
)

const (
	livingCreateURI       = "https://qyapi.weixin.qq.com/cgi-bin/living/create"                // 创建预约直播
	getWatchStatURI       = "https://qyapi.weixin.qq.com/cgi-bin/living/get_watch_stat"        // 获取直播观看明细
	getUserAllLivingidURI = "https://qyapi.weixin.qq.com/cgi-bin/living/get_user_all_livingid" // 获取成员直播ID列表
	getLivingCodeURI      = "https://qyapi.weixin.qq.com/cgi-bin/living/get_living_code"       // 在微信中观看直播或直播回放
	getLivingInfoURI      = "https://qyapi.weixin.qq.com/cgi-bin/living/get_living_info"       // 获取直播详情
)

type LiveCreateReq struct {
	AnchorUserid         string `json:"anchor_userid"`
	Theme                string `json:"theme"`
	LivingStart          int64  `json:"living_start"`
	LivingDuration       int64  `json:"living_duration"`
	Description          string `json:"description"`
	Type                 int64  `json:"type"`
	Agentid              int64  `json:"agentid"`
	RemindTime           int64  `json:"remind_time"`
	ActivityCoverMediaid string `json:"activity_cover_mediaid"`
	ActivityShareMediaid string `json:"activity_share_mediaid"`
	ActivityDetail       struct {
		Description string   `json:"description"`
		ImageList   []string `json:"image_list"`
	} `json:"activity_detail"`
}

type LivingInfo struct {
	AnchorUserid          string `json:"anchor_userid"`
	CommentNum            int64  `json:"comment_num"`
	Description           string `json:"description"`
	LivingDuration        int64  `json:"living_duration"`
	LivingStart           int64  `json:"living_start"`
	MainDepartment        int64  `json:"main_department"`
	MicNum                int64  `json:"mic_num"`
	OnlineCount           int64  `json:"online_count"`
	OpenReplay            int64  `json:"open_replay"`
	ReplayStatus          int64  `json:"replay_status"`
	ReserveLivingDuration int64  `json:"reserve_living_duration"`
	ReserveStart          int64  `json:"reserve_start"`
	Status                int64  `json:"status"`
	SubscribeCount        int64  `json:"subscribe_count"`
	Theme                 string `json:"theme"`
	Type                  int64  `json:"type"`
	ViewerNum             int64  `json:"viewer_num"`
}

type StatInfo struct {
	ExternalUsers []LivingExternalUser `json:"external_users"`
	Users         []any                `json:"users"`
}

type LivingExternalUser struct {
	ExternalUserid string `json:"external_userid"`
	IsComment      int64  `json:"is_comment"`
	IsMic          int64  `json:"is_mic"`
	Name           string `json:"name"`
	Type           int64  `json:"type"`
	WatchTime      int64  `json:"watch_time"`
}

// 创建预约直播
func (a *API) LivingCreate(req *LiveCreateReq) (*string, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := livingCreateURI + "?" + qs.Encode()
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
		Livingid string `json:"livingid"`
	}
	err = json.Unmarshal(body, &result)
	return &result.Livingid, err
}

// 获取直播观看明细
func (a *API) GetWatchStat(livingid, nextKey string) (*StatInfo, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := getWatchStatURI + "?" + qs.Encode()
	req := map[string]any{"livingid": livingid}
	if nextKey != "" {
		req["next_key"] = nextKey
	}
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
		Ending   int      `json:"ending"`
		NextKey  string   `json:"next_key"`
		StatInfo StatInfo `json:"stat_info"`
	}
	err = json.Unmarshal(body, &result)
	return &result.StatInfo, err
}

// 获取成员直播ID列表
func (a *API) GetUserAllLivingid(userid, cursor string) (any, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := getUserAllLivingidURI + "?" + qs.Encode()
	req := map[string]any{"userid": userid, "limit": 20}
	if cursor != "" {
		req["cursor"] = cursor
	}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	var result struct {
		BaseResp     `json:",inline"`
		NextCursor   string   `json:"next_cursor"`
		LivingidList []string `json:"livingid_list"`
	}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// 在微信中观看直播或直播回放
func (a *API) GetLivingCode(livingid, openid string) (*string, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	apiUrl := getLivingCodeURI + "?" + qs.Encode()
	req := map[string]any{"livingid": livingid, "openid": openid}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := a.Client.PostJSON(apiUrl, data)
	if err != nil {
		return nil, err
	}
	var result struct {
		BaseResp   `json:",inline"`
		LivingCode string `json:"living_code"`
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.Errcode > 0 {
		return nil, errors.New(result.Errmsg)
	}
	return &result.LivingCode, err
}

// 获取直播详情
func (a *API) GetLivingInfo(livingid string) (*LivingInfo, error) {
	token, err := a.Tokener.Token()
	if err != nil {
		return nil, err
	}
	qs := make(url.Values)
	qs.Add("access_token", token)
	qs.Add("livingid", livingid)
	apiUrl := getLivingInfoURI + "?" + qs.Encode()
	body, err := a.Client.GetJSON(apiUrl)
	if err != nil {
		return nil, err
	}
	var result struct {
		BaseResp   `json:",inline"`
		LivingInfo LivingInfo `json:"living_info"`
	}
	err = json.Unmarshal(body, &result)
	return &result.LivingInfo, err
}
