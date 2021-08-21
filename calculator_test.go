package main

import (
	"golang-examples/calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	c := calculator.NewCalculator()
	r := c.Add(1, 1)
	if r != 2 {
		t.Log("error should be 2", r)
		t.Fail()
	}

}

func TestSub(t *testing.T) {
	c := calculator.NewCalculator()
	r := c.Sub(2, 1)
	if r != 1 {
		t.Log("error should be 1", r)
		t.Fail()
	}

}
