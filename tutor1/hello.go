package main

import (
	"fmt"
	"github.com/1xe1/go-network/reverse"
)

func main() {
	var t bool = true
	var f bool
	r := reverse.Reverse("Hello, world!")

	fmt.Println(t)
	fmt.Println(f)
	fmt.Println(r)
}