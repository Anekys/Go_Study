package main

import "fmt"

func main()  {
	//方法1:声明一个变量默认值为0
	var a int
	fmt.Println(a)				//println会在打印后方自动换行
	fmt.Printf("a type:%T\n",a)		//printf为格式化输出,%T为类型
	//方法2:声明一个变量并赋初值
	var b int =100
	fmt.Println("b=",b)
	//方法3:通过赋初值的形式,由编译器自动匹配变量类型
	var c =200
	fmt.Println("c=",c)
	var cc="Test String"
	fmt.Println("cc Value =",cc)
	fmt.Printf("cc type is %T\n",cc)
	//方法4:var都省略直接由编译器进行自动匹配 ,此方法只能定义局部变量,定义全局变量需用方法1,2,3
	dd:="编译器自动匹配变量类型"
	fmt.Println("dd=",dd)
	fmt.Printf("dd type is %T\n",dd)
	e:=3.1415926
	fmt.Println("e=",e)
	fmt.Printf("e type is %T\n",e)
	//多个变量声明
	var aa,bb int =100,200
	fmt.Println("aa=",aa,"bb=",bb)
	//声明多个不同类型的变量
	//写法1:
	var CC,DD = 50,"Hello"
	fmt.Println("cc=",CC,"dd=",DD)
	//写法2:
	var(
		ee="o ha you"
		ff=0.1
		gg=true
	)
	fmt.Println("ee=",ee,"\nff=",ff,"\ngg=",gg)
}