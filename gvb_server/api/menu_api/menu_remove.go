package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (MenuApi) MenuRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var menuList []models.MenuModel
	// 这行代码尝试从数据库中查询所有在cr.IDList中指定的ID的MenuModel记录。
	// cr.IDList应该是一个包含菜单ID的切片。GORM的Find方法会将所有找到的记录填充到menuList中，而.RowsAffected属性则返回查询影响的行数，即实际找到的记录数量，这个值被赋给了count变量。
	count := global.DB.Find(&menuList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("The menu does not exist", c)
		return
	}

	// 启动事务：
	//global.DB.Transaction(func(tx *gorm.DB) error {...})这行代码开始了一个数据库事务。tx *gorm.DB是事务的句柄，
	// 应该用于事务内的所有数据库操作。这确保了操作的原子性：如果事务内的任何操作失败，整个事务都会回滚。
	//
	//清除关联关系：
	//err = global.DB.Model(&menuList).Association("Banners").Clear()这行代码试图清除menuList中所有菜单项与横幅的关联关系。
	//但这里有一个问题：它使用的是global.DB而不是事务句柄tx，这意味着这个操作并没有在事务中执行。
	//
	//删除菜单项：
	//err = global.DB.Delete(&menuList).Error这行代码尝试删除menuList中的所有菜单项。
	//同样，这里应该使用tx而不是global.DB来确保操作在事务中执行。
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&menuList).Association("Banners").Clear(); err != nil {
			global.Log.Error(err)
			return err
		}
		if err := tx.Delete(&menuList).Error; err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("Fail to delete the menu", c)
		return
	}
	res.OKWithMessage(fmt.Sprintf("totally delete %d menus", count), c)

}
