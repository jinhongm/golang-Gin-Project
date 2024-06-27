package advert_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// AdvertUpdateView 更新广告
// @Tags 广告管理
// @Summary 更新广告
// @Param token header string  true  "token"
// @Description 更新广告
// @Param data body AdvertRequest    true  "广告的一些参数"
// @Param id path int true "id"
// @Router /api/adverts/{id} [put]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		res.FailWithMessage("advertisement does not exist", c)
		return
	}

	maps := structs.Map(&cr)
	err = global.DB.Model(&advert).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("Fail to update the advertisement", c)
		return
	}
	res.OKWithMessage("Successful to update the advertisement", c)
}
