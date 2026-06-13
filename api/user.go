package api

import (
	"net/url"
	"strconv"
)

const (
	createUserURI      = "https://qyapi.weixin.qq.com/cgi-bin/user/create"
	updateUserURI      = "https://qyapi.weixin.qq.com/cgi-bin/user/update"
	deleteUserURI      = "https://qyapi.weixin.qq.com/cgi-bin/user/delete"
	batchDeleteUserURI = "https://qyapi.weixin.qq.com/cgi-bin/user/batchdelete"
	getUserURI         = "https://qyapi.weixin.qq.com/cgi-bin/user/get"
	listSimpleUserURI  = "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist"
	listUserURI        = "https://qyapi.weixin.qq.com/cgi-bin/user/list"
	inviteUserURI      = "https://qyapi.weixin.qq.com/cgi-bin/invite/send"
	listMemberAuthURI  = "https://qyapi.weixin.qq.com/cgi-bin/user/list_member_auth"
)

// UserAttribute struct 为用户扩展信息
type UserAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// UserAttributes struct 为用户扩展信息列表
type UserAttributes struct {
	Attrs []*UserAttribute `json:"attrs,omitempty"`
}

// User struct 为企业用户信息
type User struct {
	UserID        string         `json:"userid"`
	Name          string         `json:"name,omitempty"`
	Alias         string         `json:"alias,omitempty"`
	DepartmentIds []int64        `json:"department,omitempty"`
	Position      string         `json:"position,omitempty"`
	Mobile        string         `json:"mobile,omitempty"`
	Email         string         `json:"email,omitempty"`
	WeixinID      string         `json:"weixinid,omitempty"`
	Enable        *int           `json:"enable,omitempty"`
	Avatar        string         `json:"avatar,omitempty"`
	Status        *int           `json:"status,omitempty"`
	ExtAttr       UserAttributes `json:"extattr,omitempty"`
}

// CreateUser 方法用于创建用户
func (a *API) CreateUser(user *User) error {
	return a.PostJSON(createUserURI, nil, user, nil)
}

// UpdateUser 方法用于更新用户信息
func (a *API) UpdateUser(user *User) error {
	return a.PostJSON(updateUserURI, nil, user, nil)
}

// DeleteUser 方法用于删除某个用户
func (a *API) DeleteUser(userID string) error {
	qs := make(url.Values)
	qs.Add("userid", userID)
	return a.GetJSON(deleteUserURI, qs, nil)
}

// BatchDeleteUser 方法用于批量删除用户
func (a *API) BatchDeleteUser(userIds []string) error {
	body := map[string][]string{
		"useridlist": userIds,
	}
	return a.PostJSON(batchDeleteUserURI, nil, body, nil)
}

// GetUser 方法用于获取某个用户的信息
func (a *API) GetUser(userID string) (*User, error) {
	qs := make(url.Values)
	qs.Add("userid", userID)
	user := &User{}
	err := a.GetJSON(getUserURI, qs, user)
	return user, err
}

// ListSimpleUser 方法用于获取部门成员列表（成员仅有简单信息）
func (a *API) ListSimpleUser(departmentID int64, fetchChild *int, status *int) ([]*User, error) {
	qs := make(url.Values)
	qs.Add("department_id", strconv.FormatInt(departmentID, 10))
	if fetchChild != nil {
		qs.Add("fetch_child", strconv.Itoa(*fetchChild))
	}
	if status != nil {
		qs.Add("status", strconv.Itoa(*status))
	}

	result := &struct {
		UserList []*User `json:"userlist"`
	}{}
	err := a.GetJSON(listSimpleUserURI, qs, result)
	return result.UserList, err
}

// ListUser 方法用于获取部门成员列表（成员带有详情信息）
func (a *API) ListUser(departmentID int64, fetchChild *int, status *int) ([]*User, error) {
	qs := make(url.Values)
	qs.Add("department_id", strconv.FormatInt(departmentID, 10))
	if fetchChild != nil {
		qs.Add("fetch_child", strconv.Itoa(*fetchChild))
	}
	if status != nil {
		qs.Add("status", strconv.Itoa(*status))
	}

	result := &struct {
		UserList []*User `json:"userlist"`
	}{}
	err := a.GetJSON(listUserURI, qs, result)
	return result.UserList, err
}

// InviteUser 方法用于邀请成员关注
func (a *API) InviteUser(userID, inviteTips string) (inviteType int, err error) {
	body := map[string]string{
		"userid":      userID,
		"invite_tips": inviteTips,
	}
	result := &struct {
		Type int `json:"type"`
	}{}
	err = a.PostJSON(inviteUserURI, nil, body, result)
	if err != nil {
		return 0, err
	}
	return result.Type, nil
}

func (a *API) ListMemberAuth(cursor string, limit int) (result *ListMemberAuthRes, err error) {
	body := map[string]any{
		"cursor": cursor,
		"limit":  limit,
	}
	result = &ListMemberAuthRes{}
	err = a.PostJSON(listMemberAuthURI, nil, body, result)
	return result, err
}

type ListMemberAuthRes struct {
	Errcode        int64  `json:"errcode"`
	Errmsg         string `json:"errmsg"`
	NextCursor     string `json:"next_cursor"`
	MemberAuthList []struct {
		OpenUserid string `json:"open_userid"`
	} `json:"member_auth_list"`
}
