package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func Math(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Println(r.URL.Path)
		fmt.Fprint(w, "10的平方：", math.Pow(10, 2))
	} else {
		fmt.Println(r.URL.Path)
		fmt.Fprint(w, "hello,世界！")
	}

}

func main() {
	http.HandleFunc("/", Math)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
