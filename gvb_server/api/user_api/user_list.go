package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"gvb_server/utils/desens"
	"gvb_server/utils/jwts"
)

func (UserApi) UserListView(c *gin.Context) {
	// 如何判断是管理员
	_claims, _ := c.Get("claims")
	claims := _claims.(jwts.CustomClaims)
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var users []models.UserModel
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
	})
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			// 管理员
			user.UserName = ""
		}
		user.Tel = desens.DesensitizationTel(user.Tel)
		user.Email = desens.DesensitizationEmail(user.Email)
		// 脱敏
		users = append(users, user)
	}

	res.OKWithList(users, count, c)
}
