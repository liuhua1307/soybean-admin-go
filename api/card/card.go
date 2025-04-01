package card

import "github.com/gin-gonic/gin"

// GetCardList 获取卡片列表
func GetCardList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": map[string]interface{}{
			"customers":           101,
			"orders":              23,
			"transactionQuantity": 34,
			"transactionVolume":   5,
		},
		"code": "0000",
		"msg":  "success",
	})
}

func GetLineList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": map[string]interface{}{
			"success": []int{
				2208,
				2016,
				2916,
				4512,
				8281,
				2008,
				1963,
				2367,
				2956,
				678},
			"total": []int{
				4623,
				6145,
				6268,
				6411,
				1890,
				4251,
				2978,
				3880,
				3606,
				4311},
		},
		"code": "0000",
		"msg":  "Success",
	})
}

func GetPieList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": gin.H{
			"data": []gin.H{
				{
					"name":  "数码产品",
					"value": 20,
				},
				{
					"name":  "食品",
					"value": 10,
				},
				{
					"name":  "家电",
					"value": 40,
				},
				{
					"name":  "服装",
					"value": 30,
				},
			},
		},
		"code": "0000",
		"msg":  "",
	})
}
