package main

import "fmt"

type human struct {
	name string
	sex string
}
func (this *human) Walk() {
	fmt.Printf("%s Walk...\n",this.name)
}
func (this *human) Eat() {
	fmt.Printf("%s Eat...\n",this.name)
}

type superman struct {
	human		//继承父类
	level int
}

func (this *superman) Eat() {	//重写父类方法
	fmt.Printf("%s Eat!\n",this.name)
}
func (this *superman) Fly()  {		//子类新方法
	fmt.Printf("%s Fly!\n",this.name)
}
func main() {
	Human:=human{
		name: "张三",
		sex:  "男",
	}
	Human.Walk()
	Human.Eat()
	// //定义子类的方法1
	//Superman:=superman{
	//	human: human{"李四","男"},
	//	level: 100,
	//}

	//方法2
	var Superman superman //定义一个类,此时类的所有属性都为空
	Superman.name="李四" //通过调用属性的方法修改定义的类的值
	Superman.Walk() //调用父类方法
	Superman.Eat()//调用重写的方法
	Superman.Fly()//调用子类新方法
}