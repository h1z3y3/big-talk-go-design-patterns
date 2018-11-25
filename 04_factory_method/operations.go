package factory_method

// 运算
type Operation interface {
	SetA(float64)
	SetB(float64)
	GetResult() (float64, error)
}

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

type AddOperation struct {
	*OperationBase
}

func (oper *AddOperation) GetResult() (float64, error) {
	return oper.a + oper.b, nil
}

type SubOperation struct {
	*OperationBase
}

func (oper *SubOperation) GetResult() (float64, error) {
	return oper.a - oper.b, nil
}

type MulOperation struct {
	*OperationBase
}

func (oper *MulOperation) GetResult() (float64, error) {
	return oper.a * oper.b, nil
}


