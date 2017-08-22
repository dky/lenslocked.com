package main

import (
	"fmt"
)

func hash(s string)int {
	return len(s) % 5
}

func main() {
	//var hello string
	hello := "somecrazyjunk"
	fmt.Println(len(hello))
	//returns the remainder from division
	fmt.Println(14 % 5)
	fmt.Println(hash(hello))
}
