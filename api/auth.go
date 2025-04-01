package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"soybean-admin-go/config"
	"soybean-admin-go/db/gen"
	"soybean-admin-go/utils/log"
	"strings"
	"time"
)

// 定义JWT密钥
var jwtKey = []byte("your_secret_key")

// 定义Claims结构体
type Claims struct {
	UserID int64 `json:"userName"`
	jwt.RegisteredClaims
}

// 生成Token
func generateToken(userID int64) (string, string, error) {
	// 设置Token过期时间
	expirationTime := time.Now().Add(15 * time.Hour)
	// 创建Claims
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	// 创建Token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成Token字符串
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	// 设置Refresh Token过期时间
	refreshExpirationTime := time.Now().Add(24 * time.Hour)
	// 创建Refresh Token的Claims
	refreshClaims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpirationTime),
		},
	}
	// 创建Refresh Token对象
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	// 生成Refresh Token字符串
	refreshTokenString, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	return tokenString, refreshTokenString, nil
}

// 验证Token
func validateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	// 解析Token
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

// 登录处理函数
func Login(c *gin.Context) {
	// 模拟验证用户名和密码
	// 生成Token和Refresh Token
	var (
		user UserRequest
		u    = gen.Q.User
	)
	err := c.BindJSON(&user)
	if err != nil {
		config.Logger.Error("Invalid request", log.Field{
			Key:   "error",
			Value: err,
		})
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	userInfo, err := u.WithContext(c).Where(u.UserName.Eq(user.UserName)).First()
	if err != nil {
		config.Logger.Error("Failed to get user info", log.Field{
			Key:   "error",
			Value: err,
		})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}
	if user.Password != userInfo.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	fmt.Println(userInfo.ID)
	token, refreshToken, err := generateToken(userInfo.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	// 返回Token和Refresh Token
	c.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"token":         token,
			"refresh_token": refreshToken,
		},
		"code": "0000",
		"msg":  "请求成功",
	})
}

// 验证Token中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取Token
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
			c.Abort()
			return
		}
		stringSlice := strings.Split(tokenString, " ")
		// 验证Token
		claims, err := validateToken(stringSlice[1])
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		// 将用户ID存储到上下文
		c.Set("UserID", claims.UserID)
		c.Next()
	}
}

// 刷新Token处理函数
func RefreshToken(c *gin.Context) {
	// 从请求头中获取Refresh Token
	refreshTokenString := c.GetHeader("Refresh-Token")
	if refreshTokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing refresh token"})
		return
	}
	// 验证Refresh Token
	claims, err := validateToken(refreshTokenString)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}
	// 生成新的Token和Refresh Token
	token, newRefreshToken, err := generateToken(claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	// 返回新的Token和Refresh Token
	c.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"token":         token,
			"refresh_token": newRefreshToken,
		},
		"code": "0000",
		"msg":  "请求成功",
	})
}

func GetUserInfo(ctx *gin.Context) {
	s := ctx.GetInt64("UserID")
	q := gen.Q
	first, err := q.User.WithContext(ctx).Where(q.User.ID.Eq(s)).Select(q.User.UserName).First()
	if err != nil {
		config.Logger.Error("Failed to get user info", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}
	subQuery := q.UserRole.WithContext(ctx).Select(q.UserRole.RoleID).Where(q.UserRole.UserID.Eq(s))
	Roles, err := q.Role.WithContext(ctx).Where(q.Role.WithContext(ctx).Columns(q.Role.ID).Gt(subQuery)).Find()
	if err != nil {
		config.Logger.Error("Failed to get user info", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}
	var roleNames []string
	for _, m := range Roles {
		roleNames = append(roleNames, m.RoleCode)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"userId":   s,
			"userName": first.UserName,
			"roles":    roleNames,
			"buttons":  []string{},
		},
		"code": "0000",
		"msg":  "请求成功",
	})
}

type UserRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
