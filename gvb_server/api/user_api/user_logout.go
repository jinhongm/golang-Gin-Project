package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/utils/jwts"
)

func (UserApi) LogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	token := c.Request.Header.Get("token")

	err := service.ServiceApp.UserService.Logout(claims, token)

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("fail to cancel the account", c)
		return
	}

	res.OKWithMessage("successfully to cancel the account", c)

}
