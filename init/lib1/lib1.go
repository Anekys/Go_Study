package lib1

import "fmt"

//首字母大写的是公开函数,小写的是私有函数只能在包内访问
func Func1()  {		//注意,导出函数必须要首字母大写,否则无法调用
	fmt.Println("func1被调用")
}
func init()  {		//次函数为每个包的初始化函数,在import时会被自动调用
	fmt.Println("进入lib1")
}