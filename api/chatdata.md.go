package api

type GetCorpAuthInfoResp struct {
	BaseResp        `json:",inline"`
	AuthEditionList []AuthEditionList `json:"auth_edition_list"`
}

type AuthEditionList struct {
	Edition         int64     `json:"edition"`
	AuthScope       AuthScope `json:"auth_scope"`
	Status          int64     `json:"status"`
	BeginTime       int64     `json:"begin_time"`
	EndTime         int64     `json:"end_time"`
	MsgDurationDays int64     `json:"msg_duration_days"`
}

type AuthScope struct {
	UseridList       []string `json:"userid_list"`
	DepartmentIDList []int64  `json:"department_id_list"`
	TagIDList        []int64  `json:"tag_id_list"`
}

type AuthUserListResp struct {
	BaseResp     `json:",inline"`
	HasMore      int64          `json:"has_more"`
	NextCursor   string         `json:"next_cursor"`
	AuthUserList []AuthUserList `json:"auth_user_list"`
}

type AuthUserList struct {
	Userid      string  `json:"userid"`
	EditionList []int64 `json:"edition_list"`
}
