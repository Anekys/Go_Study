package main

import "fmt"
//指针用法大致与C无异
func swap(a *int,b*int)  {
	//temp:=*a
	//*a=*b
	//*b=temp
	*a,*b=*b,*a	//Python同款语法也支持
}
func main()  {
	var a=10
	var b=20
	swap(&a,&b)
	fmt.Println("a=",a,"b=",b)
}
