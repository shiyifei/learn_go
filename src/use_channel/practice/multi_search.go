package practice

import (
	"fmt"
	"math"
	"math/rand"
	"sync/atomic"
	"time"
)

var arrSource []int32

func init() {
	arrSource = make([]int32, 0)

	rand.Seed(time.Now().UnixNano())
	for i:=0; i<60;i++ {
		arrSource = append(arrSource, int32(rand.Int()%100))
	}
	fmt.Printf("arrSource:%+v \n", arrSource)
}

func concurrencySearch(arrSource []int32, target int32) {
	chanRet := make(chan int)
	chanFail := make(chan int)

	numOfParts := 4
	pageSize := int(math.Ceil(float64(len(arrSource)/4)))
	var offset int =0
	var end int = 0
	for i:=0; i<numOfParts; i++ {
		offset = i*pageSize
		end = offset+pageSize
		if end >len(arrSource) {
			end = len(arrSource)
		}
		arrInput := arrSource[offset:end]

		go func(arrInput []int32, offset int) {
			for k,v := range arrInput {
				if atomic.LoadInt32(&target) == v {
					chanRet <- k+offset
					return
				}
			}
			chanFail <- 1
		}(arrInput, offset)
	}
	var failedTimes int32 = 0
	for {
		select {
		case ret := <-chanRet:
				fmt.Printf("target:%d is found in arrSource[%d]\n", target, ret)
				return
			case num := <-chanFail:
				atomic.StoreInt32(&failedTimes,failedTimes+int32(num))
				if atomic.LoadInt32(&failedTimes) == int32(numOfParts) {
					fmt.Printf("target:%d is not found in arrSource \n", target)
					return
				}
		}
	}
}

func DoSearch() {
	var target int32 = 50
	concurrencySearch(arrSource, target)
}
