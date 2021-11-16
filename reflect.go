package main

import (
	"fmt"
	"reflect"
)


func _() {
	var num float32 = 3.1415926
	reflectnum(num)
}
func reflectnum(arg interface{}) {
	fmt.Println("Type is", reflect.TypeOf(arg))
	fmt.Println("Value is", reflect.ValueOf(arg))
}

type User struct {
	ID   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("user is Called...")
	fmt.Printf("%v\n", this)
}
func function(input interface{}) {
	//获取input的type,为空则返回nil
	inputType := reflect.TypeOf(input)
	//fmt.Println(inputType)
	//获取input的value,为空则返回0
	inputValue := reflect.ValueOf(input)
	//fmt.Println(inputValue)

	//通过type获取里面的方法(这里需要注意实现方法时是指针还是值,具体可见reflect2.go)
	fmt.Println(inputType.NumMethod())
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}

	//通过type获取里面的字段
	//1.获取interface的reflect.type字段,通过type得到Numfield,进行遍历
	//2.得到每个field,数据类型
	//3.通过field有一个Interface()方法得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v=%v\n", field.Type, field.Name, value)
		/* 打印结果:
		int: ID=1
		string: Name=张三
		int: Age=18
		*/
	}
}
func main() {
	user := User{
		ID:   1,
		Name: "张三",
		Age:  18,
	}
	function(user)
	//user.Call()
}
