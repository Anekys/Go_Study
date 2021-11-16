package main
//由于是自己定义的包,不在GoRoot和GoPath中,所以必须指定路径
import (
	. "study/init/lib1" //在导入包的前面加上点即可将指定包的所有内容导入至本包内,此时语法从lib1.Func1()变为Func1()
	//_ "study/init/lib2" //在导入的包前面加上下划线表示为匿名导入,此时包的init方法会正常运行,但是其他公开函数无法调用
	//在要导的包前面加上其他字符可作为别名使用
	mylib "study/init/lib2"
)
func main()  {
	Func1()
	mylib.Func2()
}
