/*
类型转换：
1.不支持隐士类型转换;
2.不能将其他类型当bool值使用;
 */
package main

var x,y,z int
var s,n = "abc",123

var (
	a int
	b float32
)

func main() {
	n,s := 0x1234,"hello,world!"
	println(x,s,n) //0,"hello,world",4660

//	data,i := [3]int{0,1,2},0 data是一个slice,其中的值为0,1,2;i 值为0;

//	i,data[i] = 2,100 i为2,data为100;

//理解重新赋值与定义新同名变量的区别
	x := "abc"
	println(&x)

	x,y := "hello",20
	println(&x,y)
	{
		x,z :=1000,30
		println(&x,z)
	}
}

