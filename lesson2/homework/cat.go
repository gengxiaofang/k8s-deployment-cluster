package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func printFile(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(buf))
}

func main() {
	// filename := os.Args[1]
	var filename string
	flag.StringVar(&filename, "f", "./cat file_to_path", "读取某个文件的内容")
	flag.Parse()
	printFile(filename)
}
