package decorator

func ExampleConcrete_Wear() {
	var p Person = &ConcreteComponent{}
	p = WearTShirt(p, "Blue")
	p = WearPants(p, 100)
	p = WearShoes(p, 42)
	p.Show()

	// Output: A Person wears sunglasses; Color: Blue, TShirt; Lenght: 100cm, Pants.; Size: 42, Shoes;
}

