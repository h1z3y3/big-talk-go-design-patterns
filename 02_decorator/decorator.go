package decorator

import "fmt"

type Person interface {
	Show()
}

type ConcreteComponent struct {
}

func (c *ConcreteComponent) Show() {
	fmt.Print("A Person wears sunglasses; ")
}

// TShirt
type TShirtDecorator struct {
	Person
	Color string
}

func (t *TShirtDecorator) Show() {
	t.Person.Show()
	fmt.Print(fmt.Sprintf("Color: %s, TShirt; ", t.Color))
}

func WearTShirt(p Person, c string) Person {
	return &TShirtDecorator{
		Person: p,
		Color:  c,
	}
}

// Pants
type PantsDecorator struct {
	Person
	Length int64
}

func (p *PantsDecorator) Show() {
	p.Person.Show()
	fmt.Print(fmt.Sprintf("Lenght: %dcm, Pants.; ", p.Length))
}

func WearPants(p Person, l int64) Person {
	return &PantsDecorator{
		Person: p,
		Length: l,
	}
}

// Shoes
type ShoesDecorator struct {
	Person
	Size int64
}

func (s *ShoesDecorator) Show() {
	s.Person.Show()
	fmt.Print(fmt.Sprintf("Size: %d, Shoes; ", s.Size))
}

func WearShoes(p Person, s int64) Person {
	return &ShoesDecorator{
		Person: p,
		Size:   s,
	}
}