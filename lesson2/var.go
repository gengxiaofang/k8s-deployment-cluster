package main

import (
	"fmt"
)

func main() {
	var i = 10
	fmt.Println(i)

	var y string = "hello,world"
	fmt.Println(y)

	x := "测试"
	fmt.Println(x)

	var (
		v1 int
		v2 string
		v3 float64
		v4 int
	)
	v1 = 1
	v2 = "v2"
	v3 = 2.0
	fmt.Println(v1, v2, v4)
	fmt.Printf("%v",v3)
}
