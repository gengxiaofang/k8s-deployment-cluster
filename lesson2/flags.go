package main

import (
	"fmt"
	"flag"
)

func main() {
	var file = flag.String("f","","go run cat.go filepath")
	flag.Parse()
	fmt.Println("the file name:",*file)
}
