package jwts

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"gvb_server/global"
	"time"
)

// 用户认证：
//用户登录后，服务器使用GenToken生成一个包含用户信息的JWT并返回给用户。
//用户在随后的请求中将JWT作为认证凭证（通常在HTTP头的Authorization字段）发送给服务器。
//服务端验证：
//服务器接收到请求后，使用ParseToken提取和验证JWT。
//如果验证成功，服务器从JWT中读取用户信息，并根据该信息处理请求。
//如果验证失败（例如，令牌过期或签名不正确），服务器拒绝请求。

// JwtPayLoad jwt中payload数据
type JwtPayLoad struct {
	Username string `json:"username"`  // 用户名
	NickName string `json:"nick_name"` // 昵称
	Role     int    `json:"role"`      // 权限  1 管理员  2 普通用户  3 游客
	UserID   uint   `json:"user_id"`   // 用户id
}

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

// GenToken 创建 Token
func GenToken(user JwtPayLoad) (string, error) {
	// 这个操作将这个字符串转换成了一个字节切片（[]byte 类型）。
	MySecret := []byte(global.Config.Jwt.Secret)
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))), // 默认2小时过期
			Issuer:    global.Config.Jwt.Issuer,                                                     // 签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

// ParseToken 解析 token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	MySecret := []byte(global.Config.Jwt.Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		global.Log.Error(fmt.Sprintf("token parse err: %s", err.Error()))
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
