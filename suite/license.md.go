package suite

type ActiveInfoRes struct {
	BaseResp       `json:",inline"`
	ActiveStatus   int64            `json:"active_status"`
	ActiveInfoList []ActiveInfoList `json:"active_info_list"`
}

type ActiveInfoList struct {
	ActiveCode string `json:"active_code"`
	Type       int64  `json:"type"`
	Userid     string `json:"userid"`
	ActiveTime int64  `json:"active_time"`
	ExpireTime int64  `json:"expire_time"`
}

type CodeActiveInfoRes struct {
	BaseResp   `json:",inline"`
	ActiveInfo ActiveInfo `json:"active_info"`
}

type ActiveInfo struct {
	ActiveCode string    `json:"active_code"`
	Type       int64     `json:"type"`
	Status     int64     `json:"status"`
	Userid     string    `json:"userid"`
	CreateTime int64     `json:"create_time"`
	ActiveTime int64     `json:"active_time"`
	ExpireTime int64     `json:"expire_time"`
	MergeInfo  MergeInfo `json:"merge_info"`
	ShareInfo  ShareInfo `json:"share_info"`
}

type MergeInfo struct {
	ToActiveCode   string `json:"to_active_code"`
	FromActiveCode string `json:"from_active_code"`
}

type ShareInfo struct {
	ToCorpid   string `json:"to_corpid"`
	FromCorpid string `json:"from_corpid"`
}

type ActivedList struct {
	BaseResp    `json:",inline"`
	HasMore     int64         `json:"has_more"`
	NextCursor  string        `json:"next_cursor"`
	AccountList []AccountList `json:"account_list"`
}

type AccountList struct {
	Type       int64  `json:"type"`        // 1:基础账号，2:互通账号
	Userid     string `json:"userid"`      // 企业的成员userid。返回加密的userid
	ActiveTime int64  `json:"active_time"` // 激活时间
	ExpireTime int64  `json:"expire_time"` // 过期时间
}

type OrderListRes struct {
	BaseResp   `json:",inline"`
	NextCursor string  `json:"next_cursor"`
	HasMore    int64   `json:"has_more"`
	OrderList  []Order `json:"order_list"`
}
type Order struct {
	OrderID   string `json:"order_id"`
	OrderType int64  `json:"order_type"`
}

type GetOrderResp struct {
	BaseResp `json:",inline"`
	Order    OrderInfo `json:"order"`
}
type OrderInfo struct {
	OrderID         string          `json:"order_id"`
	OrderType       int64           `json:"order_type"`
	OrderStatus     int64           `json:"order_status"`
	Corpid          string          `json:"corpid"`
	Price           int64           `json:"price"`
	AccountCount    AccountCount    `json:"account_count"`
	AccountDuration AccountDuration `json:"account_duration"`
	CreateTime      int64           `json:"create_time"`
	PayTime         int64           `json:"pay_time"`
}

type AccountCount struct {
	BaseCount            int64 `json:"base_count"`
	ExternalContactCount int64 `json:"external_contact_count"`
}

type AccountDuration struct {
	Months        int64 `json:"months"`
	Days          int64 `json:"days"`
	NewExpireTime int64 `json:"new_expire_time"`
}

type OrderAccountRes struct {
	BaseResp    `json:",inline"`
	HasMore     int64          `json:"has_more"`
	NextCursor  string         `json:"next_cursor"`
	AccountList []OrderAccount `json:"account_list"`
}
type OrderAccount struct {
	ActiveCode string `json:"active_code"`
	Type       int64  `json:"type"`
	Userid     string `json:"userid"`
}
