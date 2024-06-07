package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"

	"xback/goods_srv/global"
	"xback/goods_srv/model"
	"xback/goods_srv/proto"
)

func (s *GoodsServer) CategoryBrandList(ctx context.Context, req *proto.CategoryBrandFilterRequest) (*proto.CategoryBrandListResponse, error) {
	var categoryBrands []model.GoodsCategoryBrand
	categoryBrandListResponse := proto.CategoryBrandListResponse{}

	var total int64
	global.DB.Model(&model.GoodsCategoryBrand{}).Count(&total)
	categoryBrandListResponse.Total = int32(total)

	global.DB.Preload("Category").Preload("Brands").Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&categoryBrands)

	var categoryResponses []*proto.CategoryBrandResponse
	for _, categoryBrand := range categoryBrands {
		categoryResponses = append(categoryResponses, &proto.CategoryBrandResponse{
			Id: categoryBrand.ID,
			Category: &proto.CategoryInfoResponse{
				Id:             categoryBrand.Category.ID,
				Name:           categoryBrand.Category.Name,
				Level:          categoryBrand.Category.Level,
				IsTab:          categoryBrand.Category.IsTab,
				ParentCategory: categoryBrand.Category.ParentCategoryID,
			},
			Brand: &proto.BrandInfoResponse{
				Id:   categoryBrand.Brands.ID,
				Name: categoryBrand.Brands.Name,
				Logo: categoryBrand.Brands.Logo,
			},
		})
	}

	categoryBrandListResponse.Data = categoryResponses
	return &categoryBrandListResponse, nil
}

func (s *GoodsServer) GetCategoryBrandList(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.BrandListResponse, error) {
	brandListResponse := proto.BrandListResponse{}

	var category model.Category
	if result := global.DB.Find(&category, req.Id).First(&category); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var categoryBrands []model.GoodsCategoryBrand
	if result := global.DB.Preload("Brands").Where(&model.GoodsCategoryBrand{CategoryID: req.Id}).Find(&categoryBrands); result.RowsAffected > 0 {
		brandListResponse.Total = int32(result.RowsAffected)
	}

	var brandInfoResponses []*proto.BrandInfoResponse
	for _, categoryBrand := range categoryBrands {
		brandInfoResponses = append(brandInfoResponses, &proto.BrandInfoResponse{
			Id:   categoryBrand.Brands.ID,
			Name: categoryBrand.Brands.Name,
			Logo: categoryBrand.Brands.Logo,
		})
	}

	brandListResponse.Data = brandInfoResponses

	return &brandListResponse, nil
}

func (s *GoodsServer) CreateCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*proto.CategoryBrandResponse, error) {
	log.Printf("Received request: %+v\n", req)

	var category model.Category
	if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
		log.Printf("Category not found: %d\n", req.CategoryId)
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var brand model.Brands
	if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
		log.Printf("Brand not found: %d\n", req.BrandId)
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	categoryBrand := model.GoodsCategoryBrand{
		CategoryID: req.CategoryId,
		BrandsID:   req.BrandId,
	}

	if result := global.DB.Save(&categoryBrand); result.Error != nil {
		log.Printf("Error saving category brand: %v\n", result.Error)
		return nil, status.Errorf(codes.Internal, "创建分类品牌失败")
	}

	log.Printf("Created category brand with ID: %d\n", categoryBrand.ID)
	return &proto.CategoryBrandResponse{Id: categoryBrand.ID}, nil
}

func (s *GoodsServer) DeleteCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*emptypb.Empty, error) {
	if result := global.DB.Unscoped().Delete(&model.GoodsCategoryBrand{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌分类不存在")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*emptypb.Empty, error) {
	var categoryBrand model.GoodsCategoryBrand

	if result := global.DB.First(&categoryBrand, req.Id); result.RowsAffected == 0 {
		log.Printf("CategoryBrand not found: %d\n", req.Id)
		return nil, status.Errorf(codes.InvalidArgument, "品牌分类不存在")
	}

	var category model.Category
	if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
		log.Printf("Category not found: %d\n", req.CategoryId)
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var brand model.Brands
	if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
		log.Printf("Brand not found: %d\n", req.BrandId)
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	categoryBrand.CategoryID = req.CategoryId
	categoryBrand.BrandsID = req.BrandId

	if result := global.DB.Save(&categoryBrand); result.Error != nil {
		log.Printf("Error saving category brand: %v\n", result.Error)
		return nil, status.Errorf(codes.Internal, "更新分类品牌失败")
	}

	log.Printf("Updated category brand with ID: %d\n", categoryBrand.ID)
	return &emptypb.Empty{}, nil
}
