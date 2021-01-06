package main

import (
	"fmt"
	"mytypes"
)

func main() {
	x := mytypes.MyInt(9)
	fmt.Println(x.Twice())
}
