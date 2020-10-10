package facade

type AModuleAPI interface {
	TestA() string
}

type aModuleImpl struct{}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

func (a *aModuleImpl) TestA() string {
	return "A module"
}
