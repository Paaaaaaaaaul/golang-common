package model

type AttributeKey struct {
	OrgId         int    `gorm:"column:orgId" json:"orgId"`                                        // 所属项目id
	MchId         int64  `gorm:"column:mchId" json:"mchId"`                                        // 所属商家id
	AttributeId   int    `gorm:"column:attributeId;primary_key;AUTO_INCREMENT" json:"attributeId"` // 属性id
	CategoryId    int    `gorm:"column:categoryId" json:"categoryId"`                              // 商品类别id
	AttributeName string `gorm:"column:attributeName" json:"attributeName"`                        // 属性名称
	SortOrder     int    `gorm:"column:sortOrder" json:"sortOrder"`                                // 属性排列
	CreateTime    int64  `gorm:"column:createTime" json:"createTime"`                              // 创建时间
	UpdateTime    int64  `gorm:"column:updateTime" json:"updateTime"`                              // 更新时间
}

func (*AttributeKey) TableName() string {
	return "goods_attribute_key"
}
