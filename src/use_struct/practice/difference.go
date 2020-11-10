/** 演示结构体方法和结构体指针方法的区别 **/
package practice

import "fmt"

type DeptModeA interface {
	Name() string
	SetName(name string)
}

type DeptModeB interface {
	Relocate(building string, floor uint8)
}

type DeptModeFull interface {
	Name() string
	SetName(name string)
	Relocate(building string, floor uint8)
}

type Dept struct {
	name     string
	building string
	floor    uint8
	Key      string
}

func (d Dept) Name() string {
	return d.name
}

func (d Dept) SetName(name string) {
	d.name = name
}

func (d *Dept) Relocate(building string, floor uint8) {
	d.building = building
	d.floor = floor
}

func Test() {
	dept1 := Dept{
		name:     "MySohu",
		building: "Internet",
		floor:    7}
	switch v := interface{}(dept1).(type) {
	case DeptModeFull:
		fmt.Printf("The dept1 is a DeptModeFull.\n")
	case DeptModeB:
		fmt.Printf("The dept1 is a DeptModeB.\n")
	case DeptModeA:
		fmt.Printf("The dept1 is a DeptModeA.\n")
	default:
		fmt.Printf("The type of dept1 is %v\n", v)
	}
	deptPtr1 := &dept1
	if _, ok := interface{}(deptPtr1).(DeptModeFull); ok {
		fmt.Printf("The deptPtr1 is a DeptModeFull.\n")
	}
	if _, ok := interface{}(deptPtr1).(DeptModeA); ok {
		fmt.Printf("The deptPtr1 is a DeptModeA.\n")
	}
	if _, ok := interface{}(deptPtr1).(DeptModeB); ok {
		fmt.Printf("The deptPtr1 is a DeptModeB.\n")
	}
}
