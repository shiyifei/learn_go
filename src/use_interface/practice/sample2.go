package practice

type SalaryCalculator interface {
	CalculateSalary() int
}

//公司的创始员工
type Permanent struct {
	EmpId int
	Basicpay int
	Pf int
}

//公司的合同工
type Contract struct {
	EmpId int
	Basicpay int
}

//针对创始员工计算支出
func (p Permanent) CalculateSalary() int {
	return p.Basicpay + p.Pf
}

//针对合同工计算支出
func (p Contract) CalculateSalary() int {
	return p.Basicpay
}

//统计一些员工的总支出
func TotalExpense(s []SalaryCalculator) int {
	expense := 0
	for _, v := range s {
		expense += v.CalculateSalary()
	}
	return expense
}



