package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

// 获取自定义msg错误
func GetValidMsg(err error, obj interface{}) string {
	getObj := reflect.TypeOf(obj)
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
				return f.Tag.Get("json") + f.Tag.Get("msg")
			}
		}
	}
	return err.Error()
}
