package goods

import (
	"context"
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"strconv"
	"strings"
	"xfront/goods-web/forms"
	"xfront/goods-web/proto"

	"xfront/goods-web/global"
)

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	//将grpc的code转换成http的状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg:": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Code(),
				})
			}
			return
		}
	}
}

func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}

func List(ctx *gin.Context) {
	request := &proto.GoodsFilterRequest{}

	priceMin := ctx.DefaultQuery("pmin", "0")
	priceMinInt, _ := strconv.Atoi(priceMin)
	request.PriceMin = int32(priceMinInt)

	priceMax := ctx.DefaultQuery("pmax", "0")
	priceMaxInt, _ := strconv.Atoi(priceMax)
	request.PriceMax = int32(priceMaxInt)

	if request.PriceMin > request.PriceMax && request.PriceMax != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "价格区间选择错误",
		})
		return
	}

	isHot := ctx.DefaultQuery("hot", "0")
	request.IsHot = (isHot == "1")

	isNew := ctx.DefaultQuery("new", "0")
	request.IsNew = (isNew == "1")

	isTab := ctx.DefaultQuery("sale", "0")
	request.IsTab = (isTab == "1")

	categoryId := ctx.DefaultQuery("c", "0")
	categoryIdInt, _ := strconv.Atoi(categoryId)
	request.TopCategory = int32(categoryIdInt)

	pages := ctx.DefaultQuery("pn", "1")
	pagesInt, _ := strconv.Atoi(pages)
	request.Pages = int32(pagesInt)

	perNums := ctx.DefaultQuery("pnum", "10")
	perNumsInt, _ := strconv.Atoi(perNums)
	request.PagePerNums = int32(perNumsInt)

	keywords := ctx.DefaultQuery("q", "")
	request.KeyWords = keywords

	brandId := ctx.DefaultQuery("b", "0")
	brandIdInt, _ := strconv.Atoi(brandId)
	request.Brand = int32(brandIdInt)

	e, b := sentinel.Entry("goods-list", sentinel.WithTrafficType(base.Inbound))
	if b != nil {
		ctx.JSON(http.StatusTooManyRequests, gin.H{
			"msg": "请求过于频繁，请稍后重试",
		})
		return
	}
	r, err := global.GoodsSrvClient.GoodsList(context.WithValue(context.Background(), "ginContext", ctx), request)
	if err != nil {
		zap.S().Errorw("[List] 查询 【商品列表】失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	e.Exit()

	reMap := map[string]interface{}{
		"total": r.Total,
	}
	goodsList := make([]interface{}, 0)
	for _, value := range r.Data {
		goodsList = append(goodsList, map[string]interface{}{
			"id":          value.Id,
			"name":        value.Name,
			"goods_brief": value.GoodsBrief,
			"goods_sn":    value.GoodsSn,
			"sold_num":    value.SoldNum,  // 销售量
			"click_num":   value.ClickNum, // 点击
			"fav_num":     value.FavNum,   // 收藏
			"desc":        value.GoodsDesc,
			"ship_free":   value.ShipFree,
			"images":      value.Images,
			"desc_images": value.DescImages,
			"front_image": value.GoodsFrontImage,
			"shop_price":  value.ShopPrice,
			"category": map[string]interface{}{
				"id":   value.Category.Id,
				"name": value.Category.Name,
			},
			"brand": map[string]interface{}{
				"id":   value.Brand.Id,
				"name": value.Brand.Name,
				"logo": value.Brand.Logo,
			},
			"market_price": value.MarketPrice,
			"stocks":       2000, // 默认值 2000
			"is_hot":       value.IsHot,
			"is_new":       value.IsNew,
			"on_sale":      value.OnSale,
		})
	}
	reMap["data"] = goodsList

	ctx.JSON(http.StatusOK, reMap)
}

func NewHomeList(ctx *gin.Context) {
	request := &proto.GoodsFilterRequest{}

	pnumStr := ctx.DefaultQuery("pnum", "0")
	pnum, err := strconv.Atoi(pnumStr)
	if err != nil || pnum <= 0 {
		pnum = 10 // 默认值
	}

	// 设置请求的分页数量
	request.PagePerNums = int32(pnum)
	request.Pages = 1 // 只需要第一页的数据

	e, b := sentinel.Entry("goods-list", sentinel.WithTrafficType(base.Inbound))
	if b != nil {
		ctx.JSON(http.StatusTooManyRequests, gin.H{
			"msg": "请求过于频繁，请稍后重试",
		})
		return
	}

	r, err := global.GoodsSrvClient.NewHomeList(context.WithValue(context.Background(), "ginContext", ctx), request)
	if err != nil {
		zap.S().Errorw("[List] 查询 【商品列表】失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	e.Exit()

	reMap := map[string]interface{}{
		"total": r.Total,
	}
	goodsList := make([]interface{}, 0)
	for _, value := range r.Data {
		goodsList = append(goodsList, map[string]interface{}{
			"id":          value.Id,
			"name":        value.Name,
			"goods_brief": value.GoodsBrief,
			"goods_sn":    value.GoodsSn,
			"sold_num":    value.SoldNum,  // 销售量
			"click_num":   value.ClickNum, // 点击
			"fav_num":     value.FavNum,   // 收藏
			"desc":        value.GoodsDesc,
			"ship_free":   value.ShipFree,
			"images":      value.Images,
			"desc_images": value.DescImages,
			"front_image": value.GoodsFrontImage,
			"shop_price":  value.ShopPrice,
			"category": map[string]interface{}{
				"id":   value.Category.Id,
				"name": value.Category.Name,
			},
			"brand": map[string]interface{}{
				"id":   value.Brand.Id,
				"name": value.Brand.Name,
				"logo": value.Brand.Logo,
			},
			"market_price": value.MarketPrice,
			"stocks":       2000, // 默认值 2000
			"is_hot":       value.IsHot,
			"is_new":       value.IsNew,
			"on_sale":      value.OnSale,
		})
	}
	reMap["data"] = goodsList
	ctx.JSON(http.StatusOK, reMap)
}

func HotHomeList(ctx *gin.Context) {
	request := &proto.GoodsFilterRequest{}

	e, b := sentinel.Entry("goods-list", sentinel.WithTrafficType(base.Inbound))
	if b != nil {
		ctx.JSON(http.StatusTooManyRequests, gin.H{
			"msg": "请求过于频繁，请稍后重试",
		})
		return
	}

	r, err := global.GoodsSrvClient.HotHomeList(context.WithValue(context.Background(), "ginContext", ctx), request)
	if err != nil {
		zap.S().Errorw("[List] 查询 【商品列表】失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	e.Exit()

	// Construct the response data
	reMap := map[string]interface{}{
		"total": r.Total,
	}
	goodsList := make([]interface{}, 0)
	for _, value := range r.Data {
		goodsList = append(goodsList, map[string]interface{}{
			"id":        value.Id,
			"name":      value.Name,
			"sold_num":  value.SoldNum,
			"click_num": value.ClickNum,
			"fav_num":   value.FavNum,
			"desc":      value.GoodsDesc,
		})
	}
	reMap["data"] = goodsList
	ctx.JSON(http.StatusOK, reMap)
}

func New(ctx *gin.Context) {
	goodsForm := forms.GoodsForm{}
	if err := ctx.ShouldBindJSON(&goodsForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	goodsClient := global.GoodsSrvClient
	rsp, err := goodsClient.CreateGoods(context.Background(), &proto.CreateGoodsInfo{
		Name:            goodsForm.Name,        // 商品名称
		GoodsSn:         goodsForm.GoodsSn,     // 商品编号
		Stocks:          goodsForm.Stocks,      // 商品库存
		MarketPrice:     goodsForm.MarketPrice, // 市场价格
		ShopPrice:       goodsForm.ShopPrice,   // 商店价格
		GoodsBrief:      goodsForm.GoodsBrief,  // 商品简介
		ShipFree:        *goodsForm.ShipFree,   // 是否包邮
		Images:          goodsForm.Images,      // 商品图片
		DescImages:      goodsForm.DescImages,  // 商品描述图片
		GoodsFrontImage: goodsForm.FrontImage,  // 商品封面图片
		CategoryId:      goodsForm.CategoryId,  // 商品分类ID
		BrandId:         goodsForm.Brand,       // 商品品牌ID
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	//如何设置库存
	//TODO 商品的库存 - 分布式事务
	_, err = global.InventorySrvClient.SetInv(context.WithValue(context.Background(), "ginContext", ctx), &proto.GoodsInvInfo{
		GoodsId: rsp.Id,
		Num:     goodsForm.Stocks,
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, rsp)
}

func Detail(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	r, err := global.GoodsSrvClient.GetGoodsDetail(context.WithValue(context.Background(), "ginContext", ctx), &proto.GoodInfoRequest{
		Id: int32(i),
	})
	if err != nil {
		fmt.Println("err:", err)
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	invr, err := global.InventorySrvClient.InvDetail(context.WithValue(context.Background(), "ginContext", ctx), &proto.GoodsInvInfo{
		GoodsId: r.Id,
	})
	if err != nil {
		fmt.Println("invr,err:", err)
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	rsp := map[string]interface{}{
		"id":          r.Id,              // 商品ID
		"name":        r.Name,            // 商品名称
		"goods_brief": r.GoodsBrief,      // 商品简介
		"goods_sn":    r.GoodsSn,         //商品编号
		"sold_num":    r.SoldNum,         //销量
		"click_num":   r.ClickNum,        //点击
		"fav_num":     r.FavNum,          //收藏
		"desc":        r.GoodsDesc,       // 商品描述
		"ship_free":   r.ShipFree,        // 是否包邮
		"images":      r.Images,          // 商品图片
		"desc_images": r.DescImages,      // 商品描述图片
		"front_image": r.GoodsFrontImage, // 商品封面图片
		"shop_price":  r.ShopPrice,       // 商品价格
		"category": map[string]interface{}{ // 商品分类信息
			"id":   r.Category.Id,   // 商品分类ID
			"name": r.Category.Name, // 商品分类名称
		},
		"brand": map[string]interface{}{ // 商品品牌信息
			"id":   r.Brand.Id,   // 商品品牌ID
			"name": r.Brand.Name, // 商品品牌名称
			"logo": r.Brand.Logo, // 商品品牌Logo
		},
		"stocks":       invr.Num, // 跨微服务，商品库存
		"market_price": r.MarketPrice,
		"is_hot":       r.IsHot,  // 是否热销
		"is_new":       r.IsNew,  // 是否新品
		"on_sale":      r.OnSale, // 是否在售
	}
	ctx.JSON(http.StatusOK, rsp)
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsSrvClient.DeleteGoods(context.Background(), &proto.DeleteGoodsInfo{Id: int32(i)})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	ctx.Status(http.StatusOK)
	return
}

func Stocks(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	//TODO 商品的库存
	return
}

func UpdateStatus(ctx *gin.Context) {
	goodsStatusForm := forms.GoodsStatusForm{}
	if err := ctx.ShouldBindJSON(&goodsStatusForm); err != nil {
		log.Println("绑定JSON错误:", err)
		HandleValidatorError(ctx, err)
		return
	}
	log.Printf("接收到的数据: %+v\n", goodsStatusForm)

	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	if _, err = global.GoodsSrvClient.UpdateGoodsStatus(context.Background(), &proto.UpdateGoodsStatusRequest{
		Id:     int32(i),
		IsHot:  *goodsStatusForm.IsHot,
		IsNew:  *goodsStatusForm.IsNew,
		OnSale: *goodsStatusForm.OnSale,
	}); err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "修改成功",
	})
}

func Update(ctx *gin.Context) {
	goodsForm := forms.GoodsForm{}
	if err := ctx.ShouldBindJSON(&goodsForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if _, err = global.GoodsSrvClient.UpdateGoods(context.Background(), &proto.CreateGoodsInfo{
		Id:              int32(i),
		Name:            goodsForm.Name,
		GoodsSn:         goodsForm.GoodsSn,
		Stocks:          goodsForm.Stocks,
		MarketPrice:     goodsForm.MarketPrice,
		ShopPrice:       goodsForm.ShopPrice,
		GoodsBrief:      goodsForm.GoodsBrief,
		ShipFree:        *goodsForm.ShipFree,
		Images:          goodsForm.Images,
		DescImages:      goodsForm.DescImages,
		GoodsFrontImage: goodsForm.FrontImage,
		CategoryId:      goodsForm.CategoryId,
		BrandId:         goodsForm.Brand,
	}); err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}

// 搜索框搜索
func SearchGoods(ctx *gin.Context) {
	keyword := ctx.Query("keyword")

	request := &proto.SearchRequest{
		Keyword: keyword,
	}

	log.Println(request.Keyword)

	r, err := global.GoodsSrvClient.SearchGoods(context.WithValue(context.Background(), "ginContext", ctx), request)
	if err != nil {
		zap.S().Errorw("[Search] 查询 【商品列表】失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	reMap := map[string]interface{}{
		"total": r.Total,
	}
	goodsList := make([]interface{}, 0)
	for _, value := range r.Data {
		goodsList = append(goodsList, map[string]interface{}{
			"id":          value.Id,
			"name":        value.Name,
			"goods_brief": value.GoodsBrief,
			"goods_sn":    value.GoodsSn,
			"sold_num":    value.SoldNum,  // 销售量
			"click_num":   value.ClickNum, // 点击
			"fav_num":     value.FavNum,   // 收藏
			"desc":        value.GoodsDesc,
			"ship_free":   value.ShipFree,
			"images":      value.Images,
			"desc_images": value.DescImages,
			"front_image": value.GoodsFrontImage,
			"shop_price":  value.ShopPrice,
			"category": map[string]interface{}{
				"id":   value.Category.Id,
				"name": value.Category.Name,
			},
			"brand": map[string]interface{}{
				"id":   value.Brand.Id,
				"name": value.Brand.Name,
				"logo": value.Brand.Logo,
			},
			"market_price": value.MarketPrice,
			"stocks":       2000, // 默认值 2000
			"is_hot":       value.IsHot,
			"is_new":       value.IsNew,
			"on_sale":      value.OnSale,
		})
	}
	reMap["data"] = goodsList
	ctx.JSON(http.StatusOK, reMap)
}
