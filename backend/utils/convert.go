// author:lybbn
// program: golyadmin
package utils

import (
	"encoding/json"
	"reflect"
)

// 通过json来转换struct结构体到map json
func ConvertStruct2MapJson(obj any) map[string]interface{} {
	data, _ := json.Marshal(obj)
	m := make(map[string]interface{})
	json.Unmarshal(data, &m)
	return m
}

// 通过反射reflect来转换struct结构体到map json(支持一层嵌套):性能比ConvertStruct2MapJson好
func ConvertStruct2MapJsonReflect(obj any) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		jsonKey := t.Field(i).Tag.Get("json")
		if jsonKey != "-" {
			if name == "GL_BASE_MODEL" || name == "GL_CONTROL_MODEL" {
				structField := v.Field(i).Type()
				for j := 0; j < structField.NumField(); j++ {
					data[structField.Field(j).Tag.Get("json")] = v.Field(i).Field(j).Interface()
				}
				continue
			}
			data[jsonKey] = v.Field(i).Interface()
		}
	}
	return data
}
