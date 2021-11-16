package main

import "fmt"

//interface{}是万能数据类型
func Func(arg interface{})  {
	fmt.Println("Func is called ....")
	fmt.Println(arg)
	//Golang为interface{}提供"类型断言"机制判断参数类型
	value,ok:=arg.(string)
	if ok{
		fmt.Println("arg is string , value=",value)
		fmt.Printf("%T\n",value)
	} else {
		fmt.Println("arg is not string")
	}
}

type Book struct {
	tittle string
}

func main() {
	book:=Book{"Golang"}
	Func(book)
	Func("text")
}