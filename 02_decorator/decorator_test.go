package decorator

func ExampleConcrete_Wear() {
	var p Person = &ConcreteComponent{}
	p = WearTShirt(p, "Blue")
	p = WearPants(p, 100)
	p = WearShoes(p, 42)
	p.Show()

	// Output: A Person wears sunglasses; Color: Blue, TShirt; Lenght: 100cm, Pants.; Size: 42, Shoes;
}

func ExampleMan_Show() {
	var xiaoming Person = &Man{}
	xiaoming = WearShoes(xiaoming, 43)
	xiaoming = WearTShirt(xiaoming, "White")

	xiaoming.Show()

	// Output:A man wear a hat!Size: 43, Shoes; Color: White, TShirt;
}

func ExampleWoman_Show() {
	var xiaohong Person = &Woman{}
	xiaohong = WearTShirt(xiaohong, "Red")
	xiaohong = WearShoes(xiaohong, 38)

	xiaohong.Show()

	// Output: A woman wear a skirt!Color: Red, TShirt; Size: 38, Shoes;
}
