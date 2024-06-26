package model

type Category struct {
	BaseModel
	Name             string      `gorm:"type:varchar(50);not null;comment:分类名称" json:"name"`
	ParentCategoryID int32       `json:"parent"`
	ParentCategory   *Category   `json:"-"`
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID" json:"sub_category"`
	Level            int32       `gorm:"type:int;not null;default:1;comment:分类级别" json:"level"`
	IsTab            bool        `gorm:"type:tinyint;not null;default:0;comment:是否在tab栏显示" json:"is_tab"`
}

type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null;comment:品牌名称" json:"name"`
	Logo string `gorm:"type:varchar(200);not null;default:'';comment:品牌logo" json:"logo"`
}

// GoodsCategoryBrand 可以用 gorm 的 many2many 多对多关联来自动生成
type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category   Category

	BrandsID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands   Brands
}

func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

// Banner Image 的长度为 200，实际上可能会有问题，可以根据实际情况调整
type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null;comment:图片地址" json:"image"`
	Url   string `gorm:"type:varchar(200);not null;comment:跳转链接" json:"url"`
	Index int32  `gorm:"type:int;not null;default:1;comment:排序" json:"index"`
}

type Goods struct {
	BaseModel

	CategoryID int32 `gorm:"type:int;not null"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;not null"`
	Brands     Brands

	OnSale   bool `gorm:"type:tinyint;not null;default:0;comment:是否上架"`
	ShipFree bool `gorm:"type:tinyint;not null;default:0;comment:是否包邮"`
	IsNew    bool `gorm:"type:tinyint;not null;default:0;comment:是否新品"`
	IsHot    bool `gorm:"type:tinyint;not null;default:0;comment:是否热销"`

	Name            string   `gorm:"type:varchar(50);not null;comment:商品名称"`
	GoodsSn         string   `gorm:"type:varchar(50);not null;comment:商品编号"`
	ClickNum        int32    `gorm:"type:int;not null;default:0;comment:点击数"`
	SoldNum         int32    `gorm:"type:int;not null;default:0;comment:销售数量"`
	FavNum          int32    `gorm:"type:int;not null;default:0;comment:收藏数量"`
	MarketPrice     float32  `gorm:"type:decimal(10,2);not null;default:0.01;comment:市场价"`
	ShopPrice       float32  `gorm:"type:decimal(10,2);not null;default:0.01;comment:商城价"`
	GoodsBrief      string   `gorm:"type:varchar(200);not null;default:'';comment:商品简介"`
	Images          GormList `gorm:"type:json;not null;default:'[]';comment:商品图片" json:"images"`
	DescImages      GormList `gorm:"type:json;not null;default:'[]';comment:商品描述图片" json:"desc_images"`
	GoodsFrontImage string   `gorm:"type:varchar(200);not null;default:'';comment:商品主图"`
}
