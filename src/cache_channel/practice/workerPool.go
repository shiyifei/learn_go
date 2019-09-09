package practice

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id int
	no int
}

type Result struct {
	job Job
	sum int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

/**
	计算一个数各位上的和
 */
func calcu(number int ) int {
	no := number
	sum :=0
	for no != 0 {
		num := no % 10
		sum += num
		no /= 10
	}
	time.Sleep(100*time.Millisecond)
	return sum
}

/**
	将工作池中的计算工作完成后，写入Result信道
 */
func worker(wg * sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, calcu(job.no)}
		results <- output
	}
	wg.Done()
}

/**
	处理所有的工作任务
 */
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i:=0;i<noOfWorkers;i++ {
		wg.Add(1)
		go worker(&wg)  //实际上是将结果写入Result信道的过程
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i:=0; i<noOfJobs; i++ {
		random := rand.Intn(1000)
		input := Job{i, random}
		jobs <- input
	}
	close(jobs)
}

func printResult(done chan bool) {
	for result := range results {
		fmt.Printf("Job id:%d,input random no %d,sum of digits:%d \n", result.job.id, result.job.no, result.sum)
	}
	done <- true
}

func WorkerPool() {
	begin := time.Now()
	go allocate(100)
	go createWorkerPool(100)   //不在协程中运行此函数，有时会出现死锁的错误，调整该数字会发现耗时有变化

	done := make(chan bool)
	go printResult(done)
	<- done
	end := time.Now()
	interval := end.Sub(begin)
	fmt.Println("time interval:",interval)
}







