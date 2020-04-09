package practice

import ("fmt" 
"math")


func max(num1,num2 int) int {
	var result int
	if num1>num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func Go_Func() {
	var a int = 80
	var b int = 79
	var ret int
	ret = max(a, b)
	fmt.Println("max(",a,",",b,")=",ret)

	str1,str2 := swap1("hello","world")
	fmt.Println(str1,str2)

	fmt.Println("交换前的值 a=",a,",b=",b)
	ret = swap2(a, b)
	fmt.Println("交换后的值 a=,",a,",b=",b)

	fmt.Printf("交换前的值a=%d,b=%d \n", a, b)
	swap3(&a, &b)
	fmt.Printf("交换后的值a=%d,b=%d \n", a, b)

	fmt.Printf("交换前的值a=%d,b=%d \n", a, b)
	swap4(&a, &b)
	fmt.Printf("交换后的值a=%d,b=%d \n", a, b)

	getSqrtRoot := func(input float64) float64 {
		return math.Sqrt(input)
	}

	var output float64 
	output = getSqrtRoot(9)

	fmt.Printf("output=%.2f \n", output)
	fmt.Printf("output=%.2f \n", getSqrtRoot(9))


}

func swap1(a, b string) (string,string) {
	return b,a
}

func swap2(a, b int) int {
	var temp int
	temp = a
	a = b
	b = temp
	return temp
}

func swap3(a *int, b *int) int {
	var temp int 
	temp = *a
	*a = *b 
	*b = temp
	return temp
}

func swap4(a *int, b *int) {
	*a,*b = *b,*a

}
