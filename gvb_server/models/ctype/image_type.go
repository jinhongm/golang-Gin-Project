package ctype

import "encoding/json"

type ImageType int

const (
	Local ImageType = 1
	QiNiu ImageType = 2
)

func (s ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s ImageType) String() any {
	var str string
	switch s {
	case Local:
		str = "Local"
	case QiNiu:
		str = "QiNiuCloud"
	default:
		str = "others"
	}
	return str
}
