package factory_method

import "testing"

func TestAddFactory_CreateOperation(t *testing.T) {
	var factory OperationFactory
	factory = &AddFactory{}

	add := factory.CreateOperation()

	var a, b float64
	a, b = 1, 2
	add.SetA(a)
	add.SetB(b)

	res, err := add.GetResult()

	if err != nil || res != a+b {
		t.Fail()
	}
}

func TestSubFactory_CreateOperation(t *testing.T) {
	var factory OperationFactory
	factory = &SubFactory{}

	sub := factory.CreateOperation()

	var a, b float64
	sub.SetA(a)
	sub.SetB(b)

	res, err := sub.GetResult()

	if err != nil || res != a-b {
		t.Fail()
	}
}
