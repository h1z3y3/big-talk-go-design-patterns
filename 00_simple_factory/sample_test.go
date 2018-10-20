package simplefactory

import (
	"testing"
)

// test Add
func TestOperationAdd_GetResult(t *testing.T) {
	oper := NewOperation("+")
	var a, b float64
	a = 1
	b = 3
	r, _ := oper.GetResult(a, b)
	if r != a+b {
		t.Fatal("operationAdd test error")
	}
}

// test Sub
func TestOperationSub_GetResult(t *testing.T) {
	oper := NewOperation("-")
	var a, b float64
	a = 1
	b = 3
	r, _ := oper.GetResult(a, b)
	if r != a-b {
		t.Fatal("operationSub test error")
	}
}

// test Mul
func TestOperationMul_GetResult(t *testing.T) {
	oper := NewOperation("*")
	var a, b float64
	a = 1
	b = 3
	r, _ := oper.GetResult(a, b)
	if r != a*b {
		t.Fatal("operationMul test error")
	}
}

// test Div
func TestOperationDiv_GetResult(t *testing.T) {
	oper := NewOperation("/")
	var a, b float64
	a = 1
	b = 3
	r, _ := oper.GetResult(a, b)
	if r != a/b {
		t.Fatal("operationDiv test error")
	}

	a = 1
	b = 0
	r, err := oper.GetResult(a, b)
	if err == nil {
		t.Fatal("operationDiv test error")
	}
}