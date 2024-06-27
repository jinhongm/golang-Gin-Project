package config

import (
	"fmt"
)

type Es struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (es Es) URL() string {
	// 打印字段值以确认它们是否被正确设置
	return fmt.Sprintf("%s:%d", es.Host, es.Port)
}
