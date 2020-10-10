package facade

type BModuleAPI interface {
	TestB() string
}

type bModuleImpl struct{}

func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

func (b *bModuleImpl) TestB() string {
	return "B module"
}
