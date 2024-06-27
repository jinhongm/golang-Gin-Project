package qiniu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"gvb_server/config"
	"gvb_server/global"
	"time"
)

func getToken(q config.QiNiu) string {
	accessKey := q.AccessKey
	secretKey := q.SecretKey
	bucket := q.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

func getCfg(q config.QiNiu) storage.Config {
	cfg := storage.Config{}
	zone, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
	cfg.Zone = &zone
	cfg.UseHTTPS = false
	return cfg
}

func UploadImage(data []byte, imageName string, prefix string) (filePath string, err error) {
	// 确保启用了七牛云存储
	if !global.Config.QiNiu.Enable {
		return "", errors.New("Please enable QiNiu to upload ")
	}

	q := global.Config.QiNiu
	// 验证必要的配置项
	if q.AccessKey == "" || q.SecretKey == "" {
		return "", errors.New("请配置accessKey and secretKey")
	}

	// 检查文件大小是否超过限制
	if float64(len(data))/1024/1024 > q.Size {
		return "", errors.New("文件超过指定大小")
	}

	upToken := getToken(q)
	cfg := getCfg(q)

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	dataLen := int64(len(data))
	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf("%s/%s__%s", prefix, now, imageName) // 修正字符串格式化的使用

	// 假设 putExtra 已经定义
	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &storage.PutExtra{})
	if err != nil { // 修正错误处理语法
		return "", err
	}

	// 使用 q.Cdn 和 ret.Key 构建最终的文件路径
	return fmt.Sprintf("%s/%s", q.CDN, ret.Key), nil // 假设 q.Cdn 存储了CDN域名
}
