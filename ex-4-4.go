package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	s = append(s[shift:], s[:shift]...)
}
