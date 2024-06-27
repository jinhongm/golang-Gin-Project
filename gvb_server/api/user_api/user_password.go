package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/uploads/utils/pwd"
	"gvb_server/utils/jwts"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd"` // 旧密码
	Pwd    string `json:"pwd"`     // 新密码
}

// UserUpdatePassword
func (UserApi) UserUpdatePassword(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr UpdatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("The user does not exist", c)
		return
	}

	if !pwd.CheckPwd(user.PassWord, cr.OldPwd) {
		res.FailWithMessage("Password is wrong", c)
		return
	}
	hashPwd := pwd.HashPwd(cr.Pwd)
	err = global.DB.Model(&user).Update("pass_word", hashPwd).Error
	if err != nil {
		global.Log.Error(err)

		res.FailWithMessage("Fail to update the password", c)
		return
	}
	res.OKWithMessage("Update the password successfully", c)
	return
}
