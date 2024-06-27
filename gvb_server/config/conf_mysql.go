package config

import (
	"fmt"
	"strconv"
)

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Config   string `yaml:"config"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"`
}

func (m *Mysql) Dsn() string {
	// 打印字段值以确认它们是否被正确设置
	fmt.Printf("Host: %s, Port: %d, User: %s, Password: %s, DB: %s\n", m.Host, m.Port, m.User, m.Password, m.DB)

	dsn := m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
	return dsn
}
