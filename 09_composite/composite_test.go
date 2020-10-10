package composite

func ExampleHumanResource() {
	root := NewDepartment("root")

	dpt1 := NewDepartment("01")
	dpt2 := NewDepartment("02")
	dpt3 := NewDepartment("03")

	employee1 := NewEmployee("001")
	employee2 := NewEmployee("002")
	employee3 := NewEmployee("003")
	employee4 := NewEmployee("004")
	employee5 := NewEmployee("005")
	employee6 := NewEmployee("006")
	employee7 := NewEmployee("007")
	employee8 := NewEmployee("008")

	root.AddSubNode(employee1)
	root.AddSubNode(dpt1)
	root.AddSubNode(dpt2)

	dpt1.AddSubNode(employee2)
	dpt1.AddSubNode(employee3)
	dpt1.AddSubNode(dpt3)

	dpt2.AddSubNode(employee4)
	dpt2.AddSubNode(employee5)
	dpt2.AddSubNode(employee6)

	dpt3.AddSubNode(employee7)
	dpt3.AddSubNode(employee8)

	root.Print("")

	// Output:
	// + department: root
	//     - employee: 001
	//     + department: 01
	//         - employee: 002
	//         - employee: 003
	//         + department: 03
	//             - employee: 007
	//             - employee: 008
	//     + department: 02
	//         - employee: 004
	//         - employee: 005
	//         - employee: 006

}
