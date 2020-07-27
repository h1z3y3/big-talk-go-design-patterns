package main

import (
	"fmt"
	"testing"
)

func TestNewAdd01(t *testing.T) {

	builder := NewCarBuilder()

	car, err := builder.SetColor("red").SetNumber("123455").Build()

	if err != nil {
		t.Fatal("error:", err)
	}

	fmt.Println("car: ", car.Description())

}
