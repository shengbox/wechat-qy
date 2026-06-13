package api

type BaseResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type ExternalContact struct {
	Avatar          string         `json:"avatar"`
	ExternalUserid  string         `json:"external_userid"`
	Gender          int            `json:"gender"`
	Name            string         `json:"name"`
	Type            int            `json:"type"`
	Unionid         string         `json:"unionid"`
	Position        string         `json:"position"`
	CorpName        string         `json:"corp_name"`
	CorpFullName    string         `json:"corp_full_name"`
	ExternalProfile map[string]any `json:"external_profile"`
}

type FollowUser struct {
	AddWay         int      `json:"add_way"` // 该成员添加此客户的来源，具体含义详见来源定义
	Createtime     int      `json:"createtime"`
	Deletetime     int64    `json:"deletetime"`
	Description    string   `json:"description"`
	OperUserid     string   `json:"oper_userid"`
	State          string   `json:"state"` // 企业自定义的state参数，用于区分客户具体是通过哪个「联系我」添加，由企业通过创建「联系我」方式指定
	Remark         string   `json:"remark"`
	RemarkCorpName string   `json:"remark_corp_name"`
	RemarkMobiles  []string `json:"remark_mobiles"`
	Tags           []Tag    `json:"tags"`
	Userid         string   `json:"userid"`
	Username       string   `json:"username"`
}
type Tag struct {
	GroupName string `json:"group_name,omitempty"`
	TagId     string `json:"tag_id,omitempty"`
	TagName   string `json:"tag_name,omitempty"`
	Type      int    `json:"type,omitempty"`
}

type ExternalContactResp struct {
	ExternalContact ExternalContact `json:"external_contact"`
	FollowUser      []FollowUser    `json:"follow_user"`
}

type ExternalContactListResp struct {
	BaseResp       `json:",inline"`
	ExternalUserid []string `json:"external_userid"`
}

// AddContactWayReq 联系我配置
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

type BatchExternalContactReq struct {
	UseridList []string `json:"userid_list"`
	Cursor     string   `json:"cursor"`
	Limit      int64    `json:"limit"`
}

type BatchExternalContactResp struct {
	BaseResp            `json:",inline"`
	ExternalContactList []struct {
		FollowInfo      FollowInfo      `json:"follow_info"`
		ExternalContact ExternalContact `json:"external_contact"`
	} `json:"external_contact_list"`
	NextCursor string `json:"next_cursor"`
}
type FollowInfo struct {
	TagID      []string `json:"tag_id"`
	FollowUser `json:",inline"`
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
	BaseResp `json:",inline"`
	ConfigId string `json:"config_id"`
	QrCode   string `json:"qr_code"`
}

type UserBehaviorDataReq struct {
	Userid    []string `json:"userid,omitempty"`
	Partyid   []int    `json:"partyid,omitempty"`
	StartTime int64    `json:"start_time,omitempty"`
	EndTime   int64    `json:"end_time,omitempty"`
}

type UserBehaviorDataResp struct {
	BaseResp     `json:",inline"`
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
	DayBeginTime int64 `json:"day_begin_time,omitempty"`
	DayEndTime   int64 `json:"day_end_time,omitempty"`
	OwnerFilter  struct {
		UseridList []string `json:"userid_list,omitempty"`
	} `json:"owner_filter,omitempty"`
	OrderBy  int `json:"order_by,omitempty"`
	OrderAsc int `json:"order_asc,omitempty"`
	Offset   int `json:"offset,omitempty"`
	Limit    int `json:"limit,omitempty"`
}

type GroupChatStatisticResp struct {
	BaseResp   `json:",inline"`
	Total      int `json:"total"`
	NextOffset int `json:"next_offset"`
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

// ExternalContactRemark 修改客户备注信息
type ExternalContactRemark struct {
	Userid           string   `json:"userid"`
	ExternalUserid   string   `json:"external_userid"`
	Remark           string   `json:"remark,omitempty"`
	Description      string   `json:"description,omitempty"`        // 此用户对外部联系人的描述，最多150个字符
	RemarkCompany    string   `json:"remark_company,omitempty"`     // 此用户对外部联系人备注的所属公司名称，最多20个字符
	RemarkMobiles    []string `json:"remark_mobiles,omitempty"`     // 此用户对外部联系人备注的手机号
	RemarkPicMediaid string   `json:"remark_pic_mediaid,omitempty"` // 备注图片的mediaid，
}

// GroupChatReq 获取客户群列表
type GroupChatReq struct {
	StatusFilter int `json:"status_filter,omitempty"` // 客户群跟进状态过滤。
	OwnerFilter  struct {
		UseridList []string `json:"userid_list"`
	} `json:"owner_filter,omitempty"` // 群主过滤。
	Cursor string `json:"cursor,omitempty"` // 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用不填
	Limit  int    `json:"limit"`
}

// GroupChatResp 获取客户群列表
type GroupChatResp struct {
	BaseResp      `json:",inline"`
	GroupChatList []struct {
		ChatId string `json:"chat_id"`
		Status int    `json:"status"`
	} `json:"group_chat_list"`
	NextCursor string `json:"next_cursor"`
}

type GroupChatGetReq struct {
	ChatId   string `json:"chat_id"`
	NeedName int    `json:"need_name"`
}

// Generated by https://quicktype.io
type MomentListReq struct {
	StartTime  int64  `json:"start_time,omitempty"`
	EndTime    int64  `json:"end_time,omitempty"`
	Creator    string `json:"creator,omitempty"`
	FilterType int64  `json:"filter_type,omitempty"`
	Cursor     string `json:"cursor,omitempty"`
	Limit      int64  `json:"limit,omitempty"`
}

// Generated by https://quicktype.io

type MomentListResp struct {
	BaseResp   `json:",inline"`
	NextCursor string   `json:"next_cursor"`
	MomentList []Moment `json:"moment_list"`
}

type Moment struct {
	MomentID    string `json:"moment_id" bson:"moment_id"`
	Creator     string `json:"creator" bson:"creator"`
	CreateTime  int64  `json:"create_time" bson:"create_time"`
	CreateType  int64  `json:"create_type" bson:"create_type"`
	VisibleType int64  `json:"visible_type" bson:"visible_type"`
	Text        struct {
		Content string `json:"content" bson:"content"`
	} `json:"text" bson:"text"`
	Image []struct {
		MediaID string `json:"media_id" bson:"media_id"`
	} `json:"image" bson:"image"`
}

type GroupChat struct {
	ChatId     string `json:"chat_id" bson:"chat_id"`
	Name       string `json:"name"`
	Owner      string `json:"owner"`
	CreateTime int    `json:"create_time" bson:"create_time"`
	Notice     string `json:"notice"`
	MemberList []struct {
		Userid    string `json:"userid"`
		Type      int    `json:"type"`
		JoinTime  int    `json:"join_time" bson:"join_time"`
		JoinScene int    `json:"join_scene" bson:"join_scene"`
		Invitor   struct {
			Userid string `json:"userid"`
		} `json:"invitor,omitempty"`
		GroupNickname string `json:"group_nickname"`
		Name          string `json:"name"`
		Unionid       string `json:"unionid,omitempty"`
	} `json:"member_list" bson:"member_list"`
	AdminList []struct {
		Userid string `json:"userid"`
	} `json:"admin_list" bson:"admin_list"`
}

type CorpTagListResp struct {
	BaseResp `json:",inline"`
	TagGroup []TagGroup `json:"tag_group"`
}

type TagGroup struct {
	GroupId    string `json:"group_id"`
	GroupName  string `json:"group_name"`
	CreateTime int    `json:"create_time"`
	Tag        []struct {
		Id         string `json:"id"`
		Name       string `json:"name"`
		CreateTime int    `json:"create_time"`
		Order      int    `json:"order"`
	} `json:"tag"`
	Order int `json:"order"`
}

type MakeTagReq struct {
	Userid         string   `json:"userid"`
	ExternalUserid string   `json:"external_userid"`
	AddTag         []string `json:"add_tag"`
	RemoveTag      []string `json:"remove_tag"`
}

// Generated by https://quicktype.io

type GroupmsgListReq struct {
	ChatType   string `json:"chat_type"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
	Creator    string `json:"creator"`
	FilterType int64  `json:"filter_type"`
	Limit      int64  `json:"limit"`
	Cursor     string `json:"cursor"`
}

// Generated by https://quicktype.io

type GroupMsgListResp struct {
	Errcode      int64          `json:"errcode"`
	Errmsg       string         `json:"errmsg"`
	NextCursor   string         `json:"next_cursor"`
	GroupMsgList []GroupMsgList `json:"group_msg_list"`
}

type GroupMsgList struct {
	Msgid      string `json:"msgid"`
	Creator    string `json:"creator"`
	CreateTime string `json:"create_time"`
	CreateType int64  `json:"create_type"`
	Text       struct {
		Content string `json:"content"`
	} `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Msgtype     string       `json:"msgtype,omitempty"`
	Image       *Image       `json:"image,omitempty"`
	Link        *Link        `json:"link,omitempty"`
	Miniprogram *Miniprogram `json:"miniprogram,omitempty"`
	Video       *File        `json:"video,omitempty"`
	File        *File        `json:"file,omitempty"`
}

type File struct {
	MediaID string `json:"media_id"`
}

type Image struct {
	MediaID string `json:"media_id,omitempty"`
	PicURL  string `json:"pic_url,omitempty"`
}

type Link struct {
	Title   string `json:"title"`
	Picurl  string `json:"picurl"`
	Desc    string `json:"desc"`
	URL     string `json:"url"`
	MediaID string `json:"media_id,omitempty"`
}

type Miniprogram struct {
	Title      string `json:"title"`
	PicMediaID string `json:"pic_media_id"`
	Appid      string `json:"appid"`
	Page       string `json:"page"`
}

type WelcomeMsg struct {
	WelcomeCode string       `json:"welcome_code"`
	Text        Text         `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

type Text struct {
	Content string `json:"content"`
}

type ContactWayRes struct {
	BaseResp   `json:",inline"`
	ContactWay []struct {
		ConfigID string `json:"config_id"`
	} `json:"contact_way"`
	NextCursor string `json:"next_cursor"`
}

type ContactWayDetailRes struct {
	BaseResp   `json:",inline"`
	ContactWay ContactWay `json:"contact_way"`
}

type ContactWay struct {
	ConfigID    string   `json:"config_id"`
	Type        int64    `json:"type"`
	Scene       int64    `json:"scene"`
	QrCode      string   `json:"qr_code"`
	SkipVerify  bool     `json:"skip_verify"`
	Remark      string   `json:"remark"`
	State       string   `json:"state"`
	User        []string `json:"user"`
	Party       []any    `json:"party"`
	Conclusions any      `json:"conclusions"`
}

type NewExternalUseridRes struct {
	BaseResp `json:",inline"`
	Items    []struct {
		ExternalUserid    string `json:"external_userid"`
		NewExternalUserid string `json:"new_external_userid"`
	} `json:"items"`
}

type MsgTemplate struct {
	ChatType       string       `json:"chat_type,omitempty"`
	ExternalUserid []string     `json:"external_userid,omitempty"`
	ChatIDList     []string     `json:"chat_id_list,omitempty"`
	TagFilter      TagFilter    `json:"tag_filter,omitempty"`
	Sender         string       `json:"sender,omitempty"`
	AllowSelect    bool         `json:"allow_select,omitempty"`
	Text           Text         `json:"text,omitempty"`
	Attachments    []Attachment `json:"attachments,omitempty"`
}

type TagFilter struct {
	GroupList []GroupList `json:"group_list,omitempty"`
}

type GroupList struct {
	TagList []string `json:"tag_list,omitempty"`
}

type MsgTemplateRes struct {
	BaseResp `json:",inline"`
	FailList []string `json:"fail_list"`
	Msgid    string   `json:"msgid"`
}

type CreateLinkReq struct {
	LinkName string `json:"link_name"`
	Range    struct {
		UserList       []string `json:"user_list"`
		DepartmentList []int64  `json:"department_list"`
	} `json:"range"`
	SkipVerify bool `json:"skip_verify"`
}

type CreateLinkResp struct {
	BaseResp
	Link struct {
		LinkID     string `json:"link_id"`
		LinkName   string `json:"link_name"`
		URL        string `json:"url"`
		CreateTime int64  `json:"create_time"`
	} `json:"link"`
}
type MomentTask struct {
	Text         Text         `json:"text,omitempty"`
	Attachments  []Attachment `json:"attachments,omitempty"`
	VisibleRange VisibleRange `json:"visible_range,omitempty"`
}

type VisibleRange struct {
	SenderList          *SenderList          `json:"sender_list,omitempty"`
	ExternalContactList *ExternalContactList `json:"external_contact_list,omitempty"`
}

type ExternalContactList struct {
	TagList []string `json:"tag_list,omitempty"`
}

type SenderList struct {
	UserList       []string `json:"user_list,omitempty"`
	DepartmentList []int64  `json:"department_list,omitempty"`
}

type MomentTaskResp struct {
	BaseResp `json:",inline"`
	Jobid    string `json:"jobid"`
}

type AddTagReq struct {
	GroupID   string `json:"group_id"`
	GroupName string `json:"group_name"`
	Order     int64  `json:"order"`
	Tag       []struct {
		Name  string `json:"name"`
		Order int64  `json:"order"`
	} `json:"tag"`
	Agentid int64 `json:"agentid"`
}

type AddTagResp struct {
	BaseResp `json:",inline"`
	TagGroup TagGroup `json:"tag_group"`
}

type GroupmsgSendResultReq struct {
	Msgid  string `json:"msgid"`
	Userid string `json:"userid"`
	Limit  int64  `json:"limit"`
	Cursor string `json:"cursor"`
}

type GroupmsgSendResultResp struct {
	BaseResp   `json:",inline"`
	NextCursor string         `json:"next_cursor"`
	SendList   []SendListItem `json:"send_list"`
}

type GetGroupmsgTaskResp struct {
	BaseResp   `json:",inline"`
	NextCursor string `json:"next_cursor"`
	TaskList   []struct {
		Userid   string `json:"userid"`
		Status   int64  `json:"status"`
		SendTime int64  `json:"send_time"`
	} `json:"task_list"`
}

type SendListItem struct {
	ExternalUserid string `json:"external_userid"`
	ChatID         string `json:"chat_id"`
	Userid         string `json:"userid"`
	Status         int64  `json:"status"`
	SendTime       int64  `json:"send_time"`
}
type GetMomentTaskResultResp struct {
	BaseResp `json:",inline"`
	Status   int64  `json:"status"`
	Type     string `json:"type"`
	Result   struct {
		Errcode  int64  `json:"errcode"`
		Errmsg   string `json:"errmsg"`
		MomentID string `json:"moment_id"`
	} `json:"result"`
}

type RuleListRes struct {
	BaseResp `json:",inline"`
	RuleList []RuleList `json:"rule_list"`
}
type RuleList struct {
	RuleID     string `json:"rule_id"`
	RuleName   string `json:"rule_name"`
	CreateTime int64  `json:"create_time"`
}

type InterceptRule struct {
	RuleName        string          `json:"rule_name"`
	WordList        []string        `json:"word_list"`
	SemanticsList   []int64         `json:"semantics_list"`
	InterceptType   int64           `json:"intercept_type"`
	ApplicableRange ApplicableRange `json:"applicable_range"`
}

type ApplicableRange struct {
	UserList       []string `json:"user_list"`
	DepartmentList []int64  `json:"department_list"`
}
type MomentCommentsRes struct {
	BaseResp    `json:",inline"`
	CommentList []List `json:"comment_list"`
	LikeList    []List `json:"like_list"`
}

type List struct {
	ExternalUserid *string `json:"external_userid,omitempty"`
	CreateTime     int64   `json:"create_time"`
	Userid         *string `json:"userid,omitempty"`
}
