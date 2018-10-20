package simplefactory

import "fmt"

type Operation interface {
	GetResult(a float64, b float64) (float64, error)
}

func NewOperation(oper string) Operation {
	switch oper {
	case "+":
		return &operationAdd{}
	case "-":
		return &operationSub{}
	case "*":
		return &operationMul{}
	case "/":
		return &operationDiv{}
	default:
		return nil
	}
}

// 加法
type operationAdd struct{}

func (o *operationAdd) GetResult(a float64, b float64) (float64, error) {
	return a + b, nil
}

// 减法
type operationSub struct{}

func (o *operationSub) GetResult(a float64, b float64) (float64, error) {
	return a - b, nil
}

// 乘法
type operationMul struct{}

func (o *operationMul) GetResult(a float64, b float64) (float64, error) {
	return a * b, nil
}

// 除法
type operationDiv struct{}

func (o *operationDiv) GetResult(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为0")
	}

	return a / b, nil
}
