package ctype

import (
	"database/sql/driver"
	"strings"
)

type Array []string

// *t用于解引用指针t，或声明一个指向类型T的指针变量。
// *Array作为方法接收器，表示方法可以修改Array类型实例的值。
// value在Scan方法中代表从数据库中检索的原始值，需要被转换为接收器类型的值。
func (t *Array) Scan(value interface{}) error {
	v, _ := value.([]byte)
	if string(v) == "" {
		*t = []string{}
		return nil
	}

	*t = strings.Split(string(v), "\n")
	return nil
}

func (t Array) Value() (driver.Value, error) {
	return strings.Join(t, "\n"), nil
}
