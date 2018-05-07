package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"fmt"
)

func main() {
	dir_list,err := ioutil.ReadDir("/proc/")
	if err != nil {
		log.Fatal(err)
	}

	var b = true
	reg := regexp.MustCompile(`^[\d]{1,}`)
	for _,dir := range dir_list {
		b = reg.MatchString(dir.Name())
		if dir.IsDir() && true == b{
			fmt.Println(dir.Name())
		}
	}
}
