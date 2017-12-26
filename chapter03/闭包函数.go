package main

import "fmt"

func main() {
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}

func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
/*
func main() {
	var fs = [4]func() {}

	for i := 0; i < 4; i ++ {
		defer fmt.Println("defer i =",i)
		defer func() { fmt.Println("defer_closure i = ", i) }()
		fs[i] = func() { fmt.Println("closure i = ",i) } //闭包函数
	}
	for _, f := range fs {
		f()
	}
}
 */
