package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func main() {
	//用+号来链接字符串
	s1 := "hello" + "world!"
	fmt.Println(s1)

	//随机读取
	c := s1[0]
	fmt.Println(string(c))

	//切片（start,不包含end）
	s2 := s1[0:3]
	fmt.Println(string(s2))

	//取长度
	s3 := s1[0:len(s1)]
	fmt.Println(s3)

	//反射
	i,_ := strconv.Atoi("-42")
	fmt.Println(i)
	fmt.Println("type of i:",reflect.TypeOf(i))

	//浮点数转换
	f,_ :=strconv.ParseFloat("3.14",32)
	fmt.Println("type of f:",reflect.TypeOf(f),f)

	//字符串修改
	s4 := []byte(s1)
	s4[0] = 'x'
	fmt.Println(string(s4))

}