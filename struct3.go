package main

import (
	"fmt"
)
//本质是一个指针
type Animal interface {	//定义一个名为Animal的接口(interface) Go通过继承接口实现多态
	Sleep()
	GetColor() string
	GetType() string
}
//定义一个具体的类
type Cat struct {
	//类继承接口需要类重写接口的全部方法,否则不视为继承接口
	color string
}
func (this *Cat) Sleep() {
	fmt.Println("cat sleep")
}
func(this *Cat) GetColor() string{
	return this.color
}
func(this *Cat) GetType() string{
	return "Cat"
}

//定义一个狗的类
type Dog struct {
	//类继承接口需要类重写接口的全部方法,否则不视为继承接口
	color string
}
func (this *Dog) Sleep() {
	fmt.Println("Dog sleep")
}
func(this *Dog) GetColor() string{
	return this.color
}
func(this *Dog) GetType() string{
	return "Dog"
}
func showAnimal(animal Animal)  {	//参数为interface,传参的时候需要传指针
	animal.Sleep()
	fmt.Println(animal.GetType())
	fmt.Println(animal.GetColor())
}
func main() {
	/*
	//实现方法1
	var animal Animal
	animal=&Cat{"white"} //接口的本质是一个指针,所以需要对类用取地址符,此处animal已经继承了Cat的方法和属性
	fmt.Println("Color is",animal.GetColor())
	animal.Sleep() //调用cat的sleep方法
	animal=&Dog{color: "black"}
	fmt.Println("Color is",animal.GetColor())
	animal.Sleep() //调用Dog的sleep方法
	*/

	//方法2
	cat:=Cat{color:"white"}
	dog:=Dog{color: "black"}
	showAnimal(&cat) //传入cat则实现Cat的方法
	showAnimal(&dog) //传入dog则实现Dog的方法

	//多态的实现必须有一个接口作为父类,且必须有子类实现了接口的全部方法
	//最后父类的指针指向子类的具体对象,实现多态
}
