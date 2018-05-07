package main

import (
	"os"
	"flag"
	"strconv"
	"log"
	"fmt"
)

func main() {

	if os.Args[1] == "-h" {
		var help string
		flag.StringVar(&help,"h","该代码实现最基本的计算器功能","./expr 1 + 2")
		flag.Parse()
	} else {
		// strconv 字符串转行常用包
		a, err1 := strconv.Atoi(os.Args[1])
		if err1 != nil {
			log.Fatal(err1)
		}
		b, err2 := strconv.Atoi(os.Args[3])
		if err2 != nil {
			log.Fatal(err2)
		}

		c := os.Args[2]

		switch c {
		case "+":
			fmt.Printf("%d+%d=%d\n", a, b, a+b)
		case "-":
			fmt.Printf("%d-%d=%d\n", a, b, a-b)
		case "*":
			// 提示: *会被shell转义掉，使用\*代替
			fmt.Printf("%d*%d=%d\n", a, b, a*b)
		case "/":
			if b == 0 {
				fmt.Println("The divisor can't be 0...")
			} else {
				fmt.Printf("%d/%d=%f\n", a, b, float32(a)/float32(b))
			}
		default:
			fmt.Println("请执行-h查看帮助使用...")
		}
	}

}
