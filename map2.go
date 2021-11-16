package main

import (
	"fmt"
)

func printMap(Map map[string]string)  {
	//此处的形参为引用传递,在函数内部操作map会影响到实参
	for key,value:=range Map{
		fmt.Println(key,value)
	}
}

func main() {
	map1:= map[string]string{
		"China":"Beijing",
		"Japon":"Tokyo",
		"USA":"NewYouk",
	}
	//遍历
	printMap(map1)
	//删除
	delete(map1,"Japon")

	//修改
	map1["USA"]="San Francisco"

	fmt.Println("---------------")
	printMap(map1)
}