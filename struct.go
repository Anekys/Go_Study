package main

import "fmt"

type myint int	//类似于C语言的typedf
type book struct {
	tittle string
	auth string
}

func changebook(book2 book)  {
	//值传递,实参的修改不会影响形参
	book2.tittle="Python"
}
func changeBook(book *book)  {
	//引用传递,成功修改形参的值
	book.tittle="Python"
}
func _()  {
	var a myint
	fmt.Printf("value=%d type=%T\n",a,a)
	var Book book
	Book.tittle="Golang"
	Book.auth="zhang3"
	changeBook(&Book)
	fmt.Printf("%v\n",Book)
}

//类名的首字母大写表示其他包也能够访问
type Hero struct {	//定义一个Hero结构体
	//属性名的首字母大写可以让其他包能够访问,小写只有类内部可以进行访问
	name string
	ad int
	level int
}
//方法名小写只有本包可以访问该方法
func (this Hero) show() {	//通过this类名将方法绑定到结构体上使其成为一个类
	//此处的this为值传递,并不会影响类本身的数据
	fmt.Println(this.name)
	fmt.Println(this.ad)
	fmt.Println(this.level)
}
//方法名大写其他包也可以访问该方法
func (this *Hero) Setname(name string) {
	//此处传递的是类的指针,所以可以修改形参的数据,不再只是类的副本
	this.name=name
}
func main() {
	//创建一个hero对象
	hero:=Hero{
		name:  "张三",
		ad:    100,
		level: 1,
	}
	hero.show()
	hero.Setname("李四")
	hero.show()
}
