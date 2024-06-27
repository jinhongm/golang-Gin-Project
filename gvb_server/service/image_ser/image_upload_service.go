package image_ser

import (
	"fmt"
	"gvb_server/api/plugins/qiniu"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/uploads/utils"
	"io"
	"mime/multipart"
	"path"
	"strings"
)

var WhiteImageList = []string{
	"jpg",
	"png",
	"jpeg",
	"ico",
	"tiff",
	"gif",
	"svg",
	"webp",
}

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`
}

// 处理图片文件上传的方法
func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	fmt.Printf("QiNiu Enabled: %v\n", global.Config.QiNiu.Enable)
	fileName := file.Filename
	basePath := global.Config.Upload.Path
	filePath := path.Join(basePath, file.Filename)
	res.FileName = filePath

	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	// 判断白名单
	if !utils.InTheList(suffix, WhiteImageList) {
		res.Msg = "illegal file"
		return
	}

	size := float64(file.Size) / float64(1024*1024)
	// 判断文件大小
	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("The picture size exceeds the setting size, the current file's size is:%.2fMB, setting size is %dMB", global.Config.Upload.Size)
		return
	}
	// 读取文件内容
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
	}
	byteData, err := io.ReadAll(fileObj)
	imageHash := utils.Md5(byteData)
	// 去数据库中查这个照片是否存在
	var bannerModel models.BannerModel
	err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
	if err == nil {
		res.Msg = "The picture exists"
		res.FileName = bannerModel.Path
		return
	}
	fileType := ctype.Local
	res.Msg = "The picture is uploading successfully"
	res.IsSuccess = true
	if global.Config.QiNiu.Enable {
		fmt.Println("yesss")
		filePath, err = qiniu.UploadImage(byteData, fileName, "gvb")
		if err != nil {
			global.Log.Error(err)
			res.Msg = err.Error()
			return
		}
		res.FileName = filePath
		res.Msg = "上传七牛成功"
		fileType = ctype.QiNiu
	}
	// 图片入库
	global.DB.Create(&models.BannerModel{
		Path:      filePath,
		Hash:      imageHash,
		Name:      fileName,
		ImageType: fileType,
	})
	return
}
