package main

import (
	"testing"
)

func TestNewAdd01(t *testing.T) {

	var builder Builder

	builder = &Builder01{}

	a, b := 1, 3

	add, err := builder.SetA(a).SetB(b).Build()

	if err != nil {
		t.Fatal("err is not nil", err)
	}

	result := add.GetResult()

	if result != a+b {
		t.Fatal("result is error", result)
	}
}

func TestNewAdd02(t *testing.T) {

	var builder Builder

	builder = &Builder01{}

	a, b := 0, 3

	_, err := builder.SetA(a).SetB(b).Build()

	if err == nil {
		t.Fatal("err should be nil", err)
	}

}
