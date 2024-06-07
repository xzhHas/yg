package model

import (
	"context"
	"strconv"

	"gorm.io/gorm"

	"xback/goods_srv/global"
)

type Category struct {
	BaseModel
	Name             string      `gorm:"type:varchar(20);not null" json:"name"`
	ParentCategoryID int32       `json:"parent"` // 数据库存储的外键ID
	ParentCategory   *Category   `json:"-"`      // gorm中外键自己指向自己，需要使用指针
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
	Level            int32       `gorm:"type:int;not null;default:1;comment:'1为1级类目，2为2级...'" json:"level"`
	IsTab            bool        `gorm:"default:false;not null;comment:'能否展示在Tab栏'" json:"is_tab"`
}

type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null;comment:'品牌名称'"`
	Logo string `gorm:"type:varchar(200);default:'';not null;comment:'品牌图标'"`
}

// 商品和品牌的对应关系
// 使用相同的index，gorm就会建成联合的索引
type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"` // 商品外键
	Category   Category

	BrandsID int32 `gorm:"type:int;index:idx_category_brand,unique"` // 品牌外键
	Brands   Brands
}

func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null"`
	Url   string `gorm:"type:varchar(200);not null"`
	Index int32  `gorm:"type:int;default:1;not null"`
}

type Goods struct {
	BaseModel

	CategoryID int32 `gorm:"type:int;not null"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;not null"`
	Brands     Brands

	OnSale   bool `gorm:"default:false;not null;comment:'是否上架'"`
	ShipFree bool `gorm:"default:false;not null;comment:'是否免运费'"`
	IsNew    bool `gorm:"default:false;not null;comment:'是否新品'"`
	IsHot    bool `gorm:"default:false;not null;comment:'是否热卖商品'"`

	Name            string   `gorm:"type:varchar(50);not null"`
	GoodsSn         string   `gorm:"type:varchar(50);not null;comment:'商家的内部编号'"`
	ClickNum        int32    `gorm:"type:int;default:0;not null;comment:'点击数'"`
	SoldNum         int32    `gorm:"type:int;default:0;not null;comment:'销售量'"`
	FavNum          int32    `gorm:"type:int;default:0;not null;comment:'收藏数'"`
	MarketPrice     float32  `gorm:"not null;comment:'商品价格'"`
	ShopPrice       float32  `gorm:"not null;comment:'实际价格'"`
	GoodsBrief      string   `gorm:"type:varchar(100);not null;comment:'商品简介'"`
	Images          GormList `gorm:"type:varchar(1000);not null;comment:'商品图片'"`
	DescImages      GormList `gorm:"type:varchar(1000);not null;comment:'商品图片'"`
	GoodsFrontImage string   `gorm:"type:varchar(200);not null;comment:'商品展示图'"`
}

func (g *Goods) AfterCreate(tx *gorm.DB) (err error) {
	esModel := EsGoods{
		ID:          g.ID,
		CategoryID:  g.CategoryID,
		BrandsID:    g.BrandsID,
		OnSale:      g.OnSale,
		ShipFree:    g.ShipFree,
		IsNew:       g.IsNew,
		IsHot:       g.IsHot,
		Name:        g.Name,
		ClickNum:    g.ClickNum,
		SoldNum:     g.SoldNum,
		FavNum:      g.FavNum,
		MarketPrice: g.MarketPrice,
		GoodsBrief:  g.GoodsBrief,
		ShopPrice:   g.ShopPrice,
	}

	_, err = global.EsClient.Index().Index(esModel.GetIndexName()).BodyJson(esModel).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (g *Goods) AfterUpdate(tx *gorm.DB) (err error) {
	esModel := EsGoods{
		ID:          g.ID,
		CategoryID:  g.CategoryID,
		BrandsID:    g.BrandsID,
		OnSale:      g.OnSale,
		ShipFree:    g.ShipFree,
		IsNew:       g.IsNew,
		IsHot:       g.IsHot,
		Name:        g.Name,
		ClickNum:    g.ClickNum,
		SoldNum:     g.SoldNum,
		FavNum:      g.FavNum,
		MarketPrice: g.MarketPrice,
		GoodsBrief:  g.GoodsBrief,
		ShopPrice:   g.ShopPrice,
	}

	_, err = global.EsClient.Update().Index(esModel.GetIndexName()).
		Doc(esModel).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (g *Goods) AfterDelete(tx *gorm.DB) (err error) {
	_, err = global.EsClient.Delete().Index(EsGoods{}.GetIndexName()).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
