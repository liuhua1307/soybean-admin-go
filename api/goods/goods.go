package goods

import (
	"github.com/gin-gonic/gin"
	"soybean-admin-go/config"
	"soybean-admin-go/db/gen"
	"soybean-admin-go/db/model"
	"soybean-admin-go/utils/log"
	"strconv"
)

func GetGoodsList(ctx *gin.Context) {
	var (
		querys GoodsQuery
		goods  = gen.Q.Good
	)
	if err := ctx.ShouldBindQuery(&querys); err != nil {
		ctx.JSON(400, gin.H{"error": "bind query error"})
		return
	}
	goodsQuery := goods.WithContext(ctx)
	if querys.Name != "" {
		goodsQuery = goodsQuery.Where(goods.Name.Like("%" + querys.Name + "%"))
	}
	if querys.Repo != "" {
		goodsQuery = goodsQuery.Where(goods.Repo.Like("%" + querys.Repo + "%"))
	}
	if querys.Class != "" {
		goodsQuery = goodsQuery.Where(goods.Class.Like("%" + querys.Class + "%"))
	}

	find, err := goodsQuery.Offset((querys.Current - 1) * querys.Size).Limit(querys.Size).Find()
	if err != nil {
		config.Logger.Error("find goods error", log.Field{Key: "error", Value: err})
		ctx.JSON(500, gin.H{"error": "find goods error"})
		return
	}
	ctx.JSON(200, gin.H{
		"data": map[string]interface{}{
			"total":   len(find),
			"records": find,
			"size":    querys.Size,
			"current": querys.Current,
		},
		"code": "0000",
		"msg":  "success",
	})
}

func AddGoods(ctx *gin.Context) {
	var (
		goods = gen.Q.Good
	)
	var goodsInfo GoodsInfo
	if err := ctx.ShouldBindJSON(&goodsInfo); err != nil {
		ctx.JSON(400, gin.H{"error": "bind json error"})
		return
	}
	err := goods.WithContext(ctx).Create(&model.Good{
		Name:      goodsInfo.Name,
		Repo:      goodsInfo.Repo,
		Class:     goodsInfo.Class,
		Inventory: goodsInfo.Inventory,
		Weight:    goodsInfo.Weight,
		Desc:      goodsInfo.Desc,
	})
	if err != nil {
		config.Logger.Error("insert goods error", log.Field{Key: "error", Value: err})
		ctx.JSON(500, gin.H{"error": "insert goods error"})
		return
	}
	ctx.JSON(200, gin.H{"code": "0000", "msg": "success"})
}

func UpdateGoods(ctx *gin.Context) {
	var (
		goods = gen.Q.Good
	)
	var goodsInfo GoodsInfo
	if err := ctx.ShouldBindJSON(&goodsInfo); err != nil {
		ctx.JSON(400, gin.H{"error": "bind json error"})
		return
	}
	_, err := goods.WithContext(ctx).Where(goods.ID.Eq(goodsInfo.ID)).Updates(&model.Good{
		Repo:      goodsInfo.Repo,
		Class:     goodsInfo.Class,
		Inventory: goodsInfo.Inventory,
		Weight:    goodsInfo.Weight,
		Desc:      goodsInfo.Desc,
		Name:      goodsInfo.Name,
	})
	if err != nil {
		config.Logger.Error("update goods error", log.Field{Key: "error", Value: err})
		ctx.JSON(500, gin.H{"error": "update goods error"})
		return
	}
	ctx.JSON(200, gin.H{"code": "0000", "msg": "success"})
}

func DeleteGoods(ctx *gin.Context) {
	var (
		goods  = gen.Q.Good
		goodId = ctx.Param("id")
	)
	goodsIdI64, err := strconv.ParseInt(goodId, 10, 64)
	if err != nil {
		config.Logger.Error("delete goods error", log.Field{Key: "error", Value: err})
		ctx.JSON(500, gin.H{"error": "delete goods error"})
		return
	}

	_, err = goods.WithContext(ctx).Where(goods.ID.Eq(goodsIdI64)).Delete()
	if err != nil {
		config.Logger.Error("delete goods error", log.Field{Key: "error", Value: err})
		ctx.JSON(500, gin.H{"error": "delete goods error"})
		return
	}
	ctx.JSON(200, gin.H{"code": "0000", "msg": "success"})

}

type GoodsQuery struct {
	Current int    `form:"current"`
	Size    int    `form:"size"`
	Name    string `form:"name"`
	Repo    string `form:"repo"`
	Class   string `form:"class"`
}

// tag
type GoodsInfo struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Repo      string `json:"repo"`
	Class     string `json:"class"`
	Inventory string `json:"inventory"`
	Weight    string `json:"weight"`
	Desc      string `json:"desc"`
}
