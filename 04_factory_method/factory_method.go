package factory_method

// 工厂类
type OperationFactory interface {
	CreateOperation() Operation
}

// 加法工厂
type AddFactory struct {
}

func (f *AddFactory) CreateOperation() Operation {
	return &AddOperation{
		OperationBase: &OperationBase{},
	}
}

// 减法工厂
type SubFactory struct {
}

func (f *SubFactory) CreateOperation() Operation {
	return &SubOperation{
		OperationBase: &OperationBase{},
	}
}

// 乘法工厂
type MulFactory struct {
}

func (f *MulFactory) CreateOperation() Operation {
	return &MulOperation{
		OperationBase: &OperationBase{},
	}
}
