package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type ImageNameResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

// ImageNameList 图片列表
// @Tags 图片管理
// @Summary 图片名称列表
// @Description 图片名称列表
// @Router /api/image_names [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]ImageNameResponse}
func (ImagesApi) ImageNameList(c *gin.Context) {
	var imageNameList []ImageNameResponse
	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imageNameList)
	res.OKWithData(imageNameList, c)
}
