package api

import "github.com/gin-gonic/gin"

func GetUserRouter(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": "0000",
		"data": map[string]interface{}{
			"home": "/home",
			"routes": []string{
				"/",
				"/about",
				"/home",
				"/user-center",
			},
		},
		"msg": "请求成功",
	})
}

func GetSuperRouter(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": "0000",
		"data": map[string]interface{}{
			"home": "/home",
			"routes": []string{
				"/",
				"/about",
				"/home",
				"/good",
				"/order",
				"/manage",
				"/manage/role",
				"/manage/role/*",
				"/manage/user",
				"/manage/user/:id",
				"/user-center",
			},
		},
		"msg": "请求成功",
	})
}

func GetAdminRoutes(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": "0000",
		"data": map[string]interface{}{
			"home": "/home",
			"routes": []string{
				"/",
				"/about",
				"/home",
				"/good",
				"/order",
				"/manage",
				"/manage/role",
				"/manage/user",
				"/manage/user/:id",
				"/user-center",
			},
		},
		"msg": "请求成功",
	})
}
