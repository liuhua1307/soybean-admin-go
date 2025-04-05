package router

import (
	"github.com/gin-gonic/gin"
	"soybean-admin-go/api"
	"soybean-admin-go/api/card"
	"soybean-admin-go/api/goods"
	"soybean-admin-go/api/location"
	"soybean-admin-go/api/orders"
	"soybean-admin-go/api/system_msg"
)

func Init(r *gin.Engine) {
	// 使用 cookie 存储会话数据
	r.POST("/api/auth/login", api.Login)
	r.GET("/api/auth/refreshToken", api.RefreshToken)
	g := r.Group("/api", api.AuthMiddleware())
	{
		g.GET("/auth/getUserInfo", api.GetUserInfo)

		g.GET("/route/getUserRouter", api.GetSuperRouter)

		g.GET("/systemManage/getAllRoles", system_msg.GetAllRoles)
		g.POST("/systemManage/role", system_msg.AddRole)
		g.GET("/systemManage/getRoleList", system_msg.GetRoleList)
		g.GET("/systemManage/getUserList", system_msg.GetUserList)
		g.GET("/systemManage/getUserInfo/:id", system_msg.GetUserInfo)
		g.POST("/systemManage/user", system_msg.CreateUser)
		g.PUT("/systemManage/user", system_msg.UpdateUser)
		g.DELETE("/systemManage/user/:id", system_msg.DeleteUser)
		g.PUT("/systemManage/role", system_msg.UpdateRole)
		g.DELETE("/systemManage/role/:id", system_msg.DeleteRole)

		g.GET("/goodManage/getGoodList", goods.GetGoodsList)
		g.POST("/goodManage/good", goods.AddGoods)
		g.PUT("/goodManage/good", goods.UpdateGoods)
		g.DELETE("/goodManage/good/:id", goods.DeleteGoods)

		g.GET("/orderManage/getOrderList", orders.GetOrderList)
		g.POST("/orderManage/order", orders.AddOrder)
		g.PUT("/orderManage/order", orders.UpdateOrder)
		g.DELETE("/orderManage/order/:id", orders.DeleteOrder)
		g.GET("/orderManage/getRoute/:id", location.GotLocation)

		g.GET("/route/getReactUserRoutes", api.GetSuperRouter)
		gr := g.Group("/data")
		gr.GET("/card", card.GetCardList)
		gr.GET("/pie", card.GetPieList)
		gr.GET("/line", card.GetLineList)
	}
}
