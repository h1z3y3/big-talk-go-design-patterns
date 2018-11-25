package decorator

import "fmt"

type Person interface {
	Show()
}

// Person interface 的具体实现
type PersonConcreteComponent struct {
}

func (c *PersonConcreteComponent) Show() {
	fmt.Print("A Person wears sunglasses; ")
}

// 男人
type Man struct{}

func (m *Man) Show() {
	fmt.Print("A man wear a hat!")
}

// 女人
type Woman struct{}

func (w *Woman) Show() {
	fmt.Print("A woman wear a skirt!")
}

// TShirt
type TShirtDecorator struct {
	Person
	Color string
}

func (t *TShirtDecorator) Show() {
	t.Person.Show() // 调用父类的 Show() 方法
	// "装饰": 增加自己特有的属性
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
