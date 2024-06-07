package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"xback/goods_srv/global"
	"xback/goods_srv/model"
	"xback/goods_srv/proto"
)

type GoodsServer struct {
	proto.UnimplementedGoodsServer
}

func ModelToResponse(goods model.Goods) proto.GoodsInfoResponse {
	return proto.GoodsInfoResponse{
		Id:              goods.ID,
		CategoryId:      goods.CategoryID,
		Name:            goods.Name,
		GoodsSn:         goods.GoodsSn,
		ClickNum:        goods.ClickNum,
		SoldNum:         goods.SoldNum,
		FavNum:          goods.FavNum,
		MarketPrice:     goods.MarketPrice,
		ShopPrice:       goods.ShopPrice,
		GoodsBrief:      goods.GoodsBrief,
		ShipFree:        goods.ShipFree,
		GoodsFrontImage: goods.GoodsFrontImage,
		IsNew:           goods.IsNew,
		IsHot:           goods.IsHot,
		OnSale:          goods.OnSale,
		DescImages:      goods.DescImages,
		Images:          goods.Images,
		Category: &proto.CategoryBriefInfoResponse{
			Id:   goods.Category.ID,
			Name: goods.Category.Name,
		},
		Brand: &proto.BrandInfoResponse{
			Id:   goods.Brands.ID,
			Name: goods.Brands.Name,
			Logo: goods.Brands.Logo,
		},
	}
}

func (s *GoodsServer) GoodsList(ctx context.Context, req *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
	goodsListResponse := &proto.GoodsListResponse{}

	// Create a bool query for Elasticsearch
	q := elastic.NewBoolQuery()

	if req.KeyWords != "" {
		q = q.Must(elastic.NewMultiMatchQuery(req.KeyWords, "name", "goods_brief"))
	}
	if req.IsHot {
		q = q.Filter(elastic.NewTermQuery("is_hot", req.IsHot))
	}
	if req.IsNew {
		q = q.Filter(elastic.NewTermQuery("is_new", req.IsNew))
	}
	if req.PriceMin > 0 {
		q = q.Filter(elastic.NewRangeQuery("shop_price").Gte(req.PriceMin))
	}
	if req.PriceMax > 0 {
		q = q.Filter(elastic.NewRangeQuery("shop_price").Lte(req.PriceMax))
	}
	if req.Brand > 0 {
		q = q.Filter(elastic.NewTermQuery("brands_id", req.Brand))
	}

	// Handle category filtering
	categoryIds := make([]interface{}, 0)
	if req.TopCategory > 0 {
		var category model.Category
		if result := global.DB.First(&category, req.TopCategory); result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "商品分类不存在")
		}

		var subQuery string
		if category.Level == 1 {
			subQuery = fmt.Sprintf("SELECT id FROM category WHERE parent_category_id IN (SELECT id FROM category WHERE parent_category_id=%d)", req.TopCategory)
		} else if category.Level == 2 {
			subQuery = fmt.Sprintf("SELECT id FROM category WHERE parent_category_id=%d", req.TopCategory)
		} else if category.Level == 3 {
			subQuery = fmt.Sprintf("SELECT id FROM category WHERE id=%d", req.TopCategory)
		}

		type Result struct {
			ID int32
		}
		var results []Result
		global.DB.Raw(subQuery).Scan(&results)
		for _, re := range results {
			categoryIds = append(categoryIds, re.ID)
		}

		// Add terms query for categories
		q = q.Filter(elastic.NewTermsQuery("category_id", categoryIds...))
	}

	// Pagination setup
	if req.Pages == 0 {
		req.Pages = 1
	}
	if req.PagePerNums > 100 {
		req.PagePerNums = 100
	} else if req.PagePerNums <= 0 {
		req.PagePerNums = 10
	}

	// Execute the Elasticsearch query
	result, err := global.EsClient.Search().
		Index(model.EsGoods{}.GetIndexName()).
		Query(q).
		From(int((req.Pages - 1) * req.PagePerNums)). // Adjust for page number
		Size(int(req.PagePerNums)).
		Do(context.Background())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Elasticsearch query failed: %v", err)
	}

	// Collect goods IDs from Elasticsearch results
	goodsIds := make([]int32, 0)
	goodsListResponse.Total = int32(result.Hits.TotalHits.Value)
	for _, value := range result.Hits.Hits {
		var goods model.EsGoods
		if err := json.Unmarshal(value.Source, &goods); err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to unmarshal Elasticsearch response: %v", err)
		}
		goodsIds = append(goodsIds, goods.ID)
	}

	// Fetch goods details from MySQL
	var goods []model.Goods
	if err := global.DB.Preload("Category").Preload("Brands").Find(&goods, goodsIds).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "MySQL query failed: %v", err)
	}

	// Convert goods models to response format
	for _, good := range goods {
		goodsInfoResponse := ModelToResponse(good)
		goodsListResponse.Data = append(goodsListResponse.Data, &goodsInfoResponse)
	}

	return goodsListResponse, nil
}

func (s *GoodsServer) NewHomeList(ctx context.Context, req *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
	goodsListResponse := &proto.GoodsListResponse{}

	var goods []model.Goods
	query := global.DB.Preload("Category").Preload("Brands").Where("is_new = ?", true)

	// 获取总数
	var total int64
	query.Model(&model.Goods{}).Count(&total)
	goodsListResponse.Total = int32(total)

	// 设置分页参数
	if req.Pages == 0 {
		req.Pages = 1
	}
	if req.PagePerNums > 100 {
		req.PagePerNums = 100
	} else if req.PagePerNums <= 0 {
		req.PagePerNums = 10
	}

	// 分页查询
	err := query.Limit(int(req.PagePerNums)).Offset(int((req.Pages - 1) * req.PagePerNums)).Find(&goods).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "MySQL query failed: %v", err)
	}

	for _, good := range goods {
		goodsInfoResponse := ModelToResponse(good) // 自定义的模型转换函数
		goodsListResponse.Data = append(goodsListResponse.Data, &goodsInfoResponse)
	}
	return goodsListResponse, nil
}

func (s *GoodsServer) HotHomeList(ctx context.Context, req *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
	goodsListResponse := &proto.GoodsListResponse{}

	var goods []model.Goods
	query := global.DB.Preload("Category").Preload("Brands").Where("is_hot = ?", true)

	// Query the sum of sold_num, click_num, and fav_num, sort by the sum in descending order, and limit to top 8
	query = query.Select("*, (sold_num + click_num + fav_num) as total_score").Order("total_score desc").Limit(5)

	err := query.Find(&goods).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "MySQL query failed: %v", err)
	}

	// Get total count
	var total int64
	query.Model(&model.Goods{}).Count(&total)
	goodsListResponse.Total = int32(total)

	for _, good := range goods {
		goodsInfoResponse := ModelToResponse(good)
		goodsListResponse.Data = append(goodsListResponse.Data, &goodsInfoResponse)
	}
	return goodsListResponse, nil
}

// 现在用户提交订单有多个商品，你得批量查询商品的信息吧
func (s *GoodsServer) BatchGetGoods(ctx context.Context, req *proto.BatchGoodsIdInfo) (*proto.GoodsListResponse, error) {
	goodsListResponse := &proto.GoodsListResponse{}
	var goods []model.Goods

	//调用where并不会真正执行sql 只是用来生成sql的 当调用find， first才会去执行sql，
	result := global.DB.Where(req.Id).Find(&goods)
	for _, good := range goods {
		goodsInfoResponse := ModelToResponse(good)
		goodsListResponse.Data = append(goodsListResponse.Data, &goodsInfoResponse)
	}
	goodsListResponse.Total = int32(result.RowsAffected)
	return goodsListResponse, nil
}

func (s *GoodsServer) GetGoodsDetail(ctx context.Context, req *proto.GoodInfoRequest) (*proto.GoodsInfoResponse, error) {
	var goods model.Goods

	if result := global.DB.Preload("Category").Preload("Brands").First(&goods, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}
	goodsInfoResponse := ModelToResponse(goods)
	return &goodsInfoResponse, nil
}

func (s *GoodsServer) CreateGoods(ctx context.Context, req *proto.CreateGoodsInfo) (*proto.GoodsInfoResponse, error) {
	var category model.Category
	if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var brand model.Brands
	if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}
	//先检查redis中是否有这个token
	//防止同一个token的数据重复插入到数据库中，如果redis中没有这个token则放入redis
	//这里没有看到图片文件是如何上传， 在微服务中 普通的文件上传已经不再使用，我们使用阿里云的oss
	goods := model.Goods{
		Brands:          brand,
		BrandsID:        brand.ID,
		Category:        category,
		CategoryID:      category.ID,
		Name:            req.Name,
		GoodsSn:         req.GoodsSn,
		MarketPrice:     req.MarketPrice,
		ShopPrice:       req.ShopPrice,
		GoodsBrief:      req.GoodsBrief,
		ShipFree:        req.ShipFree,
		Images:          req.Images,
		DescImages:      req.DescImages,
		GoodsFrontImage: req.GoodsFrontImage,
		IsNew:           req.IsNew,
		IsHot:           req.IsHot,
		OnSale:          req.OnSale,
	}

	//srv之间互相调用了
	tx := global.DB.Begin()
	result := tx.Save(&goods)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return &proto.GoodsInfoResponse{
		Id: goods.ID,
	}, nil
}

func (s *GoodsServer) DeleteGoods(ctx context.Context, req *proto.DeleteGoodsInfo) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Goods{BaseModel: model.BaseModel{ID: req.Id}}, req.Id); result.Error != nil {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateGoods(ctx context.Context, req *proto.CreateGoodsInfo) (*emptypb.Empty, error) {
	var goods model.Goods

	if result := global.DB.First(&goods, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}

	var category model.Category
	if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var brand model.Brands
	if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	goods = model.Goods{
		Brands:          brand,
		BrandsID:        brand.ID,
		Category:        category,
		CategoryID:      category.ID,
		Name:            req.Name,
		GoodsSn:         req.GoodsSn,
		MarketPrice:     req.MarketPrice,
		ShopPrice:       req.ShopPrice,
		GoodsBrief:      req.GoodsBrief,
		ShipFree:        req.ShipFree,
		Images:          req.Images,
		DescImages:      req.DescImages,
		GoodsFrontImage: req.GoodsFrontImage,
		IsNew:           req.IsNew,
		IsHot:           req.IsHot,
		OnSale:          req.OnSale,
	}

	tx := global.DB.Begin()
	result := tx.Save(&goods)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateGoodsStatus(ctx context.Context, req *proto.UpdateGoodsStatusRequest) (*emptypb.Empty, error) {
	var goods model.Goods

	if result := global.DB.First(&goods, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}
	log.Println(goods)
	goods.IsNew = req.IsNew
	goods.IsHot = req.IsHot
	goods.OnSale = req.OnSale

	if result := global.DB.Save(&goods); result.Error != nil {
		return nil, result.Error
	}

	return &emptypb.Empty{}, nil
}

// 搜索框接口
func (s *GoodsServer) SearchGoods(ctx context.Context, request *proto.SearchRequest) (*proto.GoodsListResponse, error) {
	goodsListResponse := &proto.GoodsListResponse{}

	var goods []model.Goods
	query := global.DB.Preload("Category").Preload("Brands").Where("name LIKE ?", "%"+request.Keyword+"%")

	// 获取总数
	var total int64
	query.Model(&model.Goods{}).Count(&total)
	goodsListResponse.Total = int32(total)

	// 查询所有符合条件的商品
	err := query.Find(&goods).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "MySQL query failed: %v", err)
	}

	for _, good := range goods {
		goodsInfoResponse := ModelToResponse(good)
		goodsListResponse.Data = append(goodsListResponse.Data, &goodsInfoResponse)
	}

	return goodsListResponse, nil
}
