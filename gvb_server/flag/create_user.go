package flag

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/uploads/utils/pwd"
)

func CreateUser(permissions string) {
	// 创建用户的逻辑
	// 用户名 昵称 密码 确认密码 邮箱
	var (
		userName   string
		nickName   string
		password   string
		rePassword string
		email      string
	)
	fmt.Printf("Please enter your user name：")
	fmt.Scan(&userName)
	fmt.Printf("Please enter your nick name：")
	fmt.Scan(&nickName)
	fmt.Printf("Please enter the email：")
	fmt.Scan(&email)
	fmt.Printf("please enter the password：")
	fmt.Scan(&password)
	fmt.Printf("please enter your password again：")
	fmt.Scan(&rePassword)

	// 判断用户名是否存在
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		// 存在
		global.Log.Error("The user name already existed, please enter again")
		return
	}
	// 校验两次密码
	if password != rePassword {
		global.Log.Error("These two passwords are not consistent, please enter again")
		return
	}
	// 对密码进行hash
	hashPwd := pwd.HashPwd(password)

	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	}

	// 头像问题
	// 1. 默认头像
	// 2. 随机选择头像
	avatar := "/uploads/avatar/5.png"

	// 入库
	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		PassWord:   hashPwd,
		Email:      email,
		Role:       ctype.Role(role),
		Avatar:     avatar,
		IP:         "127.0.0.1",
		Addr:       "intranet address",
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("User %s is created succefully!", userName)

}
