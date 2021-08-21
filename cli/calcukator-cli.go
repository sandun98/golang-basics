package main

import (
	"fmt"
	"golang-basics/calculator"
)

func main() {
	c := NewCalculator()

	fmt.Println(c.Add(10, 23))
}
