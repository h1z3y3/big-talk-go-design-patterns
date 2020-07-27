package main

import (
	"errors"
)

type Calculate interface {
	GetResult() int
}

type Builder interface {
	SetA(int) Builder
	SetB(int) Builder
	GetA() int
	GetB() int
	Build() (Calculate, error)
}

type Builder01 struct {
	A int
	B int
}

func (b1 *Builder01) SetA(a int) Builder {
	b1.A = a
	return b1
}

func (b1 *Builder01) SetB(b int) Builder {
	b1.B = b
	return b1
}

func (b1 *Builder01) GetA() int {
	return b1.A
}

func (b1 *Builder01) GetB() int {
	return b1.B
}

func (b1 *Builder01) Build() (Calculate, error) {
	if b1.A == 0 {
		return nil, errors.New("A can't be 0")
	}

	if b1.B == 0 {
		return nil, errors.New("B can't be 0")
	}

	return newAdd(b1), nil
}

type Add struct {
	builder Builder
}

func newAdd(builder Builder) *Add {
	return &Add{
		builder: builder,
	}
}

func (a *Add) GetResult() int {
	return a.builder.GetA() + a.builder.GetB()
}
