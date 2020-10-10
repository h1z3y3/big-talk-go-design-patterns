package facade

import (
	"fmt"
)

type API interface {
	Test() string
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func NewAPIFacede() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

func (a *apiImpl) Test() string {
	retA := a.a.TestA()
	retB := a.b.TestB()

	return fmt.Sprintf("facade: a->%s b->%s", retA, retB)
}
