package res

type ErrorCode int

const (
	SettingsError ErrorCode = 1001 // 系统错误
	ArgumentError ErrorCode = 1002
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "System Error",
		ArgumentError: "Argument Error",
	}
)
