package system_msg

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"soybean-admin-go/config"
	"soybean-admin-go/db/gen"
	"soybean-admin-go/db/model"
	"soybean-admin-go/utils/log"
	"strconv"
	"sync"
	"time"
)

func CreateUser(ctx *gin.Context) {
	var (
		userDto UserDto
		tx      = gen.Q.Begin()
	)
	err := ctx.BindJSON(&userDto)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind userDto"})
		return
	}

	modelUser := &model.User{
		CreateBy:   userDto.CreateBy,
		UpdateBy:   userDto.UpdateBy,
		Status:     userDto.Status,
		UserName:   userDto.UserName,
		UserGender: userDto.UserGender,
		NickName:   userDto.NickName,
		UserPhone:  userDto.UserPhone,
		UserEmail:  userDto.UserEmail,
	}
	err = tx.User.WithContext(ctx).Create(modelUser)
	if err != nil {
		tx.Rollback()
		config.Logger.Error("Failed to create user", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	IDs, err := tx.Role.WithContext(ctx).Where(tx.Role.RoleCode.In(userDto.Roles...)).Select(tx.Role.ID).Find()
	if err != nil || len(IDs) != len(userDto.Roles) {
		tx.Rollback()
		config.Logger.Error("Failed to get role info", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get role info"})
		return
	}
	wg := &sync.WaitGroup{}
	once := &sync.Once{}
	wg.Add(len(IDs))
	errChan := make(chan error, len(IDs))
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	for _, ID := range IDs {
		go func(ID *model.Role) {
			defer wg.Done()
			select {
			case <-cancelCtx.Done():
				return
			default:
				err = tx.UserRole.WithContext(ctx).Create(&model.UserRole{
					UserID: modelUser.ID,
					RoleID: ID.ID,
				})
				if err != nil {
					once.Do(func() {
						cancelFunc()
						errChan <- err
						tx.Rollback()
					})
				}
			}

		}(ID)
	}
	wg.Wait()
	close(errChan)
	for errs := range errChan {
		if errs != nil {
			config.Logger.Error("Failed to create user role", log.Field{
				Key:   "error",
				Value: errs,
			})
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user role"})
		}
	}

	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{"msg": "success", "code": "0000"})
}

func UpdateUser(ctx *gin.Context) {
	var (
		userDto UserDto
		tx      = gen.Q.Begin()
	)
	err := ctx.BindJSON(&userDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind userDto"})
		return
	}
	var password string
	err = tx.User.WithContext(ctx).Where(tx.User.UserName.Eq(userDto.UserName)).Select(tx.User.Password).Scan(&password)
	if err != nil {
		config.Logger.Error("Failed to get password", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get password"})
		return
	}
	modelUser := &model.User{
		CreateBy:   userDto.CreateBy,
		UpdateBy:   userDto.UpdateBy,
		Status:     userDto.Status,
		UserName:   userDto.UserName,
		UserGender: userDto.UserGender,
		NickName:   userDto.NickName,
		UserPhone:  userDto.UserPhone,
		UserEmail:  userDto.UserEmail,
		Password:   password,
	}
	_, err = tx.User.WithContext(ctx).Where(tx.User.ID.Eq(userDto.ID)).Updates(modelUser)
	if err != nil {
		tx.Rollback()
		config.Logger.Error("Failed to update user", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	// 先删了再加
	_, err = tx.UserRole.WithContext(ctx).Where(tx.UserRole.UserID.Eq(userDto.ID)).Delete()
	if err != nil {
		tx.Rollback()
		config.Logger.Error("Failed to delete user role", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user role"})
		return
	}
	IDs, err := tx.Role.WithContext(ctx).Where(tx.Role.RoleCode.In(userDto.Roles...)).Select(tx.Role.ID).Find()
	if err != nil || len(IDs) != len(userDto.Roles) {
		tx.Rollback()
		config.Logger.Error("Failed to get role info", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get role info"})
		return
	}
	wg := &sync.WaitGroup{}
	once := &sync.Once{}
	wg.Add(len(IDs))
	errChan := make(chan error, len(IDs))
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	for _, ID := range IDs {
		go func(ID *model.Role) {
			defer wg.Done()
			select {
			case <-cancelCtx.Done():
				return
			default:
				err = tx.UserRole.WithContext(ctx).Create(&model.UserRole{
					UserID: userDto.ID,
					RoleID: ID.ID,
				})
				if err != nil {
					once.Do(func() {
						cancelFunc()
						errChan <- err
						tx.Rollback()
					})
				}
			}

		}(ID)
	}
	wg.Wait()
	close(errChan)
	for errs := range errChan {
		if errs != nil {
			config.Logger.Error("Failed to create user role", log.Field{
				Key:   "error",
				Value: errs,
			})
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user role"})
		}
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{"msg": "success", "code": "0000"})
}

func DeleteUser(ctx *gin.Context) {

	var (
		tx = gen.Q.Begin()
	)
	userId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind userId"})
		return
	}
	_, err = tx.User.WithContext(ctx).Where(tx.User.ID.Eq(userId)).Delete()
	if err != nil {
		tx.Rollback()
		config.Logger.Error("Failed to delete user", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	_, err = tx.UserRole.WithContext(ctx).Where(tx.UserRole.UserID.Eq(userId)).Delete()
	if err != nil {
		tx.Rollback()
		config.Logger.Error("Failed to delete user role", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user role"})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{"msg": "success", "code": "0000"})
}

func GetUserList(ctx *gin.Context) {
	var (
		querys UserListQueryParams
		users  = gen.Q.User
	)
	err := ctx.BindQuery(&querys)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "bind error",
		})
	}
	fmt.Println(querys)
	userQuery := users.WithContext(ctx)
	if querys.NickName != "" {
		userQuery = userQuery.Where(users.NickName.Like("%" + querys.NickName + "%"))
	}
	if querys.UserName != "" {
		userQuery = userQuery.Where(users.UserName.Like("%" + querys.UserName + "%"))
	}
	if querys.UserGender != "" {
		userQuery = userQuery.Where(users.UserGender.Eq(querys.UserGender))
	}
	if querys.UserPhone != "" {
		userQuery = userQuery.Where(users.UserPhone.Like("%" + querys.UserPhone + "%"))
	}
	if querys.UserEmail != "" {
		userQuery = userQuery.Where(users.UserEmail.Like("%" + querys.UserEmail + "%"))
	}
	if querys.UserStatus != 0 {
		userQuery = userQuery.Where(users.Status.Eq(querys.UserStatus))
	}
	usersInfo, err := userQuery.Preload(users.Roles).Limit(querys.Size).Offset((querys.Current - 1) * querys.Size).Find()
	if err != nil {
		config.Logger.Error("Failed to get user list", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user list"})
		return
	}
	var userDtos []UserDto
	for _, user := range usersInfo {
		var roleCode []string
		for _, role := range user.Roles {
			roleCode = append(roleCode, role.RoleCode)
		}
		userDtos = append(userDtos, UserDto{
			ID:         user.ID,
			CreateBy:   user.CreateBy,
			UpdateBy:   user.UpdateBy,
			Status:     user.Status,
			UserName:   user.UserName,
			UserGender: user.UserGender,
			NickName:   user.NickName,
			UserPhone:  user.UserPhone,
			UserEmail:  user.UserEmail,
			CreateTime: user.CreatedAt,
			UpdateTime: user.CreatedAt,
			Roles:      roleCode,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"records": userDtos,
			"total":   len(userDtos),
			"size":    querys.Size,
			"current": querys.Current,
		},
		"code": "0000",
		"msg":  "请求成功",
	})
}

func GetUserInfo(ctx *gin.Context) {
	var (
		userId = ctx.Param("id")
		query  = gen.Q
	)
	userIdi64, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind userId"})
		return
	}
	userMsgInfo, err := query.WithContext(ctx).User.Preload(query.User.Roles).Where(query.User.ID.Eq(userIdi64)).First()
	if err != nil {
		config.Logger.Error("Failed to get user info", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}
	var roleNames []string
	for _, role := range userMsgInfo.Roles {
		roleNames = append(roleNames, role.RoleCode)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": UserDto{
			ID:         userMsgInfo.ID,
			CreateBy:   userMsgInfo.CreateBy,
			UpdateBy:   userMsgInfo.UpdateBy,
			Status:     userMsgInfo.Status,
			UserName:   userMsgInfo.UserName,
			UserGender: userMsgInfo.UserGender,
			NickName:   userMsgInfo.NickName,
			UserPhone:  userMsgInfo.UserPhone,
			UserEmail:  userMsgInfo.UserEmail,
			CreateTime: userMsgInfo.CreatedAt,
			UpdateTime: userMsgInfo.CreatedAt,
			Roles:      roleNames,
		},
		"code": "0000",
		"msg":  userMsgInfo.UserName,
	})
}

type UserListQueryParams struct {
	Current    int    `form:"current" binding:"required"`
	Size       int    `form:"size" binding:"required"`
	NickName   string `form:"nickName"`
	UserName   string `form:"userName"`
	UserGender string `form:"userGender"`
	UserPhone  string `form:"userPhone"`
	UserEmail  string `form:"userEmail"`
	UserStatus int64  `form:"userStatus"`
}

type UserDto struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreateBy   string    `gorm:"column:create_by;comment:创建者名称" json:"createBy"`                     // 创建者名称
	UpdateBy   string    `gorm:"column:update_by;comment:修改者名称" json:"updateBy"`                     // 修改者名称
	Status     int64     `gorm:"column:status;default:1;comment:状态（1-正常，2-禁用）" json:"status"`        // 状态（1-正常，2-禁用）
	UserName   string    `gorm:"column:user_name;not null;comment:用户名" json:"userName"`              // 用户名
	UserGender string    `gorm:"column:user_gender;default:1;comment:性别（1-男，2-女）" json:"userGender"` // 性别（1-男，2-女）
	NickName   string    `gorm:"column:nick_name;comment:昵称" json:"nickName"`                        // 昵称
	UserPhone  string    `gorm:"column:user_phone;comment:电话" json:"userPhone"`                      // 电话
	UserEmail  string    `gorm:"column:user_email;comment:邮箱" json:"userEmail"`                      // 邮箱
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	Roles      []string  `json:"userRoles"`
}
