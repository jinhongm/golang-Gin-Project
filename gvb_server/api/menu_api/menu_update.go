package menu_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	id := c.Param("id")
	var menuModel models.MenuModel
	err = global.DB.Preload("Banners").Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("The menu does not exist", c)
		return
	}
	global.DB.Model(&menuModel).Association("Banners").Clear()
	// 选择了banner那就更新
	if len(cr.ImageSortList) > 0 {
		// 操作第三方表
		var bannerList []models.MenuBannerModel
		for _, sort := range cr.ImageSortList {
			bannerList = append(bannerList, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}
		err = global.DB.Create(&bannerList).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("Fail to create Menu photos", c)
			return
		}
	}

	maps := structs.Map(&cr)
	err = global.DB.Model(&menuModel).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("Fail to update the menu", c)
		return
	}
	res.OKWithMessage("Successful to update the menu", c)

}
