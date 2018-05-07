package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"os"
)

func printFile(name string) {
	buf,err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(buf))
}

func main() {
	A := os.Args[1]
	printFile(A)
}
