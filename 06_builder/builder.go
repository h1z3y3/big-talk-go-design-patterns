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
	return fmt.Sprintf("Car: %s, %s", c.Color, c.Number)
}

// Concrete Builder
type CarBuilder struct {
	car Car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (c *CarBuilder) SetColor(a string) *CarBuilder {
	c.car.Color = a

	return c
}

func (c *CarBuilder) SetNumber(b string) *CarBuilder {
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
