package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"please fulfill the menu title" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"please fulfill the menu path" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime  *int        `json:"abstract_time"`                                                             // 切换的时间，单位秒
	BannerTime    int         `json:"banner_time" structs:"banner_time"`                                         // 切换的时间，单位秒
	Sort          int         `json:"sort" binding:"required" msg:"please enter the menu number" structs:"sort"` // 菜单的序号
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`                                               // 具体图片的顺序
}

func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var menuList []models.MenuModel
	// 这行代码通过全局的数据库实例global.DB执行一个查询，寻找数据库中所有title字段值等于cr.Title并且path字段值等于cr.Path的MenuModel记录。
	// 这里使用的是GORM的Find方法，它将查询结果存储到menuList变量中。
	// "title = ? and path = ?"是SQL查询的条件部分，cr.Title和cr.Path是这些条件的参数。
	// 查询执行后，.RowsAffected属性用于获取查询影响的行数，即数据库中满足条件的记录数量。这个值被赋给了count变量。
	count := global.DB.Find(&menuList, "title = ? and path = ?", cr.Title, cr.Path).RowsAffected
	if count > 0 {
		res.FailWithMessage("repeated menu", c)
		return
	}

	//创建banner 数据入库
	menuModel := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	err = global.DB.Create(&menuModel).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("Fail to add the menu", c)
	}

	if len(cr.ImageSortList) == 0 {
		res.OKWithMessage("Successfully to add the menu", c)
		return
	}

	//批量入库
	var menuBannerList []models.MenuBannerModel
	for _, sort := range cr.ImageSortList {
		// 判断图片id是否存在
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		res.FailWithMessage("Failed to associate menu image", c)
		return
	}
	res.OKWithMessage("Successfully add it to the menu", c)
}
