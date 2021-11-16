package main

import "fmt"
//map属于hashlib本身是无序的
func main() {
	//第一种创建方式
	var map1 map[string]string	//声明一个空的map
	if map1==nil{
		fmt.Println("这是一个空map")
	}
	map1=make(map[string]string,3) //与切片类似,使用make开辟空间
	map1["1"]="one"	//为map赋值
	map1["2"]="two"
	map1["3"]="three"
	fmt.Println(map1)

	//第二种创建方式
	map2:=make(map[string]string,3)
	map2["1"]="one"
	map2["2"]="two"
	map2["3"]="three"
	fmt.Println(map2)

	//第三种创建方式
	map3:= map[int]string{
		1:"1",
		2:"2",
		3:"3", //注意,最后一个键值对后面也需要加逗号
	}
	fmt.Println(map3)
}
