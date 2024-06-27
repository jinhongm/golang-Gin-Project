package advert_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"strings"
)

// AdvertListView 广告列表
// @Tags 广告管理
// @Summary 广告列表
// @Description 广告列表
// @Param data query models.PageInfo    false  "查询参数"
// @Router /api/adverts [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.AdvertModel]}
func (AdvertApi) AdvertListView(c *gin.Context) {
	var cr models.PageInfo

	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}
	referer := c.GetHeader("Referer")
	isShow := true
	if strings.Contains(referer, "admin") {
		isShow = false
	}
	// 判断 Referer是否包含admin 如果是就全部返回 不是就返回is_show=true
	list, count, _ := common.ComList(models.AdvertModel{IsShow: isShow}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	res.OKWithList(list, count, c)
	return
}
