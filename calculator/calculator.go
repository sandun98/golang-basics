package calculator

import (
	"errors"
	"fmt"
)

// interface
type ICalculator interface {
	Add(a int, b int) int
	Sub(a int, b int) int
	Mul(a int, b int) int
	Div(a float64, b float64) (float64, error)
	Pow(base float64, power float64) float64
	privatefunction() // package visible ?

}

// implementation
type calculator struct {
}

func (c calculator) privatefunction() {
	fmt.Println("private function")
}

func (c calculator) Add(a int, b int) int {
	return a + b
}

func (c calculator) Sub(a int, b int) int {
	return a - b
}

func (c calculator) Mul(a int, b int) int {
	return a * b
}

func (c calculator) Div(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func (c calculator) Pow(base float64, power float64) float64 {
	var result float64 = 1

	for i := 0; float64(i) < power; i++ {
		result = result * base
	}
	return result
}

func NewCalculator() ICalculator {
	return &calculator{}
}
