package main

import (
	"fmt"
)

func main() {
	aSlice := []int{1, 2, 3, 4, 5}
  fmt.Println(rotate(aSlice))  
}
  
  func rotate(s[]int, shift int) {
    return append(s[shift:], s[:shift]...)
  }
