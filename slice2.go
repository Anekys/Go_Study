package main

import "fmt"

func main()  {
	//声明一个切片slice1,并且赋初值为[1,2,3]
	//slice1:=[]int{1,2,3}

	//声明一个切片slice1,不赋初值,此时仅仅是声明,变量本身的内存为空,长度也为0
	var slice1 []int
	//slice1=make([]int,3)//make关键字类似于C语言的malloc,第一个参数为开辟内存的类型,第二个参数为开辟内存的数量,开辟出来的空间默认值为0
	//slice1[0]=1 //开辟好内存后就可以为切片的指定成员赋值了

	//声明一个切片slice1,为其开辟三份内存空间,全部为默认值0
	//var slice1 []int=make([]int,3)

	//与上一个方法等价
	//slice1 :=make([]int,3)
	fmt.Printf("len=%d,slice=%v",len(slice1),slice1) //%v表示打印变量的全部信息

	//判断一个slice是否为空
	if slice1==nil {		//只要slice1有元素,就算全部都是初值0,也不是空切片,只有一个元素都没有时才会成立
		fmt.Println("这是一个空切片")
	}else {		//注意这里的两个大括号一定要与else在同一行上,否则会报语法错误
		fmt.Println("这不是一个空切片")
	}
}
