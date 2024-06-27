package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb_server/config"
	"gvb_server/global"
	"io/fs"
	"io/ioutil"
	"log"
)

const ConfigFile = "settings.yaml"

// InitConf 读取yaml文件的配置
func InitConf() {
	// 这里创建了一个指向 config.Config 结构体的新实例的指针。
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		// 使用 panic 抛出一个包装了原始错误信息的新错误
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("Failed to load config yamlFile: %s", err) // 更正为适当的错误处理
	}
	log.Println("Config yamlFile load Init success.") // 使用正确的日志函数

	global.Config = c // global 是一个包含全局变量的包，Config 是该包中定义的一个变量，
	//是公共可访问的，以便程序的其他部分可以读取配置数据。
}

func SetYaml() error {
	// 将整个 global.Config 对象转换为 YAML 格式的字节流
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		global.Log.Error(err)
		return err
	}

	// 将 YAML 字节流写入配置文件
	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm) // 确保 ConfigFile 是全局配置文件的正确路径
	if err != nil {
		global.Log.Error(err)
		return err
	}

	global.Log.Info("The configuration file has been successfully modified.")
	return nil
}
