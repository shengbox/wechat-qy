package suite

import "github.com/shengbox/wechat-qy/base"

type preAuthCodeInfo struct {
	Code      string `json:"pre_auth_code"`
	ExpiresIn int64  `json:"expires_in"`
}

type corpTokenInfo struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// Corporation 用于表示授权方企业信息
type Corporation struct {
	ID              string `json:"corpid"`
	Name            string `json:"corp_name"`
	CorpFullName    string `json:"corp_full_name"`
	Type            string `json:"corp_type"`
	RoundLogoURI    string `json:"corp_round_logo_url"`
	SquareLogoURI   string `json:"corp_square_logo_url"`
	UserMax         int    `json:"corp_user_max"`
	AgentMax        int    `json:"corp_agent_max"`
	QRCode          string `json:"corp_wxqrcode"`
	SubjectType     string `json:"subject_type"`
	VerifiedEndTime string `json:"verified_end_time"`
}

// Agent 用于表示应用基本信息
type Agent struct {
	ID                   int64  `json:"agentid"`
	Name                 string `json:"name,omitempty"`
	RoundLogoURI         string `json:"round_logo_url,omitempty"`
	SquareLogoURI        string `json:"square_logo_url,omitempty"`
	Description          string `json:"description,omitempty"`
	RedirectDomain       string `json:"redirect_domain,omitempty"`
	RedirectLocationFlag int64  `json:"report_location_flag,omitempty"`
	IsReportUser         int64  `json:"isreportuser,omitempty"`
	IsReportEnter        int64  `json:"isreportenter,omitempty"`
}

type authorizedAgent struct {
	Agent
	AppID    int64    `json:"appid"`
	APIGroup []string `json:"api_group"`
}

type authorizedDepartment struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	ParentID int64  `json:"parentid"`
	Writable bool   `json:"writable"`
}

// AuthInfo 表示授权基本信息
type AuthInfo struct {
	Agent      []*authorizedAgent      `json:"agent"`
	Department []*authorizedDepartment `json:"department"`
}

// PermanentCodeInfo 代表获取企业号永久授权码时的响应信息
type PermanentCodeInfo struct {
	AccessToken   string       `json:"access_token"`
	ExpiresIn     int64        `json:"expires_in"`
	PermanentCode string       `json:"permanent_code"`
	AuthCorpInfo  *Corporation `json:"auth_corp_info"`
	AuthInfo      *AuthInfo    `json:"auth_info"`
	Ticket        *base.Ticket `json:"ticket"`
}

type operator struct {
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}

// CorpAuthInfo 代表企业号的授权信息
type CorpAuthInfo struct {
	AuthCorpInfo *Corporation `json:"auth_corp_info"`
	AuthInfo     *AuthInfo    `json:"auth_info"`
	AuthUserInfo *operator    `json:"auth_user_info"`
}

type allowUser struct {
	UserID string `json:"userid"`
	Status string `json:"status"`
}

type allowUsers struct {
	User []*allowUser `json:"user"`
}

type allowPartys struct {
	PartyID []int64 `json:"partyid"`
}

type allowTags struct {
	TagID []int64 `json:"tagid"`
}

// CorpAgent 用于表示授权方企业号某个应用的基本信息
type CorpAgent struct {
	Agent
	AllowUsers  *allowUsers  `json:"allow_userinfos"`
	AllowPartys *allowPartys `json:"allow_partys"`
	AllowTags   *allowTags   `json:"allow_tags"`
	Close       int64        `json:"close"`
}

// AgentEditInfo 代表设置授权方企业号某个应用时的应用信息
type AgentEditInfo struct {
	Agent
	LogoMediaID string `json:"logo_mediaid,omitempty"`
}

// RecvSuiteTicket 用于记录应用套件 ticket 的被动响应结果
type RecvSuiteTicket struct {
	SuiteId     string
	InfoType    string
	TimeStamp   float64
	SuiteTicket string
}

// RecvSuiteAuth 用于记录应用套件授权变更和授权撤销的被动响应结果
type RecvSuiteAuth struct {
	SuiteId    string
	InfoType   string
	TimeStamp  float64
	AuthCorpId string
}

// RecvCreateAuth 用于记录应用套件授权创建被动响应结果
type RecvCreateAuth struct {
	SuiteId   string
	AuthCode  string
	InfoType  string
	TimeStamp float64
}

type RecvChangeEvent struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	AgentID      int64
	Event        string
}

//======= 客户事件 =======

type RecvChangeExternalEvent struct {
	SuiteId    string //第三方应用ID
	AuthCorpId string //授权企业的CorpID
	InfoType   string
	TimeStamp  int64 //时间戳
	ChangeType string
}

type RecvChangeContactEvent struct {
	RecvChangeExternalEvent `xml:",inline"`
	UserID                  string
	OpenUserID              string
}

type RecvChangeExternalTagEvent struct {
	RecvChangeExternalEvent `xml:",inline"`
	Id                      string
}

type RecvChangeExternalContactEvent struct {
	RecvChangeExternalEvent `xml:",inline"`
	UserID                  string //企业服务人员的UserID
	ExternalUserID          string //外部联系人的userid，注意不是企业成员的帐号
	State                   string //添加此用户的「联系我」方式配置的state参数，可用于识别添加此用户的渠道
	WelcomeCode             string //欢迎语code，可用于发送欢迎语
	FailReason              string
}

type RecvChangeExternalChatEvent struct {
	RecvChangeExternalEvent `xml:",inline"`
	ChatId                  string
	UpdateDetail            string
	JoinScene               int
	QuitScene               int
	MemChangeCnt            int
}

// Admin 获取应用的管理员列表
type Admin struct {
	Userid     string `json:"userid"`
	OpenUserid string `json:"open_userid"`
	AuthType   int    `json:"auth_type"`
}

type RegisterCodeInfo struct {
	Code      string `json:"register_code"`
	ExpiresIn int64  `json:"expires_in"`
}

// RecRegisterCorp 注册完成回调事件
type RecRegisterCorp struct {
	ServiceCorpId string
	InfoType      string
	TimeStamp     float64
	RegisterCode  string
	AuthCorpId    string
	ContactSync   *ContactSync
	AuthUserInfo  *AuthUserInfo
	State         string
}

type ContactSync struct {
	AccessToken string
	ExpiresIn   int64
}

type AuthUserInfo struct {
	UserId string
}

// Generated by https://quicktype.io

type UserInfo3RD struct {
	Errcode    int64  `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	CorpID     string `json:"CorpId"`
	UserID     string `json:"UserId"`
	DeviceID   string `json:"DeviceId"`
	UserTicket string `json:"user_ticket"`
	ExpiresIn  int64  `json:"expires_in"`
	OpenUserid string `json:"open_userid"`
}

// Generated by https://quicktype.io

type UserInfoDetail3RD struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Corpid  string `json:"corpid"`
	Userid  string `json:"userid"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Avatar  string `json:"avatar"`
	QrCode  string `json:"qr_code"`
}

// Generated by https://quicktype.io

type LoginInfo struct {
	Errcode  int64    `json:"errcode"`
	Errmsg   string   `json:"errmsg"`
	Usertype int64    `json:"usertype"`
	UserInfo UserInfo `json:"user_info"`
	CorpInfo CorpInfo `json:"corp_info"`
	Agent    []Agent  `json:"agent"`
	AuthInfo AuthInfo `json:"auth_info"`
}

type Department struct {
	ID       int64 `json:"id"`
	Writable bool  `json:"writable"`
}

type CorpInfo struct {
	Corpid string `json:"corpid"`
}

type UserInfo struct {
	Userid     string `json:"userid"`
	OpenUserid string `json:"open_userid"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
}

type MediaInfo struct {
	Errcode   int64  `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

type JobInfo struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Jobid   string `json:"jobid"`
}

// Generated by https://quicktype.io

type JobResult struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Status  int64  `json:"status"`
	Type    string `json:"type"`
	Result  struct {
		ContactIDTranslate struct {
			URL string `json:"url"`
		} `json:"contact_id_translate"`
	} `json:"result"`
}
