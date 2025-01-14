package config

type QiNiu struct {
	Enable bool `json:"enable" yaml:"enable"`

	AccessKey string `json:"access_key" yaml:"access_key"`

	SecretKey string `json:"secret_key" yaml:"secret_key"`

	Bucket string `json:"bucket" yaml:"bucket"`

	CDN string `json:"cdn" yaml:"cdn"` //访问图片的地址的前缀

	Zone string `json:"zone" yaml:"zone"`

	Prefix string `json:"prefix" yaml:"prefix"`

	Size float64 `json:"size" yaml:"size"`
}
