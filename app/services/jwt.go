package services

import (
	"Gin_Start/global"
	"Gin_Start/utils"
	"context"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type jwtService struct {
}

var JwtService = new(jwtService)

type JwtUser interface {
	GetUid() string
}

// CustomClaims 自定义声明
type CustomClaims struct {
	jwt.StandardClaims // jwt的标准字段
}

const (
	TokenType    = "bearer" // token类型
	AppGuardName = "app"    // app守卫名称
)

type TokenOutPut struct {
	AccessToken string `json:"access_token"` // token
	TokenType   string `json:"token_type"`   // token类型
	ExpiresIn   int    `json:"expires_in"`   // 过期时间
}

// CreateToken 创建token
func (jwtService *jwtService) CreateToken(GuardName string, user JwtUser) (tokenData TokenOutPut, err error, token *jwt.Token) {
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + global.App.Config.Jwt.JwtTtl, // 过期时间
			Id:        user.GetUid(),                                    // 用户id
			Issuer:    GuardName,                                        // 签发人 用于 标识这个JWT的签发主体 可以避免token被其他人使用
			NotBefore: time.Now().Unix() - 1000,                         // 生效时间 -1000 是为了防止服务器时间和客户端时间不一致
		},
	},
	)
	// 生成token
	tokenStr, err := token.SignedString([]byte(global.App.Config.Jwt.Secret)) // 生成token
	tokenData = TokenOutPut{
		tokenStr,
		TokenType,
		int(global.App.Config.Jwt.JwtTtl),
	}
	return
}

// 获取黑名单缓存 key
func (jwtService *jwtService) getBlackListKey(tokenStr string) string {
	return "jwt_black_list:" + utils.MD5([]byte(tokenStr))
}

// JoinBlackList token 加入黑名单
func (jwtService *jwtService) JoinBlackList(token *jwt.Token) (err error) {
	nowUnix := time.Now().Unix()
	timer := time.Duration(token.Claims.(*CustomClaims).ExpiresAt-nowUnix) * time.Second
	// 将 token 剩余时间设置为缓存有效期，并将当前时间作为缓存 value 值
	err = global.App.Redis.SetNX(context.Background(), jwtService.getBlackListKey(token.Raw), nowUnix, timer).Err()
	return
}

// IsInBlacklist token 是否在黑名单中
func (jwtService *jwtService) IsInBlacklist(tokenStr string) bool {
	joinUnixStr, err := global.App.Redis.Get(context.Background(), jwtService.getBlackListKey(tokenStr)).Result()
	joinUnix, err := strconv.ParseInt(joinUnixStr, 10, 64)
	if joinUnixStr == "" || err != nil {
		return false
	}
	// JwtBlacklistGracePeriod 为黑名单宽限时间，避免并发请求失效
	if time.Now().Unix()-joinUnix < global.App.Config.Jwt.JwtBlacklistGracePeriod {
		return false
	}
	return true
}
