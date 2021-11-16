package main

import "fmt"

func _()  {	//匿名函数,类似于匿名导入,可以防止编译时编译器报错
	slice1:=make([]int,3,5) //3为长度(len),5为容量(cap) 如果未指定容量,make则默认容量与长度相等
	fmt.Printf("slice1长度:%d,容量:%d,  %v\n",len(slice1),cap(slice1),slice1)
	//切片的长度表示的是切片的合法成员数,而容量则表示实际开辟的内存空间
	slice1=append(slice1,1,2)	//append可以为切片追加新的元素,可以追加多个也可以追加单个,此处追加1和2
	//此时打印出来的slice长度已经变为5了,但是容量也还是5
	fmt.Printf("slice1长度:%d,容量:%d,  %v\n",len(slice1),cap(slice1),slice1)
	//注意,现在slice1的长度已经满了
	slice1=append(slice1,3)
	//此时slice1的长度再次＋1变为6,但是容量却变为了10.在切片容量已满的情况下再次append,编译器会自动在后面开辟出大小与容量相同的内存区域
	//因为slice1的容量为5,所以再次append时会再追加5份内存空间,并把这五份的第一个设置值为6
	fmt.Printf("slice1长度:%d,容量:%d,  %v\n",len(slice1),cap(slice1),slice1)
}
func main() {
	slice2:=[]int{1,2,3,4,5}
	fmt.Println(slice2)
	slice3:=slice2[:2]	//切片支持截取操作,与Python截取方法一致,左闭右开
	fmt.Println(slice3)
	slice3[0]=100		//需要注意的是,go截取后的数组如果值发生改变,被截取的数组的值也会发生改变,因为两个数组指向的是同一内存空间
	fmt.Printf("slice2长度:%d,容量:%d,  %v\n",len(slice2),cap(slice2),slice2)//截取产生的切片与原切片容量一致
	fmt.Printf("slice3长度:%d,容量:%d,  %v\n",len(slice3),cap(slice3),slice3)
	//如果不想修改截取的切片后原切片被改变,则可以使用copy函数进行复制
	slice4:=make([]int,6)
	copy(slice4,slice2) //copy会将原切片中的数据全部复制到新的切片中,以新切片的长度和容量为准依次填入数据
	//若原切片比复制的切片长,则多出部分会被丢弃,反之则新切片多出来的部分不变,依然为默认值0
	fmt.Printf("slice3长度:%d,容量:%d,  %v\n",len(slice4),cap(slice4),slice4)
}
