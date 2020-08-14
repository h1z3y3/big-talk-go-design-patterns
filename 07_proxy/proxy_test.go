package main

import (
	"testing"
)

func TestProxy(t *testing.T) {
	var subject ISubject
	subject = &Proxy{}

	result := subject.Do()

	if result != "pre:real:after" {
		t.Fail()
	}
}
