package main

import "fmt"
//定义数组方式1
func func1()  {
	//定义一个固定长度的数组
	var myaArray[10]int
	for i:=0;i<10;i++{
		fmt.Println(myaArray[i])
	}
}
//定义数组方式2
func func2()  {
	myArray2:=[10]int{1,2,3,4}	//定义一个固定数量的数组,并不完全初始化
	for index,value:=range myArray2{	//index为遍历myArray2的下标,value为遍历的值
		fmt.Println("index=",index,"value=",value)
	}
	//fmt.Printf("myArray2的数据类型为:%T",myArray2) //此处打印的值类型为[10]int,如果有四个成员则为[4]int
}
func printArray(Array[10]int)  {	//注意这里数组作为参数实质也是值传递,与C不同.成员数不同的数组,类型不同,本函数传递不是10成员的int数组会报错
	for index,value:=range Array{
		fmt.Println("index=",index,"value=",value)
	}
}
func func3()  {
	mySlice:=[]int{1,2,3,4,5,6,7,8,9,0}//定义一个切片(动态数组)slice
	//printArray(mySlice) //这里不能调用定义的printArray函数
	fmt.Printf("切片的数据类型为:%T\n", mySlice) //此处打印的数据类型为[]int,
	printSlice(mySlice)

}
func main()  {
	func1()
	func2()
	func3()

}
func printSlice(Slice []int)  {	//本函数可以传入任意长度的Slice(切片)而不用担心像数组一样因为成员数不同而报错
	for _,value:=range Slice{	//此处的_表示匿名变量,由于range返回两个元素,而又不需要index,所以可用匿名变量顶替
		fmt.Println("value=",value)
	}
	Slice[0]=100	//参数为Slice时是引用传递,修改对应元素的值时会使被传进来的变量的值也发生改变
	for _,value:=range Slice{
		fmt.Println("value=",value)
	}
}