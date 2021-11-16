package main

import (
	"fmt"
	"io"
)
import "os"

func _()  {
	//每个变量都有一个pair,有两个组成部分,一个type,一个value,value为变量的值,type分为static type(系统自带类型,int,string)
	//																	   concrete type(interface指向的具体的数据类型)
	//pair <statictype:string,value:"Golang">
	var a string
	a="Golang"
	var Alltype interface{}
	Alltype=a	//interface的type指向了string,而value指向了"Golang"
	fmt.Println(Alltype)
}
func _()  {
	//tty:pair<type:*os.File,value:"test"文件描述符>
	tty,err:=os.OpenFile("test",os.O_RDWR,0)
	if err!=nil{
		fmt.Println("Open file error:",err)
		return
	}
	//r:pair<type: , value:>
	var r io.Reader
	//r:pair<type:*os.File , value:"test"文件描述符>
	r=tty
	//w:pair<type: , value:>
	var w io.Writer
	//w:pair<type:*os.File , value:"test"文件描述符>
	w=r.(io.Writer)
	//向test文件中写入"Hellow world"
	w.Write([]byte("Hellow world"))
	//此实验可以证明在变量与变量互相赋值时,会保持pair的一致
}

type Writer interface {
	WriteBook()
}
type Reader interface {
	ReadBook()
}
//具体类型
type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read Book...")
}
func (this *Book) WriteBook() {
	fmt.Println("Write Book...")
}
func main() {
	//b:pair<type:Book,value:book{}地址>
	b:=&Book{}
	//r:pair<type: , value:>
	var r Reader
	//r:pair<type:Book, value:book{}地址>
	r=b
	r.ReadBook()
	var w Writer
	//w:pair<type:Book, value:book{}地址>
	w=r.(Writer) //此处的断言为什么会成功?因为w和r具体的type一致的
	w.WriteBook()
}