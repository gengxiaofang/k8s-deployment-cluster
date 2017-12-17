//常量必须是编译期确定的数字/字符串/布尔值,并且常量不占用内存空间;
//常量值还可以是 len/cap/unsafe.Sizeof 等编译期可确定结果的函数返回值;

package main

const x,y int  = 1,2         // 多常量初始化
const s  = "hello,Golang" //类型推算

const (                      // 常量组
	a,b = 10,100
	c bool = false
	d                         //如果不提供类型和初始值,那么视作与上一常量相同;
)

// 枚举
const (
	Sunday = iota     // 0
	Monday            // 1，通常省略后续⾏行表达式。
	Tuesday           // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

const (
	_ = iota                    // iota = 0
	KB int64 = 1 << (10 * iota) // iota = 1
	MB                          // 与 KB 表达式相同，但 iota = 2
	GB
	TB
)

func main()  {                // 未使用局部常量不会引发编译错误;
	const x  = "xxx"
	println(x,y,a,b,c,d)      // println{xxx 2 10 100 false false}

	// 修改字符串,字符串修改需先转换成[]rune或[]byte,修改完之后再转换成string;
	s := "abcd"
	bs := []byte(s)

	bs[1] = 'B'
	println(string(bs))

	us := []rune(s)
	us[1] = '您'
	println(string(us))
}
