/**
	本示例演示如何使用interface和struct来实现类似其他语言的多态，不同的类针对同一个方法有不同的功能
 */
package main

import (
	UsePolym "use_polym/practice"
)


func main() {
	project1 := UsePolym.FixedBilling{"build house", 100000}
	project2 := UsePolym.FixedBilling{"build road", 5000}
	project3 := UsePolym.TimeAndMaterial{"car service", 8, 200}
	project4 := UsePolym.Advertisement{"sale book", 1, 10000}
	incomes := []UsePolym.Income{project1, project2, project3, project4}
	UsePolym.CalculateNetIncome(incomes)
}
