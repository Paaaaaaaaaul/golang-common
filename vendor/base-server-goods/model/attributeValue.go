package model

type AttributeValue struct {
	OrgId          int    `gorm:"column:orgId" json:"orgId"`                                // 所属项目id
	MchId          int64  `gorm:"column:mchId" json:"mchId"`                                // 所属商家id
	ValueId        int64  `gorm:"column:valueId;primary_key;AUTO_INCREMENT" json:"valueId"` // 属性值id
	AttributeId    int64  `gorm:"column:attributeId" json:"attributeId"`                    // 属性id
	AttributeValue string `gorm:"column:attributeValue" json:"attributeValue"`              // 属性值
	SortOrder      int    `gorm:"column:sortOrder" json:"sortOrder"`                        // 排序次序
	CreateTime     int64  `gorm:"column:createTime" json:"createTime"`                      // 创建时间
	UpdateTime     int    `gorm:"column:updateTime" json:"updateTime"`                      // 更新时间
}

func (*AttributeValue) TableName() string {
	return "goods_attribute_value"
}
