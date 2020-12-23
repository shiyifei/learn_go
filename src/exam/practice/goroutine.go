package practice

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

func printChar(times int, begin, end byte) {
	for count:=0; count<times; count++ {
		for char := begin; char <= end; char++ {
			fmt.Printf("%c", char)
		}
	}
}

func getBytes(times int, begin, end byte) []byte {
	bytes := make([]byte,0)
	for count:=0; count<times; count++ {
		for char := begin; char <= end; char++ {
			bytes = append(bytes, char)
		}
	}
	return bytes
}

var times int
func init() {
	times = 5000
}

func concurrency() {
	begin := time.Now()
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

	wg.Add(3)

	//打印100次小写字母
	go func() {
		defer wg.Done()
		printChar(times, 'a', 'a')
	}()

	//打印100次大写字母
	go func() {
		defer wg.Done()
		printChar(times, 'A', 'A')
	}()

	//打印100次数字
	go func() {
		defer wg.Done()
		printChar(times, '0', '0')
	}()

	wg.Wait()

	diff :=  time.Now().Sub(begin)
	fmt.Println("\n 并发打印完所有字母，耗时：",diff)
}

func serial() {
	begin := time.Now()
	printChar(times, 'a', 'z')
	printChar(times, 'A', 'Z')
	printChar(times, '0', '9')
	diff :=  time.Now().Sub(begin)
	fmt.Println("\n 串行打印完所有字母，耗时：",diff)
}

func writeFile(filePath string, times int, begin, end byte) {
		f, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		dataByte := getBytes(times, begin, end)
		total, err := f.Write(dataByte)
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		fmt.Println(total, "bytes written successfully")
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
}
/**
	并发写三个文件，内容相同
 */
func writeFileConcurrency() {
	begin := time.Now()
	//runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

	wg.Add(3)

	//打印100次小写字母
	go func() {
		defer wg.Done()
		file  := "/home/shiyf/1.txt"
		writeFile(file, times, 'a', 'z')
	}()

	//打印100次大写字母
	go func() {
		defer wg.Done()
		file := "/home/shiyf/2.txt"
		writeFile(file, times, 'A', 'Z')
	}()

	//打印100次数字
	go func() {
		defer wg.Done()
		file := "/home/shiyf/3.txt"
		writeFile(file, times, '0', '9')
	}()
	wg.Wait()

	diff :=  time.Now().Sub(begin)
	fmt.Println("\n 并发写入所有文件，耗时：",diff)
}

/**
	串行写三个文件
 */
func writeFileSerial() {
	begin := time.Now()
	file  := "/home/shiyf/11.txt"
	writeFile(file, times, 'a', 'z')
	file = "/home/shiyf/12.txt"
	writeFile(file, times, 'A', 'Z')
	file = "/home/shiyf/13.txt"
	writeFile(file, times, '0', '9')

	diff :=  time.Now().Sub(begin)
	fmt.Println("\n 串行写入所有文件，耗时：",diff)
}

func TestConcurrency() {
	//concurrency()
	//serial()

	writeFileConcurrency()
	writeFileSerial()
}

