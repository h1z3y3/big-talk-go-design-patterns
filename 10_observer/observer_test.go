package observer

func ExampleObserver() {
	subject := NewSubject()

	aObserver := NewConcreteObserver("A")
	bObserver := NewConcreteObserver("B")
	cObserver := NewConcreteObserver("C")

	subject.register(aObserver)
	subject.register(bObserver)
	subject.register(cObserver)

	subject.UpdateContext("update string")

	// Output:
	// A receive update string
	// B receive update string
	// C receive update string
}
