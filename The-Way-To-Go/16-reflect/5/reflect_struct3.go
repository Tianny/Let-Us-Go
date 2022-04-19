package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Person struct {
	Name        string `label:"Person Name: " uppercase:"true"`
	Age         int    `label:"Age is: "`
	Sex         string `label:"Sex is: "`
	Description string
}

func PrintUseTag(ptr interface{}) error {
	// 获取入参的类型
	reType := reflect.TypeOf(ptr)
	// 入参类型校验
	if reType.Kind() != reflect.Ptr || reType.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("参数应该为结构体指针")
	}
	// 取指针指向的结构体变量
	v := reflect.ValueOf(ptr).Elem()

	// 解析字段, NumField() 4 个字段。
	for i := 0; i < v.NumField(); i++ {
		// 获取结构体字段信息
		structField := v.Type().Field(i)
		// 取tag
		tag := structField.Tag
		// 解析label tag，获取tag值
		label := tag.Get("label")
		if label == "" {
			label = structField.Name + ": "
		}
		// 解析uppercase tag, 除 uppercase 标签外，全部小写输出。
		value := fmt.Sprintf("%v", v.Field(i))

		// 字段类型为字符串时， 判断tag值是否为真，进行大小写转换。
		if structField.Type.Kind() == reflect.String {
			tagVal := tag.Get("uppercase")
			if tagVal == "true" {
				value = strings.ToUpper(value)
			} else {
				value = strings.ToLower(value)
			}
		}
		fmt.Println(label + value)
	}
	return nil
}
func main() {
	person := Person{"Tom", 29, "Male", "Cool"}
	err := PrintUseTag(&person)
	if err != nil {
		return
	}
}
