/**
	这几个结构体都实现了接口中声明的方法，默认都隐式实现了接口
 */

package practice

import "fmt"

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	ProjectName  string
	BiddedAmount int
}

type TimeAndMaterial struct {
	ProjectName string
	NoOfHours   int
	HourlyRate  int
}

func (e FixedBilling) calculate() int {
	return e.BiddedAmount
}

func (e FixedBilling) source() string {
	return e.ProjectName
}

func (e TimeAndMaterial) calculate() int {
	return e.NoOfHours * e.HourlyRate
}

func (e TimeAndMaterial) source() string {
	return e.ProjectName
}

func CalculateNetIncome(ic []Income) {
	var netIncome int = 0
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d \n", income.source(), income.calculate())
		netIncome += income.calculate()
	}
	fmt.Printf("Net income of organisation= $%d\n", netIncome)
}


