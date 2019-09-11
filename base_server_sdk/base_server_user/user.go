package base_server_user

import (
	"errors"
	"github.com/becent/commom/base_server_sdk"
	json "github.com/json-iterator/go"
	"strconv"
)

type User struct {
	UserId     int64  `json:"userId"`
	OrgId      int    `json:"orgId"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Account    string `json:"account"`
	LoginPwd   string `json:"loginPwd"`
	TransPwd   string `json:"transPwd"`
	NickName   string `json:"nickName"`
	Avatar     string `json:"avatar"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	IdCard     string `json:"idCard"`
	Sex        int    `json:"sex"`
	BirthDay   string `json:"birthDay"`
	Status     int    `json:"status"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	Ext        string `json:"ext"`
}

var (
	ErrInvalidParams = errors.New("params invalid or missed")
	ErrServiceBusy   = errors.New("service busy")
)

// Register register a user to base_server_user
func Register(user *User) (*User, error) {
	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(user.OrgId)
	params["phone"] = user.Phone
	params["email"] = user.Email
	params["account"] = user.Account
	params["loginPwd"] = user.LoginPwd
	params["transPwd"] = user.TransPwd
	params["nickName"] = user.NickName
	params["avatar"] = user.Avatar
	params["sex"] = strconv.Itoa(user.Sex)
	params["birthDay"] = user.BirthDay
	params["ext"] = user.Ext

	if params["orgId"] == "" || (params["phone"] == "" && params["account"] == "" && params["email"] == "") || (params["account"] != "" && params["loginPwd"] == "") {
		return nil, ErrInvalidParams
	}

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.UserServerHost, "user", "register", params)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, user); err != nil {
		return nil, ErrServiceBusy
	}

	return user, nil
}
