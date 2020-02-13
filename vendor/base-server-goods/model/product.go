package model

type Product struct {
	ProductId  int64  `gorm:"column:productId;primary_key;AUTO_INCREMENT" json:"productId"` // 商品id
	OrgId      int    `gorm:"column:orgId" json:"orgId"`                                    // 所属项目id
	MchId      int64  `gorm:"column:mchId" json:"mchId"`                                    // 所属商家id
	CategoryId int    `gorm:"column:categoryId" json:"categoryId"`                          // 类别id
	Title      string `gorm:"column:title" json:"title"`                                    // 商品名称
	SubTitle   string `gorm:"column:subTitle" json:"subTitle"`                              // 商品子名称
	Detail     string `gorm:"column:detail" json:"detail"`                                  // 商品详情
	MainImg    string `gorm:"column:mainImg" json:"mainImg"`                                // 主图
	SubImg     string `gorm:"column:subImg" json:"subImg"`                                  // 副图集
	Video      string `gorm:"column:video" json:"video"`                                    // 视频
	Status     int    `gorm:"column:status" json:"status"`                                  // 商品状态 1:上架 2:下架
	CreateTime int64  `gorm:"column:createTime" json:"createTime"`                          // 创建时间
	UpdateTime int64  `gorm:"column:updateTime" json:"updateTime"`                          // 更新时间
}

type ProductDetail struct {
	*Product
	SkuList []*Sku
}

func (*Product) TableName() string {
	return "goods_product"
}
