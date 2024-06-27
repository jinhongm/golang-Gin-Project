package flag

import (
	sys_flag "flag"
	"github.com/fatih/structs"
)

type Option struct {
	DB   bool
	User string // -u admin -u user
	Es   string // -es create -es delete
}

func Parse() Option {
	db := sys_flag.Bool("db", false, "initializing database")
	user := sys_flag.String("u", "", "create the accounts")
	es := sys_flag.String("es", "", "create the elastic search")
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
		Es:   *es,
	}
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) (f bool) {
	maps := structs.Map(&option)
	for _, v := range maps {
		// 在 switch 语句内部，根据 v 的类型分别处理。
		// 如果 v 的类型是 string 或 bool，则执行相应的代码块。这允许函数根据不同的字段类型进行条件判断。
		switch val := v.(type) {
		case string:
			if val != "" {
				f = true
			}
		case bool:
			if val == true {
				f = true
			}
		}
	}
	return f
}

func SwitchOption(option Option) {
	if option.DB { // 同上，修改Option为option
		Makemigrations() // 确保这个函数在其他地方定义
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
	}
	if option.Es == "create" {
		EsCreateIndex()
	}
}
