package api

type ExternalContact struct {
	Avatar         string `json:"avatar"`
	ExternalUserid string `json:"external_userid"`
	Gender         int    `json:"gender"`
	Name           string `json:"name"`
	Type           int    `json:"type"`
}

type FollowUser struct {
	AddWay         int      `json:"add_way"` // 该成员添加此客户的来源，具体含义详见来源定义
	Createtime     int      `json:"createtime"`
	Description    string   `json:"description"`
	OperUserid     string   `json:"oper_userid"`
	State          string   `json:"state"` // 企业自定义的state参数，用于区分客户具体是通过哪个「联系我」添加，由企业通过创建「联系我」方式指定
	Remark         string   `json:"remark"`
	RemarkCorpName string   `json:"remark_corp_name"`
	RemarkMobiles  []string `json:"remark_mobiles"`
	Tags           []struct {
		GroupName string `json:"group_name"`
		TagId     string `json:"tag_id"`
		TagName   string `json:"tag_name"`
		Type      int    `json:"type"`
	} `json:"tags"`
	Userid string `json:"userid"`
}

type ExternalContactResp struct {
	ExternalContact ExternalContact `json:"external_contact"`
	FollowUser      []FollowUser    `json:"follow_user"`
}

type ExternalContactListResp struct {
	Errcode        int      `json:"errcode"`
	Errmsg         string   `json:"errmsg"`
	ExternalUserid []string `json:"external_userid"`
}

//AddContactWayReq 联系我配置
type AddContactWayReq struct {
	Type          int         `json:"type"`            // 联系方式类型,1-单人, 2-多人
	Scene         int         `json:"scene"`           // 场景，1-在小程序中联系，2-通过二维码联系
	Style         int         `json:"style"`           // 在小程序中联系时使用的控件样式，详见附表
	Remark        string      `json:"remark"`          // 联系方式的备注信息，用于助记，不超过30个字符
	SkipVerify    bool        `json:"skip_verify"`     // 外部客户添加时是否无需验证，默认为true
	State         string      `json:"state"`           // 企业自定义的state参数，用于区分不同的添加渠道，在调用“获取外部联系人详情”时会返回该参数值，不超过30个字符
	User          []string    `json:"user"`            // 使用该联系方式的用户userID列表，在type为1时为必填，且只能有一个
	Party         []int       `json:"party"`           // 使用该联系方式的部门id列表，只在type为2时有效
	IsTemp        bool        `json:"is_temp"`         // 是否临时会话模式，true表示使用临时会话模式，默认为false
	ExpiresIn     int         `json:"expires_in"`      // 临时会话二维码有效期，以秒为单位。该参数仅在is_temp为true时有效，默认7天
	ChatExpiresIn int         `json:"chat_expires_in"` // 临时会话有效期，以秒为单位。
	Unionid       string      `json:"unionid"`         // 可进行临时会话的客户unionid，该参数仅在is_temp为true时有效，如不指定则不进行限制
	Conclusions   Conclusions `json:"conclusions"`
}

type Conclusions struct {
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
	Link struct {
		Title  string `json:"title"`
		Picurl string `json:"picurl"`
		Desc   string `json:"desc"`
		Url    string `json:"url"`
	} `json:"link"`
	Miniprogram struct {
		Title      string `json:"title"`
		PicMediaId string `json:"pic_media_id"`
		Appid      string `json:"appid"`
		Page       string `json:"page"`
	} `json:"miniprogram"`
}

type AddContactWayResp struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	ConfigId string `json:"config_id"`
	QrCode   string `json:"qr_code"`
}

type UserBehaviorDataReq struct {
	Userid    []string `json:"userid"`
	Partyid   []int    `json:"partyid"`
	StartTime int64    `json:"start_time"`
	EndTime   int64    `json:"end_time"`
}

type UserBehaviorDataResp struct {
	Errcode      int            `json:"errcode"`
	Errmsg       string         `json:"errmsg"`
	BehaviorData []BehaviorData `json:"behavior_data"`
}

type BehaviorData struct {
	StatTime            int     `json:"stat_time"`
	ChatCnt             int     `json:"chat_cnt"`
	MessageCnt          int     `json:"message_cnt"`
	ReplyPercentage     float64 `json:"reply_percentage"`
	AvgReplyTime        int     `json:"avg_reply_time"`
	NegativeFeedbackCnt int     `json:"negative_feedback_cnt"`
	NewApplyCnt         int     `json:"new_apply_cnt"`
	NewContactCnt       int     `json:"new_contact_cnt"`
}

type GroupChatStatisticReq struct {
	DayBeginTime int64 `json:"day_begin_time"`
	DayEndTime   int64 `json:"day_end_time"`
	OwnerFilter  struct {
		UseridList []string `json:"userid_list"`
	} `json:"owner_filter"`
	OrderBy  int `json:"order_by"`
	OrderAsc int `json:"order_asc"`
	Offset   int `json:"offset"`
	Limit    int `json:"limit"`
}

type GroupChatStatisticResp struct {
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	Total      int    `json:"total"`
	NextOffset int    `json:"next_offset"`
	Items      []struct {
		Owner string `json:"owner"`
		Data  struct {
			NewChatCnt   int `json:"new_chat_cnt"`
			ChatTotal    int `json:"chat_total"`
			ChatHasMsg   int `json:"chat_has_msg"`
			NewMemberCnt int `json:"new_member_cnt"`
			MemberTotal  int `json:"member_total"`
			MemberHasMsg int `json:"member_has_msg"`
			MsgTotal     int `json:"msg_total"`
		} `json:"data"`
	} `json:"items"`
}
