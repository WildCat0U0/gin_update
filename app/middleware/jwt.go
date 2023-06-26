package middleware

import (
	"Gin_Start/app/common/response"
	"Gin_Start/app/services"
	"Gin_Start/global"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//tokenStr := c.Request.Header.Get("Authorization")
		tokenStr, _ := c.Cookie("username")
		if tokenStr == "" {
			response.TokenEmpty(c)
			c.Abort()
			return
		}
		//tokenStr = tokenStr[len(services.TokenType)+1:]

		// Token 解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &services.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Secret), nil
		})
		//if err != nil || services.JwtService.IsInBlacklist(tokenStr) || token.Valid != true {
		//	response.TokenFail(c)
		//	c.Abort()
		//	return
		//}
		if err != nil || services.JwtService.IsInBlacklist(tokenStr) {
			response.TokenFail(c)
			c.Abort()
			return
		}

		claims := token.Claims.(*services.CustomClaims)
		// Token 发布者校验
		if claims.Issuer != GuardName {
			response.TokenFail(c)
			c.Abort()
			return
		}

		//token 续签
		//这段代码的作用是在 JWT Token 即将过期时，自动刷新该 Token。
		//具体地，代码首先计算了当前 Token 距离过期还有多少秒，如果剩余时间小于配置文件中的 RefreshGracePeriod，就表示该 Token 即将过期，需要刷新。
		//接下来，代码通过获取一个名为 refresh_token_lock 的 Redis 锁来实现并发控制。如果成功获取到该锁，代码调用 JWT Service 的 GetUserInfo 方法获取当前用户的信息，并使用该信息创建一个新的 JWT Token，将新 Token 的 Access Token 和 ExpiresIn 分别放入响应头中的 new-token 和 new-expires-in 字段中。最后，代码调用 JWT Service 的 JoinBlackList 方法将原 Token 放入黑名单，使其无法再次使用。
		//需要注意的是，该代码中使用了 Redis 锁来进行并发控制。这样做的好处在于，可以避免多个请求同时刷新 Token，从而保证数据的一致性和安全性。
		//global.App.Log.Error(strconv.FormatInt(claims.ExpiresAt, 10))
		//global.App.Log.Error(strconv.FormatInt(time.Now().Unix(), 10))
		if claims.ExpiresAt-time.Now().Unix() < global.App.Config.Jwt.RefreshGracePeriod {
			lock := global.Lock("refresh_token_lock", global.App.Config.Jwt.JwtBlacklistGracePeriod)
			if lock.Get() {
				err, user := services.JwtService.GetUserInfo(GuardName, claims.Id)
				if err != nil {
					global.App.Log.Error(err.Error())
					lock.Release()
				} else {
					tokenData, _, _ := services.JwtService.CreateToken(GuardName, user)
					//c.Header("new-token", tokenData.AccessToken)
					//c.Header("new-expires-in", strconv.Itoa(tokenData.ExpiresIn))
					cookie1 := &http.Cookie{Name: "username", Value: tokenData.AccessToken, Expires: time.Unix(0, 0).Add(time.Duration(tokenData.ExpiresIn) * time.Second), HttpOnly: false}
					c.SetCookie(cookie1.Name, cookie1.Value, cookie1.MaxAge, cookie1.Path, cookie1.Domain, false, cookie1.HttpOnly)
					_ = services.JwtService.JoinBlackList(token)
				}
			}
		}

		c.Set("token", token)
		c.Set("id", claims.Id)
	}
}
