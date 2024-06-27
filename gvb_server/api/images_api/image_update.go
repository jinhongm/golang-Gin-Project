package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"please choose file id"`
	Name string `json:"name" binding:"required" msg:"please enter file's name'"`
}

// c.ShouldBindJSON(&cr)尝试将客户端发来的JSON格式的请求体解析到cr变量中。
// 这意味着，如果客户端发送了一个包含图像更新信息的请求，那么这些信息将会被存储在cr实例中。
func (ImagesApi) ImageUpdateView(c *gin.Context) {
	var cr ImageUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var imageModel models.BannerModel
	err = global.DB.Take(&imageModel, cr.ID).Error
	if err != nil {
		res.FailWithMessage("The file does not exist", c)
		return
	}
	err = global.DB.Model(&imageModel).Update("name", cr.Name).Error
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OKWithMessage("Success", c)
	return
}
