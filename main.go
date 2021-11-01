package main

import (
	"fmt"
)

func Sum(v1 int, v2 int) (v int) {
	return v1 + v2
}

func main() {

	var v1, v2 int

	fmt.Printf("v1: ")
	fmt.Scan(&v1)
	fmt.Print("v2: ")
	fmt.Scan(&v2)
	fmt.Println("result:", Sum(v1, v2))
}
