package practice

import (
	"fmt"
	"math"
	"reflect"
	"sync"
	"time"
)

type Task struct {
	CycleRound  int		    //轮次
	Delay    int			//延时多少秒执行
	JsonData string			//传入的json字符串
	TaskFunc string			//传入的回调函数 需要在Processor.go中定义
}

type Slot struct {
	TaskList []Task
}

const (
	TEMPLATE = "2006-01-02 15:04:05.000" //golang时间格式 加入毫秒显示 如果是6个0，则是微秒，3个0表示不足三位时左侧会补齐0
	MAX      = 30
)

type CycleList struct {
	CurrentIndex int
	Slots        [MAX]Slot
}

var wg sync.WaitGroup
var list *CycleList

/**
 * @Description: 初始化方法，用于测试
 */
func init() {
	list = new(CycleList)
	for i := 0; i < MAX; i++ {
		list.Slots[i].TaskList = make([]Task, 0)
	}
	list.CurrentIndex = 0

	var task Task
	task.Delay = 22
	task.CycleRound = getCycleRound(task.Delay)
	task.JsonData = `{"pre_order_id":2, "user_id":11,"op_type":100065,"app_id":301}`
	task.TaskFunc = "ProcessOrder"
	index := getPosition(list.CurrentIndex, task.Delay)
	list.Slots[index].TaskList = append(list.Slots[index].TaskList, task)

	task.Delay = 44
	task.CycleRound = getCycleRound(task.Delay)
	task.JsonData = `{"pre_order_id":3, "user_id":22,"op_type":100066,"app_id":301}`
	task.TaskFunc = "ProcessOrder"
	index = getPosition(list.CurrentIndex, task.Delay)
	list.Slots[index].TaskList = append(list.Slots[index].TaskList, task)

	fmt.Printf("222, list.Slots[index].TaskList:%#v,currentIndex:%d, index:%d \n", list.Slots[index].TaskList, list.CurrentIndex, index)

	task.Delay = 66
	task.CycleRound = getCycleRound(task.Delay)
	task.JsonData = `{"pre_order_id":4, "user_id":33,"op_type":100067,"app_id":301}`
	task.TaskFunc = "ProcessOrder"
	index = getPosition(list.CurrentIndex, task.Delay)
	list.Slots[index].TaskList = append(list.Slots[index].TaskList, task)
}

/**
 * @Description: 启动定时器
 */
func StartTimer() {
	wg.Add(1)
	ticker := time.NewTicker(time.Second)
	go func(t *time.Ticker) {
		defer wg.Done()
		for {
			select {
			case received := <-t.C: //注意这里的返回值是时间类型
				fmt.Printf(
					"get ticker,received:%s,current index:%d, tasklist:%#v \n",
					received.Format(TEMPLATE),
					list.CurrentIndex, list.Slots[list.CurrentIndex].TaskList)

				//如果当前solt内有任务,要根据传入的参数来处理
				arrTask := list.Slots[list.CurrentIndex].TaskList
				lenTasks := len(arrTask)
				if lenTasks > 0 {
					var arrDelIndex = make([]int, 0)
					for i := 0; i < lenTasks; i++ {
						objTask := arrTask[i]
						//判断轮次是否为0，如果大于0表示本轮不处理，只将轮次减1即可。
						if objTask.CycleRound > 0 {
							list.Slots[list.CurrentIndex].TaskList[i].CycleRound--
						} else {
							wg.Add(1)
							processData(&wg, objTask.TaskFunc, objTask.JsonData)
							arrDelIndex = append(arrDelIndex, i)
						}
					}
					//从后往前删除切片中的已处理任务
					delCount := len(arrDelIndex)
					if delCount > 0 {
						for i := delCount - 1; i >= 0; i-- {
							arrTask = append(arrTask[:i], arrTask[i+1:]...)
						}
						list.Slots[list.CurrentIndex].TaskList = arrTask
					}
				}
				if list.CurrentIndex == MAX-1 {
					list.CurrentIndex = 0
				} else {
					list.CurrentIndex += 1
				}
			}
		}
	}(ticker)
	wg.Wait()
}

/**
 * @Description: 根据传入的延迟时间（s）来计算轮次
 * @param delay
 */
func getCycleRound(delay int) int {
	return int(math.Floor(float64(delay)/MAX))
}

func getPosition(currentIndex, delay int) int {
	return (currentIndex + delay) % MAX
}
//处理其他比较费时的任务
func processData(wg *sync.WaitGroup, funcName string, params string) {
	defer wg.Done()
	fmt.Println("funcName is:", funcName)
	InvokeObjectMethod(new(Processor), funcName, params)
}

 /**
 * @Description: 反射调用指定的处理方法
 * @param object Processor
 * @param methodName 方法名
 * @param args []interface{}
 */
func InvokeObjectMethod(object interface{}, methodName string, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, param := range args {
		inputs[i] = reflect.ValueOf(param)
	}
	defer MyRecover()
	reflect.ValueOf(object).MethodByName(methodName).Call(inputs)
}

 /**
 * @Description: 向环形队列中指定位置的任务切片中插入一个任务
 * @param task
 */
func SendOneTask(task Task) {
	task.CycleRound = getCycleRound(task.Delay)
	index := getPosition(list.CurrentIndex, task.Delay)
	list.Slots[index].TaskList = append(list.Slots[index].TaskList, task)
}
