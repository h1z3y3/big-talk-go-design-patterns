package main

import (
	"fmt"
	"testing"
)

func TestNewAdd01(t *testing.T) {

	director := NewDirector(&CarBuilder{})
	adCar, err := director.Generate()

	if err != nil {
		t.Fatal("error:", err)
	}

	fmt.Println(adCar.Description())

}
