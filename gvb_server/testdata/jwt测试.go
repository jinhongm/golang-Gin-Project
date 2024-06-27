package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/utils/jwts"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	fmt.Println(global.Config.Jwt)
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   1,
		Role:     1,
		Username: "ggboy",
		NickName: "sb",
	})
	fmt.Println(token, err)
	claims, err := jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImdnYm95Iiwibmlja19uYW1lIjo\nic2IiLCJyb2xlIjoxLCJ1c2VyX2lkIjoxLCJleHAiOjE3MDk4MDk5MzcuOTg1NzkzLCJpc3MiOiIxMjM\n0In0.I36Iqys_6lc7BxXbpqcb-JkFB1CnzuU1qAETE7yxJEQ")
	fmt.Println(claims, err)
}
