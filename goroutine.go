package main

import (
	"fmt"
	"runtime"
	"time"
)

func newtask()  {
	i:=0
	//下示for循环写法代表死循环
	for  {
		i++
		fmt.Printf("new goroutine : i=%d\n",i)
		time.Sleep(1*time.Second)
	}
}
func _()  {
	//使用go关键字创建一个goroutine来执行newtask()流程
	go newtask()
	i:=0
	for  {
		i++
		fmt.Printf("main goroutine:%d\n",i)
		time.Sleep(1*time.Second)
	}
}
func main() {
	//使用go创建承载一个形参为空返回值为空的一个函数
	go func(){
		defer fmt.Println("A.defer")
		func(){
			defer fmt.Println("B.defer")
			//return只能退出当前函数,如果想退出整个goroutine需要用以下指令
			runtime.Goexit()
			fmt.Println("B")
		}() //此处加上小括号代表函数立即执行,如果次匿名函数有参数,则要提供的参数写在这个括号中
		fmt.Println("A")
	}()
	//注意使用go承载的函数如果要获取返回值的话,是不能直接通过创建变量的方式来接受的,需要用到channel
	//为了防止goroutine随主程的销毁而销毁,此处加上for来阻塞主程
	for  {
		time.Sleep(1*time.Second)
	}
}
