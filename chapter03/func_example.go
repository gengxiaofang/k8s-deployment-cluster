// 51reoobt golang 课程
package main

import (
	"sort"
	"fmt"
)


type Student struct {
	Name string
	Id   int
}

func main() {

	s := []int{2,6,7,3,4,9}

	sort.Slice(s, func(i, j int) bool { // 通过匿名函数实现对属组 s 进行排序;
		return s[i] < s[j]

	})


	ss := []Student{}
	ss = append(ss,Student{
		Name: "a",
		Id: 2,
	})

	ss = append(ss,Student{
		Name: "b",
		Id: 3,
	})

	ss = append(ss,Student{
		Name: "c",
		Id: 7,

	})

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Name < ss[j].Name
	})

	fmt.Println(ss)
}