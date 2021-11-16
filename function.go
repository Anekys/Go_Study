package main
//函数的用法
import "fmt"
//声明函数格式如下
//func 函数名(参数名1 参数1类型 , 参数名2 参数2类型) (返回值1类型,返回值2类型){函数体}
//单个返回值时可不加()
func func1(a string,b int) int {
	fmt.Println(a)
	fmt.Println(b)
	c:=100
	return c
}
func func2(a int,b int) (int,int) {
	return b,a
}
//返回多个参数值,有形参名称的   这里定义的参数名顺序决定了返回值的顺序
func func3(a int , b int) (r1 int ,r2 int) {	//这里的返回值若类型一致也可以写成(r1,r2 int)
	//这里的r1和r2的作用域为整个func3,默认值是0
	r2=a
	r1=b
	//如果未给r1,r2进行赋值,也可以在return后面写上要返回的值
	return	r1,r2	//如果已经给r1和r2赋值,这里可以直接return,如果未赋值也可以在这里直接填值,会自动对应r1和r2的顺序进行返回
}
func main()  {
	var a=func1("hello",666)
	fmt.Println(a)
	b,c:= func2(1,2)
	fmt.Println(b,c)
	b,c= func2(1,2)
	fmt.Println(b,c)
}
