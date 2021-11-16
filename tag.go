package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type People struct {
	Name string `info:"name" doc:"名字"`
	Sex string `info:"sex"`
}

func findtag(str interface{})  {
	//这里需要注意,使用reflect.Valueof(str).Elem()时,这个str必须是传入的地址值,要不然会直接Panic
	t:=reflect.TypeOf(str).Elem()
	for i:=0;i<t.NumField();i++{
		taginfo:=t.Field(i).Tag.Get("info")
		tagdoc:=t.Field(i).Tag.Get("doc")
		fmt.Printf("taginfo:%s , tagdoc:%s\n",taginfo,tagdoc)
	}
}
func _()  {
	//通过反射获取结构体标签
	var people People
	findtag(&people)
}

type Movie struct {
	//注意,如果有两个相同的标签,那么两个标签的内容都会被吞,无法转为JSON
	Tittle string	`json:"tittle"`
	Year int		`json:"year"`
	Price int		`json:"rmb"`
	Actors []string `json:"actors"`

}
func main() {
	movie:=Movie{
		Tittle:  "喜剧之王",
		Year:    2000,
		Price:   50,
		Actors: []string{"周星驰", "张柏芝"},
	}
	//编码的过程:结构体-->json
	jsonstr,err:=json.Marshal(movie)
	if err!=nil {
		fmt.Println("结构体转Json转换失败!")
		return
	}else {
		fmt.Printf("%s\n",jsonstr)
	}
	//解码的过程: json-->结构体
	mymovie:=Movie{}//创建一个空的用来接受json的对象
	err=json.Unmarshal(jsonstr,&mymovie)
	if err!=nil{
		fmt.Println("json转结构体失败!")
		return
	}
	fmt.Printf("%v\n",mymovie)

}