package main

import "fmt"

func main()  {
	//const length int =100
	//const length =100			//与上面的方法效果一致
	//fmt.Println(length)
	//const可以定义枚举类型
	//const (			//iota关键字会自动迭代数值,每个加一,第一个数值为0
	//	A=iota	//0			注意iota只能出现在const()的括号中,其他地方都是不可用的
	//	B		//1
	//	C		//2
	//)
	//fmt.Println(A,B,C)
	const (					//iota可以用作公式,且在公式未改变时,一直迭代下去,若公式改变则会按照新的公式迭代,但是iota的值不会变
		A,B=iota+2,iota*2		//A=0+2=2 B=0*2=0
		C,D						//C=1+2=3 D=1*2=2
		E,F=iota+3,iota*3		//此处公式发生了改变 E=2+3=5 F=2*3=6
		G,H						//G=3+3=6 H=3*3=9
	)
	fmt.Println(A,B)
	fmt.Println(C,D)
	fmt.Println(E,F)
	fmt.Println(G,H)
}
