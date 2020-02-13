package model

type Task struct {
	TaskId       int64  `gorm:"column:taskId;primary_key;AUTO_INCREMENT" json:"taskId"`
	TaskType     int    `gorm:"column:taskType" json:"taskType"`
	Status       int    `gorm:"column:status" json:"status"`
	OrgId        int    `gorm:"column:orgId" json:"orgId"`
	Detail       string `gorm:"column:detail" json:"detail"`
	ParentTaskId int64  `gorm:"column:parentTaskId" json:"parentTaskId"`
	CreateTime   int64  `gorm:"column:createTime" json:"createTime"`
	UpdateTime   int64  `gorm:"column:updateTime" json:"updateTime"`
}

type TaskOperateAmount struct {
	OpType        int    `json:"opType"`
	BsType        int    `json:"bsType"`
	AccountId     int64  `json:"accountId"`
	AllowNegative int    `json:"allowNegative"`
	Amount        string `json:"amount"`
	UserId        int64  `json:"userId"`
	Currency      string `json:"currency"`
	Detail        string `json:"detail"`
	Ext           string `json:"ext"`
}

type TaskCallBack struct {
	CallBackUrl string            `json:"callBackUrl"`
	Data        map[string]string `json:"data"`
}

func (*Task) TableName() string {
	return "account_task"
}
