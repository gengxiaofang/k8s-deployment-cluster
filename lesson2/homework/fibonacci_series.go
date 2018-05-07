package main

import (
	"fmt"
	"flag"
	"time"
)

// 递归函数
func fibA(n int) int {
	if n > 2 {
		return fibA(n-1) + fibA(n-2)
	} else {
		return n
	}

}

// for迭代
func fibB(n int) int {
	a,b := 0,1
	for i :=0; i <= n;i++  {
		a,b =b,a+b
	}
	return a
}



func main() {
	var num int
	flag.IntVar(&num, "n", 0, "the number of Fibonacci series")
	flag.Parse()

	t1 := time.Now()
	var f1 = fibA(num)
	fmt.Printf(" Fibonacci series of %d is:%d\n", num, f1)
	fmt.Println("执行时间：", time.Since(t1))

	t2 := time.Now()
	var f2 = fibB(num)
	fmt.Printf(" Fibonacci series of %d is:%d\n", num, f2)
	fmt.Println("执行时间：", time.Since(t2))

}
