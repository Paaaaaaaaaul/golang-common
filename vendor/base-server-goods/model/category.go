package model

type CategoryStatus int

const (
	//商品类别状态
	STATUS_ENABLE  CategoryStatus = 1 //启用
	STATUS_DISABLE CategoryStatus = 2 //禁用
)

type Category struct {
	OrgId      int            `gorm:"column:orgId" json:"orgId"`                                      // 所属项目id
	MchId      int64          `gorm:"column:mchId" json:"mchId"`                                      // 所属商家id
	CategoryId int            `gorm:"column:categoryId;primary_key;AUTO_INCREMENT" json:"categoryId"` // 类别id
	ParentId   int            `gorm:"column:parentId" json:"parentId"`                                // 父类id
	Name       string         `gorm:"column:name" json:"name"`                                        // 类别名称
	Status     CategoryStatus `gorm:"column:status" json:"status"`                                    // 状态 1:可用 2:禁用
	SortOrder  int            `gorm:"column:sortOrder" json:"sortOrder"`                              // 排列次序
	CreateTime int64          `gorm:"column:createTime" json:"createTime"`                            // 创建时间
	UpdateTime int64          `gorm:"column:updateTime" json:"updateTime"`                            // 更新时间
}

func (*Category) TableName() string {
	return "goods_category"
}
