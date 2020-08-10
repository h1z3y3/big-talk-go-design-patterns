package main

import (
	"errors"
	"fmt"
)

type IProduct interface {
	Description() string
}

type Car struct {
	Color  string
	Number string
}

func (c *Car) Description() string {
	return fmt.Sprintf("Car: color: %s, number: %s", c.Color, c.Number)
}

// Builder interface
type Builder interface {
	SetColor(string) Builder
	SetNumber(string) Builder
	Build() (IProduct, error)
}

// Director
type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) Generate() (IProduct, error) {
	return d.builder.SetColor("red").SetNumber("12345").Build()
}

// Concrete Builder
type CarBuilder struct {
	car Car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (c *CarBuilder) SetColor(a string) Builder {
	c.car.Color = a

	return c
}

func (c *CarBuilder) SetNumber(b string) Builder {
	c.car.Number = b
	return c
}

func (c *CarBuilder) Build() (IProduct, error) {

	if c.car.Color == "" {
		return nil, errors.New("color can not be empty")
	}

	if c.car.Number == "" {
		return nil, errors.New("number can not be empty")
	}

	return &c.car, nil
}
