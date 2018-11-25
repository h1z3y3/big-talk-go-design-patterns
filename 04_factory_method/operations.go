package factory_method

// 运算
type Operation interface {
	SetA(float64)
	SetB(float64)
	GetResult() (float64, error)
}

// 运算基类，实现公共的方法
type OperationBase struct {
	a float64
	b float64
}

func (oper *OperationBase) SetA(a float64) {
	oper.a = a
}

func (oper *OperationBase) SetB(b float64) {
	oper.b = b
}

// 加法运算
type AddOperation struct {
	*OperationBase
}

func (oper *AddOperation) GetResult() (float64, error) {
	return oper.a + oper.b, nil
}

// 减法运算
type SubOperation struct {
	*OperationBase
}

func (oper *SubOperation) GetResult() (float64, error) {
	return oper.a - oper.b, nil
}

// 乘法元算
type MulOperation struct {
	*OperationBase
}

func (oper *MulOperation) GetResult() (float64, error) {
	return oper.a * oper.b, nil
}


