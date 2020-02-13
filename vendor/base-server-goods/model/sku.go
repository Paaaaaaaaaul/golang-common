package model

type StockOpType int

const (
	ADD_STOCK        StockOpType = 1 //加库存
	SUB_STOCK        StockOpType = 2 //减库存
	FREEZE           StockOpType = 3 //冻结库存
	SUB_FREEZE_STOCK StockOpType = 4 //减冻结库存
	UN_FREEZE        StockOpType = 5 //解冻库存

	STOCK_PRECISION = 18 //库存精度
)

type Sku struct {
	SkuId       int64  `gorm:"column:skuId;primary_key;AUTO_INCREMENT" json:"skuId"` // skuId
	OrgId       int    `gorm:"column:orgId" json:"orgId"`                            // 所属项目id
	MchId       int64  `gorm:"column:mchId" json:"mchId"`                            // 所属商家id
	ProductId   int64  `gorm:"column:productId" json:"productId"`                    // 商品id
	Specs       string `gorm:"column:specs" json:"specs"`                            // 规格列表
	SortOrder   int    `gorm:"column:sortOrder" json:"sortOrder"`                    // 规格序列
	Price       string `gorm:"column:price" json:"price"`                            // 价格
	Stock       string `gorm:"column:stock" json:"stock"`                            // 库存
	FreezeStock string `gorm:"column:freezeStock" json:"freezeStock"`                // 冻结库存
	CreateTime  int64  `gorm:"column:createTime" json:"createTime"`                  // 创建时间
	UpdateTime  int64  `gorm:"column:updateTime" json:"updateTime"`                  // 更新时间
}

//批量操作库存
type BatchOperateStock struct {
	MchId     int64
	SkuId     int64
	ProductId int64
	Qty       string
	OpType    StockOpType
}

func (*Sku) TableName() string {
	return "goods_sku"
}
