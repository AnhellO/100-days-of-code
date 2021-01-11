package main

import (
	"fmt"
	"mytypes"
	"strings"
)

func main() {
	var sb strings.Builder
	sb.WriteString("Hello, ")
	sb.WriteString("Gophers!")
	fmt.Println(sb.Len())
	fmt.Println(sb.String())

	// Testing my custom builder
	var mb mytypes.MyBuilder
	fmt.Println(mb.Hello())
	// The following line won't work!
	// fmt.Println(mb.Len())
	mb.WriteString("Hello, Gophers!")
	fmt.Println(mb.Len())
}
