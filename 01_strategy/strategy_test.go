package strategy

import (
	"testing"
)

// test Add
func TestAdd_GetResult(t *testing.T) {
	ctx := NewContext(&Add{})
	var a, b float64
	a = 1
	b = 3
	r, _ := ctx.GetResult(a, b)
	if r != a-b {
		t.Fatal("operationSub test error")
	}
}

// test Sub
func TestSub_GetResult(t *testing.T) {
	ctx := NewContext(&Sub{})
	var a, b float64
	a = 1
	b = 3
	r, _ := ctx.GetResult(a, b)
	if r != a-b {
		t.Fatal("operationSub test error")
	}
}

// test Mul
func TestMul_GetResult(t *testing.T) {
	ctx := NewContext(&Mul{})
	var a, b float64
	a = 1
	b = 3
	r, _ := ctx.GetResult(a, b)
	if r != a*b {
		t.Fatal("operationMul test error")
	}
}

// test Div
func TestOperationDiv_GetResult(t *testing.T) {
	ctx := NewContext(&Div{})
	var a, b float64
	a = 1
	b = 3
	r, _ := ctx.GetResult(a, b)
	if r != a/b {
		t.Fatal("operationDiv test error")
	}

	a = 1
	b = 0
	r, err := ctx.GetResult(a, b)
	if err == nil {
		t.Fatal("operationDiv test error")
	}
}
