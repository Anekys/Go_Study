package main

import (
	"fmt"
	"time"
)

func _()  {		//无缓冲channel
	//定义一个channel,默认无缓存
	c:=make(chan int)
	go func() {
		defer fmt.Println("Goroutine结束")
		fmt.Println("goroutine 正在运行...")
		c<-666//表示将666传入已经定义的c(channel)中
	}()
	num:=<-c //从channel(c)中取出一个数据并复制给num
	//注意此处的<-c实质上是一个阻塞操作,如果channel中没有数据可以读出,则此时会阻塞主程,直到获取到数据后再向下执行
	fmt.Println(num)
}

func _()  {
	c:=make(chan int,3) //创建一个带缓冲的channel,缓存容量为3
	go func() {
		defer fmt.Println("Goroutine结束")
		fmt.Println("goroutine 正在运行...")
		for i:=0;i<3;i++ {
			c<-i//将i的内容写入channel中,如果channel中已满,则此操作会阻塞Goroutine
		}
	}()
	time.Sleep(2*time.Second)//暂停两秒钟保证Channel中已经写满3个元素
	for i:=0;i<3;i++{
		fmt.Println(<-c) //从channel中取出一个元素并打印
	}
}

func _()  {		//channel的关闭与range
	c:=make(chan int)
	//如果是nil channel(通过var新建一个不初始化的channel)则收发数据都会阻塞
	go func() {
		for i:=0;i<5;i++{
			c<-i
		}
		//close()可以关闭一个channel
		close(c)//此处如果不关闭channel,则主程中的for循环会一直从channel中获取数据,造成死锁,导致编译器报错
		//关闭channel后,依然可以从channel中读取数据,但是不能发送数据,否则会引发panic报错并立即返回0
	}()
	//接收数据写法1
	/*for{
		if data,ok:=<-c;ok{ //通过判断ok的布尔值来得知Channel是否关闭
			fmt.Println(data)
		}else {
			break
		}
	}
	fmt.Println("main goroutine over")*/
	//写法2
	for data:=range c{
		fmt.Println(data)
	}
}
func fibonacii(c,q chan int)  {	//斐波那契数列
	x,y:=1,1
	for {
		select {	//select具备多路监听功能,可以同时监听多个channel
		case c<-x: //如果c可写,则会执行下列语句
			x=y
			y=x+y
		case <-q: //如果q可读出数据
			fmt.Println("quit")
			return //退出当前函数
		/*default: //只有当上面所有的case都无法执行时,才会执行default的语句
			fmt.Println("~~~~~")*/
		}
	}
}
func main()  {
	c:=make(chan int)
	q:=make(chan int)
	go func() {
		for i:=0;i<10;i++{
			fmt.Println(<-c)
		}
		q<-0
	}()
	fibonacii(c,q)
}
