package middleware

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/redis_ser"
	"gvb_server/utils/jwts"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("do not have token", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token error", c)
			c.Abort()
			return
		}
		// 判断是否在redis中
		if redis_ser.CheckLogout(token) {
			res.FailWithMessage("token failure", c)
			c.Abort()
			return
		}

		// 登录的用户
		c.Set("claims", claims)
	}
}

func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if "logout_"+token == "" {
			res.FailWithMessage("do not have token", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token error", c)
			c.Abort()
			return
		}
		// 判断是否在redis中
		if redis_ser.CheckLogout(token) {
			res.FailWithMessage("token failure", c)
			c.Abort()
			return
		}
		// 登录的用户
		if claims.Role != int(ctype.PermissionAdmin) {
			res.FailWithMessage("admin error", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
