package main

import (
	"fmt"
)

func main() {
	shift := 3
	s := []int{1, 2, 3, 4, 5}
	s = append(s[shift:], s[:shift]...)
}
