package composite

import (
	"fmt"
)

type HumanResource interface {
	AddSubNode(HumanResource)

	Print(prefix string)
}

type humanResource struct {
	name       string
	department HumanResource
}

func (h *humanResource) AddSubNode(_ HumanResource) {

}

func (h *humanResource) Print(_ string) {}

type Employee struct {
	humanResource
}

func NewEmployee(name string) HumanResource {
	return &Employee{
		humanResource{
			name: name,
		},
	}
}

func (e *Employee) Print(prefix string) {
	fmt.Printf("%s- employee: %s\n", prefix, e.name)
}

type Department struct {
	humanResource
	subNodes []HumanResource
}

func NewDepartment(name string) HumanResource {
	return &Department{
		humanResource: humanResource{
			name: name,
		},
		subNodes: make([]HumanResource, 0),
	}
}

func (d *Department) AddSubNode(n HumanResource) {
	d.subNodes = append(d.subNodes, n)
}

func (d *Department) Print(prefix string) {
	fmt.Printf("%s+ department: %s\n", prefix, d.name)
	prefix += "    "
	for _, v := range d.subNodes {
		v.Print(prefix)
	}

}
