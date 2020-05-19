package practice

type SalaryCalculator interface {
	CalculateSalary() int
}

//公司的创始员工
type Permanent struct {
	EmpId int
	BasicPay int
	Pf int
}

//公司的合同工
type Contract struct {
	EmpId int
	BasicPay int
}

//针对创始员工计算支出
func (p Permanent) CalculateSalary() int {
	return p.BasicPay + p.Pf
}

//针对合同工计算支出
func (p Contract) CalculateSalary() int {
	return p.BasicPay
}

//统计一些员工的总工资支出
func TotalExpense(s []SalaryCalculator) int {
	expense := 0
	for _, v := range s {
		expense += v.CalculateSalary()
	}
	return expense
}



