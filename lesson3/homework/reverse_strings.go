package main

import "fmt"

// rune 实现字符串反转
func reverse (s string) string {
	s1 := []rune(s)
	var s2 []rune
	for i := 0; i < len(s1);i ++ {
		s2 = append(s2,s1[len(s1)-i-1])
	}
	return string(s2)
}

func reverse1(s string) string {
	var s1 string
	for i :=0;i < len(s);i++ {
		s1 = s1 + fmt.Sprintf("%c",s[len(s)-i-1])
	}
	return s1
}

func main() {
	str := "hello,中国"
	fmt.Println(reverse(str))
}
