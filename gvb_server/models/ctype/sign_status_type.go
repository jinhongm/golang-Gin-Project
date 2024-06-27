package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ    = 1
	SignGitee = 2
	SignEmail = 3
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() any {
	var str string
	switch s {
	case SignQQ:
		str = "QQ"
	case SignGitee:
		str = "Gitee"
	case SignEmail:
		str = "Email"
	default:
		str = "others"
	}
	return str
}
