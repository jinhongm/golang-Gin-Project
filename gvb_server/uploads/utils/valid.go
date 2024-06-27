package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

func GetValidMsg(err error, obj any) string {
	getObj := reflect.TypeOf(obj)
	if errs, ok := err.(validator.ValidationErrors); ok { // 修改变量名为errs
		for _, e := range errs { // 这里遍历的是errs
			if f, exists := getObj.Elem().FieldByName(e.Field()); exists { // 修改拼写错误exits为exists
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}
