package model

const (
	//账户状态
	STATUS_ENABLE  = 1	//启用
	STATUS_DISABLE = 2	//禁用
)

type Account struct {
	AccountId    int64  `gorm:"column:accountId;primary_key;AUTO_INCREMENT" json:"accountId"`
	OrgId        int    `gorm:"column:orgId" json:"orgId"`
	UserId       int64  `gorm:"column:userId" json:"userId"`
	Currency     string `gorm:"column:currency" json:"currency"`
	AvailAmount  string `gorm:"column:availAmount" json:"availAmount"`
	FreezeAmount string `gorm:"column:freezeAmount" json:"freezeAmount"`
	Status       int    `gorm:"column:status" json:"status"`
	CreateTime   int64  `gorm:"column:createTime" json:"createTime"`
	UpdateTime   int64  `gorm:"column:updateTime" json:"updateTime"`
}

func (*Account) TableName() string {
	return "account"
}
