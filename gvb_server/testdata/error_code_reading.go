package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

const file = "models/res/error_code.json" // 指定了错误码文件的路径。

// ErrMap是一个映射类型，键是res.ErrorCode（假设是int的别名），值是对应的错误信息字符串。
type ErrMap map[string]string

func main() {
	byteData, err := os.ReadFile(file) // 读取文件内容到byteData。
	if err != nil {
		logrus.Error(err) // 如果读取文件时发生错误，则记录错误并退出。
		return
	}

	var errMap ErrMap                       // 声明一个ErrMap类型的变量。
	err = json.Unmarshal(byteData, &errMap) // 解析JSON数据填充到errMap中。
	if err != nil {
		logrus.Error(err) // 如果解析JSON数据时发生错误，则记录错误并退出。
		return
	}

	fmt.Println(errMap) // 打印解析后的错误码映射。
}
