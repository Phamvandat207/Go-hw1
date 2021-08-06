package main

import (
	"fmt"
)

func main() {
	a, b := 0, 1

	for a < 100 {
		b = a + b
		a = b - a
		fmt.Println(a)
	}
}
