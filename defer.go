package main

import "fmt"

//func func1()  {
//	fmt.Println("A")
//}
//func func2()  {
//	fmt.Println("B")
//}
//func func3()  {
//	fmt.Println("C")
//}
func testFunc() int {
	defer deferfunc()//
	return returnfunc()//
}
func deferfunc() int {
	fmt.Println("defer called")
	return 0
}
func returnfunc() int {
	fmt.Println("return caller")
	return 0
}
func main()  {
	//defer fmt.Println("程序即将结束")	//defer关键字类似于C++的析构函数,在程序即将结束时会被自动调用
	//fmt.Println("Hellow")
	//fmt.Println("World!")
	//defer func1()
	//defer func2()
	//defer func3()	//同一个程序可以有多个defer,defer的调用顺序与堆栈一致,先进后出
	testFunc()		//当defer和return同时出现时,return先执行,defer后执行
}
