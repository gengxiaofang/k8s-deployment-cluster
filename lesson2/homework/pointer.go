// "&"取变量地址,"*"取一个指针所指向的对象,也就是如果一个指针保存着一个内存地址，那么它就返回在那个地址的对象;
package main

import (
	"fmt"
)


func main() {

	var a = 10	//赋值变量 a 为 10;
	var p = &a	//获取 a 的内存地址并赋值给变量p;

	fmt.Println("Before assignment, *p: ", *p)	// 打印结果为 10;
	fmt.Println("Before assignment, p: ", p)	// 打印结果为 a 的内存地址;
	fmt.Println("Before assignment, &a: ", &a)	// 打印结果为 a 的内存地址;
	fmt.Println()


	*p = 20
	fmt.Println("After assignment, *p: ", *p)	// 打印结果为 20,将 a 内存地址的对应的对象更改为 20,因此打印结果为20;
	fmt.Println("After assignment, p: ", p)		// 打印结果为 a 的内存地址;
	fmt.Println("After assignment, &a: ", &a)	// 打印结果为 a 的内存地址;
	fmt.Println("After assignment, a: ", a)		// 打印结果为 20;
	fmt.Println()


	a = 30
	fmt.Println("After assignment, *p: ", *p)	// 打印结果为 30,内存地址依旧不变,变的是内存地址对应的对象发生改变;
	fmt.Println("After assignment, p: ", p)		// 打印结果为 a 的内存地址;
	fmt.Println("After assignment, &a: ", &a)	// 打印结果为 a 的内存地址;
	fmt.Println("After assignment, a: ", a)		// 打印结果为 30;

}
