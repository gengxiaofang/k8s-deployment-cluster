// 99乘法表并打印到文件
package main

import (
	"fmt"
	"os"
	log2 "log"
)

func main() {
	f,err := os.Create("/tmp/multiply_list.txt")
	if err != nil {
		log2.Fatal(err)
	}
	for i := 1;i < 10; i ++ {
		for n := 1; n <= i; n ++ {
			fmt.Printf("%d*%d=%d ",n,i,i*n)
			fmt.Fprintf(f,"%d*%d=%d ",n,i,i*n)
		}
		fmt.Println()
		fmt.Fprintln(f)
	}
	f.Close()
}
