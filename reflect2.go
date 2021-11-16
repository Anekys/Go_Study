package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

// Call 使用指针接收者实现了一个方法
func (this *User) CallPoint() {
	this.Id = 100
	fmt.Printf("%v\n", this)
}

// CallValue 使用值接收者实现了一个方法
func (this User) CallValue() {
	this.Id = 200
	fmt.Printf("%v\n", this)
}

func main() {
	fmt.Println("---1----")
	user := User{1, "ding", 18}
	// User 类型的值可以用来调用使用值接收者声明的方法
	GetField(user)
	fmt.Println("---2----")
	// User 类型值的指针也可以用来调用使用值接收者声明的方法
	GetField(&user)

	fmt.Println("===3====")
	user2 := User{1, "ding", 28}
	// 实参接收者是 T 类型的变量而形参接收者是 *T 类型
	// 编译器会隐式地获取变量的地址 user2.CallPoint() // 隐式转换为 (&user)
	user2.CallValue()

	fmt.Println("===4====")
	user3 := &User{1, "ding", 38}
	// 实参接收者是 *T 类型而形参接收者是 T 类型
	// 编译器会隐式地解引用接收者，获得实际的取值
	user3.CallValue() // 隐式转换为 (*user3)
	user3.CallPoint()
}

func GetField(arg interface{}) {
	argType := reflect.TypeOf(arg)
	fmt.Println(argType.NumMethod())
	for i := 0; i < argType.NumMethod(); i++ {
		m := argType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}