package routers

import "gvb_server/api"

func (router RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi
	router.GET("images", app.ImageListView)
	router.GET("image_names", app.ImageNameList)
	router.POST("images", app.ImagesUploadView)
	router.DELETE("images", app.ImageRemoveView)
	router.PUT("images", app.ImageUpdateView)
}
