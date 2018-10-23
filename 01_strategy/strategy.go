package strategy

import "fmt"

// Context 类
type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{
		strategy: strategy,
	}
}

func (c *Context) GetResult(a float64, b float64) (float64, error) {
	return c.GetResult(a, b)
}

// 策略接口
type Strategy interface {
	GetResult(a float64, b float64) (float64, error)
}

// 以下类实现策略接口
// 加法
type Add struct{}

func (o *Add) GetResult(a float64, b float64) (float64, error) {
	return a + b, nil
}

// 减法
type Sub struct{}

func (o *Sub) GetResult(a float64, b float64) (float64, error) {
	return a - b, nil
}

// 乘法
type Mul struct{}

func (o *Mul) GetResult(a float64, b float64) (float64, error) {
	return a * b, nil
}

// 除法
type Div struct{}

func (o *Div) GetResult(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为0")
	}

	return a / b, nil
}
