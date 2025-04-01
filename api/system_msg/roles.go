package system_msg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"soybean-admin-go/config"
	"soybean-admin-go/db/gen"
	"soybean-admin-go/db/model"
	"soybean-admin-go/utils/log"
	"strconv"
)

func GetAllRoles(ctx *gin.Context) {
	roles := gen.Q.Role
	var rolesRes []struct {
		ID       int64  `gorm:"column:id" json:"id"`
		RoleName string `gorm:"column:role_name" json:"roleName"`
		RoleCode string `gorm:"column:role_code" json:"roleCode"`
	}
	err := roles.WithContext(ctx).Select(roles.ID, roles.RoleName, roles.RoleCode).Scan(&rolesRes)
	if err != nil {
		config.Logger.Error("Failed to get roles info", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(500, gin.H{"error": "Failed to get roles info"})
		return
	}

	ctx.JSON(200, gin.H{
		"data": rolesRes,
	})
}

func AddRole(ctx *gin.Context) {
	var req RoleRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		config.Logger.Error("Failed to bind json", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(500, gin.H{"error": "Failed to bind json"})
		return
	}
	role := gen.Q.Role
	err = role.WithContext(ctx).Create(&model.Role{
		RoleName: req.RoleName,
		RoleCode: req.RoleCode,
		RoleDesc: req.RoleDesc,
		CreateBy: req.CreateBy,
		UpdateBy: req.UpdateBy,
		Status:   req.Status,
	})
	if err != nil {
		config.Logger.Error("Failed to insert role", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(500, gin.H{"error": "Failed to insert role"})
		return
	}
	ctx.JSON(200, gin.H{
		"code": "0000",
		"msg":  "请求成功",
	})
}

func UpdateRole(ctx *gin.Context) {
	var req RoleRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		config.Logger.Error("Failed to bind json", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(500, gin.H{"error": "Failed to bind json"})
		return
	}
	role := gen.Q.Role
	_, err = role.WithContext(ctx).Where(role.RoleCode.Eq(req.RoleCode)).Updates(&model.Role{
		RoleName: req.RoleName,
		RoleDesc: req.RoleDesc,
		UpdateBy: req.UpdateBy,
		Status:   req.Status,
	})
	if err != nil {
		config.Logger.Error("Failed to update role", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(500, gin.H{"error": "Failed to update role"})
		return
	}
	ctx.JSON(200, gin.H{
		"code": "0000",
		"msg":  "请求成功",
	})
}

func GetRoleList(ctx *gin.Context) {
	var (
		req   RoleQuery
		roles = gen.Q.Role
	)
	err := ctx.BindQuery(&req)
	fmt.Println(req)
	if err != nil {
		config.Logger.Error("Failed to bind query", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(500, gin.H{"error": "Failed to bind query"})
		return
	}
	roleQuery := roles.WithContext(ctx)
	if req.RoleCode != "" {
		roleQuery = roleQuery.Where(roles.RoleCode.Like("%" + req.RoleCode + "%"))
	}
	if req.RoleName != "" {
		roleQuery = roleQuery.Where(roles.RoleName.Like("%" + req.RoleName + "%"))
	}
	if req.Status != 0 {
		roleQuery = roleQuery.Where(roles.Status.Eq(int64(req.Status)))
	}
	list, err := roleQuery.Offset((req.Current - 1) * req.Size).Limit(req.Size).Find()
	if err != nil {
		config.Logger.Error("Failed to get roles info", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(500, gin.H{"error": "Failed to get roles info"})
		return
	}
	var records []RoleRequest
	for _, item := range list {
		records = append(records, RoleRequest{
			ID:       item.ID,
			RoleName: item.RoleName,
			RoleCode: item.RoleCode,
			RoleDesc: item.RoleDesc,
			CreateBy: item.CreateBy,
			UpdateBy: item.UpdateBy,
			Status:   item.Status,
		})
	}
	ctx.JSON(200, gin.H{
		"data": map[string]interface{}{
			"records": records,
			"total":   len(list),
			"size":    req.Size,
			"current": req.Current,
		},
		"code": "0000",
		"msg":  "请求成功",
	})

}

func DeleteRole(ctx *gin.Context) {

	role := gen.Q.Role
	roleidI64, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		config.Logger.Error("Failed to bind query", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(400, gin.H{"error": "Failed to bind query"})
		return
	}
	_, err = role.WithContext(ctx).Where(role.ID.Eq(roleidI64)).Delete()
	if err != nil {
		config.Logger.Error("Failed to delete role", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(500, gin.H{"error": "Failed to delete role"})
		return
	}
	ctx.JSON(200, gin.H{
		"code": "0000",
		"msg":  "请求成功",
	})
}

type RoleQuery struct {
	RoleCode string `form:"roleCode"`
	RoleName string `form:"roleName"`
	Current  int    `form:"current"`
	Size     int    `form:"size"`
	Status   int    `form:"status"`
}
type RoleRequest struct {
	ID       int64  `json:"id"`
	RoleName string `json:"roleName"`
	RoleCode string `json:"roleCode"`
	RoleDesc string `json:"roleDesc"`
	CreateBy string `json:"createBy"`
	UpdateBy string `json:"updateBy"`
	Status   int64  `json:"status"`
}
