package model

type AccountLog struct {
	LogId           int64  `gorm:"column:logId;primary_key;AUTO_INCREMENT" json:"logId"`
	OrgId           int    `gorm:"column:orgId" json:"orgId"`
	UserId          int64  `gorm:"column:userId" json:"userId"`
	Currency        string `gorm:"column:currency" json:"currency"`
	BsType          int    `gorm:"column:bsType" json:"bsType"`
	LogType         int    `gorm:"column:logType" json:"logType"`
	Amount          string `gorm:"column:amount" json:"amount"`
	BeforeAvailable string `gorm:"column:beforeAvailable" json:"beforeAvailable"`
	AfterAvailable  string `gorm:"column:afterAvailable" json:"afterAvailable"`
	BeforeFreeze    string `gorm:"column:beforeFreeze" json:"beforeFreeze"`
	AfterFreeze     string `gorm:"column:afterFreeze" json:"afterFreeze"`
	Detail          string `gorm:"column:detail" json:"detail"`
	Ext             string `gorm:"column:ext" json:"ext"`
	AttachId        int64  `gorm:"column:attachId" json:"attachId"`
	CreateTime      int64  `gorm:"column:createTime" json:"createTime"`
}

func (*AccountLog) TableName() string {
	return "account_log"
}

type LogList struct {
	LogId      int64  `json:"logId"`
	UserId     int64  `json:"userId"`
	Currency   string `json:"currency"`
	BsType     int    `json:"bsType"`
	LogType    int    `json:"logType"`
	Amount     string `json:"amount"`
	Detail     string `json:"detail"`
	Ext        string `json:"ext"`
	CreateTime int64  `json:"createTime"`
	AttachId   int64  `json:"attachId"`
}
