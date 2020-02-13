package model

type TaskStatus int
type TaskType int
type ProductStatus int

const (
	TASK_STATUS_SUCCESS TaskStatus = 1 //任务待处理状态
	TASK_STATUS_PROCESS TaskStatus = 2 //任务完成状态

	TASK_TYPE_GOODS_CHANGE_STATUS TaskType = 1 //上下架任务类型
	TASK_TYPE_OPERATE_STOCK       TaskType = 2 //批量修改库存

	PRODUCT_STATUS_ON_SHELVES  ProductStatus = 1 //上架
	PRODUCT_STATUS_OFF_SHELVES ProductStatus = 2 //下架
)

type Task struct {
	TaskId       int64      `gorm:"column:taskId;primary_key;AUTO_INCREMENT" json:"taskId"`
	TaskType     TaskType   `gorm:"column:taskType" json:"taskType"`
	Status       TaskStatus `gorm:"column:status" json:"status"`
	OrgId        int        `gorm:"column:orgId" json:"orgId"`
	MchId        int64      `gorm:"column:mchId" json:"mchId"`
	Detail       string     `gorm:"column:detail" json:"detail"`
	ParentTaskId int64      `gorm:"column:parentTaskId" json:"parentTaskId"`
	ExecTime     int64      `gorm:"column:execTime" json:"execTime"`
	CreateTime   int64      `gorm:"column:createTime" json:"createTime"`
	UpdateTime   int64      `gorm:"column:updateTime" json:"updateTime"`
}

type TaskGoodsChangeStatus struct {
	ProductId int64         `json:"productId"`
	Status    ProductStatus `json:"status"`
}

type TaskCallBack struct {
	CallBackUrl string            `json:"callBackUrl"`
	Data        map[string]string `json:"data"`
}

func (*Task) TableName() string {
	return "goods_task"
}
