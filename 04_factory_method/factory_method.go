package factory_method

// 工厂
type OperationFactory interface {
	CreateOperation() Operation
}

type AddFactory struct {
}

func (f *AddFactory) CreateOperation() Operation {
	return &AddOperation{
		OperationBase: &OperationBase{},
	}
}

type SubFactory struct {
}

func (f *SubFactory) CreateOperation() Operation {
	return &SubOperation{
		OperationBase: &OperationBase{},
	}
}

type MulFactory struct {
}

func (f *MulFactory) CreateOperation() Operation {
	return &MulOperation{
		OperationBase: &OperationBase{},
	}
}
