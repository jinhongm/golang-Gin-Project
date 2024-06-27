package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/uploads/utils/pwd"
	"gvb_server/utils/jwts"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"Please enter user name"`
	Password string `json:"password" binding:"required" msg:"Please enter passwords"`
}

func (UserApi) EmailLoginView(c *gin.Context) {
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		// 没找到
		global.Log.Warn("user name does not exist")
		res.FailWithMessage("User name or password is wrong", c)
		return
	}
	// 校验密码
	isCheck := pwd.CheckPwd(userModel.PassWord, cr.Password)
	if !isCheck {
		global.Log.Warn("password is wrong")
		res.FailWithMessage("User name or password is wrong", c)
		return
	}
	// 登录成功，生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("fail to create token", c)
		return
	}
	res.OKWithData(token, c)

}
