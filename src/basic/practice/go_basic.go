package practice

import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

var age int32 = 30
var username string = "shiyf"
var score float32 = 1189.5
var isOk bool = false
var a, b, c int = 4, 5, 6

type Person struct {
	name string
	age  int
}

var P = Person{age:32}

const LENGTH int = 5
const WIDTH = 10

const (
	Unknown = 0
	Male    = 1
	Famale  = 2
)

type T struct {
	t1 byte
	t2 int32
	t3 int64
	t4 string
	t5 bool
}

/**
	defer是后进先出的，panic需要等defer技术后才会向上传递
	出现panic恐慌的时候，会先按照defer的后入先出的顺序执行，最后才会执行panic
	近期有同学遇到多次执行的时候发现panic的执行顺序不定，那么是不是因为panic与defer没有先后关系呢
	那为什么没有加recover()时候，panic执行顺序不定呢？
	defer的执行顺序肯定是FILO的，但是没有被recover的panic协程（线程）可能争夺CPU的顺序比defer快，
	所以造成了这样的情况，也可能是写缓存问题，所以对panic进行recover将其加入到defer队列中。
 */
func defer_call() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("one=",err)
		}
	}()
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		fmt.Println("打印中")
	}()
	defer func() {
		fmt.Println("打印后")
	}()

	panic("触发异常")
}

/**
	打印后
	打印中
	打印前
	panic:触发异常
 */

 type student struct {
 	Name string
 	Age int
 }

 /**
 	你会发现最终生成map实际上并不是我们想要的，里面只含有切片中的最后一个元素了
 	与Java的foreach一样，for range 都是使用副本的方式
 	m[stu.Name]=&stu实际上一致指向同一个指针， 最终该指针的值为遍历的最后一个struct的值拷贝
  */
 func pase_student() {
 	m := make(map[string]*student)
 	arrStudent := []student{
 		{Name:"zhao", Age:22},
 		{Name:"wang", Age:25},
 		{Name:"li", Age:24},
	}
 	for _, stu := range arrStudent {
 		m[stu.Name] = &stu
	}

 	for _, stu := range m {
		fmt.Printf("stu:%+v\n", stu)
	}
 }

 /**
 	将一个切片复制到map中的正确写法
  */
func parse_student() {
	m := make(map[string]*student)
	arrStudent := []student{
		{Name:"zhao", Age:22},
		{Name:"wang", Age:25},
		{Name:"li", Age:24},
	}

	for i:=0;i<len(arrStudent);i++ {
		m[arrStudent[i].Name] = &arrStudent[i]
	}

	for _, stu := range m {
		fmt.Printf("stu:%+v\n", stu)
	}
}


func exam(){
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(20)
	for i:=0;i<10;i++ {
		go func() {
			fmt.Println("A:", i)
			defer wg.Done()
		}()
	}
	for i:=0;i<10;i++ {
		go func(b int) {
			fmt.Println("B:", b)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}

type people struct{}

func (p *people) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *people) ShowB() {
	fmt.Println("showB")
}

type teacher struct {
	people
}

func (t *teacher) ShowB() {
	fmt.Println("teacher showB")
}

func testSlice(){
	s1 := make([]int, 0, 3)
	for i := 0; i < cap(s1); i++ {
		s1 = append(s1, i)
	}
	s2 := s1
	fmt.Println("len(s1):", len(s1))
	for i := 0; i < len(s1); i++ {
		s1[i] = s1[i] + 1
	}
	fmt.Println(s1) //[1 2 3]
	fmt.Println(s2) //[1 2 3]
}

type UserAges struct{
	ages map[string]int
	sync.Mutex
}
func (p *UserAges) Add(name string, age int) {
	p.Lock()
	defer p.Unlock()
	p.ages[name] = age
}

/**
	如果没有Lock()和Unlock()会有读写冲突的发生，竟态条件会有出现
	加入锁，也就意味着当前时刻只能有一个协程来读或写map
 */
func (p *UserAges) Get(name string) int {
	p.Lock()
	defer p.Unlock()
	if ret, ok := p.ages[name]; ok {
		return ret
	}
	return -1
}


func testMultiReadWrite() {
	var wg sync.WaitGroup
	wg.Add(3)
	ages := make(map[string]int)
	userAges := &UserAges{ages:ages}
	go func(p *UserAges){
		defer wg.Done()
		p.Add("zhangfei", 20)
	}(userAges)

	go func(p *UserAges) {
		defer wg.Done()
		v := p.Get("zhangfei")
		fmt.Println("age:",v)
	}(userAges)

	go func(p *UserAges) {
		defer wg.Done()
		v := p.Get("zhangfei")
		fmt.Println("age:",v)
	}(userAges)
	wg.Wait()
}


func Go_basic() {
	testMultiReadWrite()
	fmt.Println("============================")

	testSlice()
	fmt.Println("============================")
	return
	ter := teacher{}
	/**
	这是Golang的组合模式，可以实现OOP的继承。
	被组合的类型People所包含的方法虽然升级成了外部类型Teacher这个组合类型的方法（一定要是匿名字段），但它们的方法(ShowA())调用时接受者并没有发生变化。
	此时People类型并不知道自己会被什么类型组合，当然也就无法调用方法时去使用未知的组合者Teacher类型的功能。
	 */
	ter.ShowA()
	ter.ShowB()
	fmt.Println("============================")
	exam()
	fmt.Println("============================")
	pase_student()
	fmt.Println("============================")
	parse_student()
	fmt.Println("============================")
	defer_call()

	var n4, name, n5 = 100,"tom",888
	fmt.Println("n4:",n4, "name:",name, "n5:",n5)



	fmt.Printf("Hello,%s,your score s:%.2f,what are you doing now?age:%d,are you ok?%t \n",
		username, score, age, isOk)

	score = score + float32(age)  //强制类型转换，不然会报错
	//unsafe.Sizeof(score)返回占用字节数
	fmt.Printf("score is:%.2f, type of score is:%T, size of score is:%d \n", score, score, unsafe.Sizeof(score))


	//unsafe.Sizeof(age)返回占用字节数，32位系统下大小是 32 位（4 字节）, 64位系统下，age会占用 64 位（8 字节）的大小
	fmt.Printf("type of age is:%T, size of age is:%d \n", age, unsafe.Sizeof(age))

	const strName string = "shiyf,what are you doing now?"
	fmt.Printf("value of strName:%s,type of strName is:%T, size of strName is:%d \n", strName, strName, unsafe.Sizeof(strName))

	a, b := 6, 8 //a,b,c 局部赋值优先

	// a,b,c := 1,2,3
	fmt.Println(a, b, c)

	fmt.Printf("person name:%s,age:%d \n", P.name, P.age)

	//P := Person{name:"wanggengke"}
	var P Person
	P.name = "areyouok"
	P.age = 33
	fmt.Printf("person name:%s,age:%d \n", P.name, P.age)
	fmt.Println(P)

	var _, ret, retStr = numbers()  //调用多个返回值的方法
	fmt.Println(ret, retStr)

	var area int
	area = LENGTH * WIDTH
	fmt.Printf("area is:%d \n", area)

	fmt.Printf("male is:%d \n", Male)

	var str string = "abc"
	fmt.Println(str, len(str), unsafe.Sizeof(str))

	//unsafe包其实是指针
	fmt.Println("----------unsafe.Pointer---------")
	t := &T {1,2,3, "this is an example", true}
	ptr := unsafe.Pointer(t)  //获取变量t的通用指针
	t1 := (*byte)(ptr)		//unsafe.Pointer可以和普通指针进行相互转换
	fmt.Println(*t1)	 	//t.t1当前的值

	//unsafe.Pointer 可以和 uintptr 进行相互转换 uintptr(ptr)
	t2 := (*int32)(unsafe.Pointer(uintptr(ptr) + unsafe.Offsetof(t.t2)))  	//指针移动到t.t2位置
	fmt.Println(*t2)  //t.t2当前的值
	*t2 = 99	//实际上会更改t.t2的值
	fmt.Println(t) //可以直接打印一个stuct对象,其中的元素会用空格分开

	t3 := (*int64)(unsafe.Pointer(uintptr(ptr) + unsafe.Offsetof(t.t3)))	//指针移动到t.t3位置
	fmt.Println(*t3)  //t.t3当前的值
	*t3 = 123  //实际上会更改t.t3的值
	fmt.Println(t)

	var len float64 = 20.0
	var width float64 = 30.0
	var f11, f12 float64
	f11,f12 = rectProps(len, width)  //调用多个返回值的方法

	fmt.Println("area=",f11,",perimeter=",f12)

	printRows()

	useSwitch()
}

/**
返回多个值的建议写法,比较好理解
 */
func numbers() (int, int, string) {
	a, b, c := 1, 2, "are you ok?"
	return a, b, c
}

/**
返回多个值的方法，不建议这样写，因为不好理解
 */
func rectProps(length,width float64)(area, perimeter float64) {
	area = length * width
	perimeter = 2*(length+width)
	return
}


func printRows() {
	for i:=10;i<20;i++ {
		for j:=1;j<=10;j++ {
			if  j-1 == i-10 {
				fmt.Printf("%d * %d = %d \n", i, j, i*j)
			}
		}
	}
}

func useSwitch() {
	letter := "i"
	//switch不需要在每个case语句后写break;
	switch letter {
	case "a","e","i","o","u":
		fmt.Println("这是一个元音字母")
	default:
		fmt.Println("这个一个非元音字母")
	}
}
