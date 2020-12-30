package practice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Resp struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data []interface{} `json:"data"`
}

var resp *Resp

func init() {
	resp = new(Resp)
	resp.Code = 0
	resp.Msg = "ok"
	resp.Data = []interface{}{}
}

func handleIndex(w http.ResponseWriter, r * http.Request) {
	input := r.ParseForm()
	fmt.Printf("%#v", input)
	fmt.Println("PATH:", r.URL.Path)
	fmt.Println("SCHEME", r.URL.Scheme)
	fmt.Println("METHOD:", r.Method)
	fmt.Fprintf(w, "<h1>It works!</h1>")
}

/**
 * @Description: 处理用户发来的处理预订单的请求，要求传入三个参数:json,delay,taskFunc
 * json:表示预订单数据
 * delay:表示该预订单将在多少秒后执行
 * taskFunc:表示该预订单的处理函数，需要在Processor.go文件中定义该方法，先定义后使用
 * @param w
 * @param r
 */
func handleProcess(w http.ResponseWriter, r * http.Request) {
	strJson := r.PostFormValue("json")
	fmt.Printf("post:%#v", strJson)
	strJson = strings.TrimSpace(strJson)

	if strJson == "" {
		resp.Code = -1
		resp.Msg = "请传入json参数"
		bytes, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(bytes))
		return
	}

	taskFunc := r.PostFormValue("taskFunc")
	fmt.Printf("taskFunc:%#v", taskFunc)
	taskFunc = strings.TrimSpace(taskFunc)

	if taskFunc == "" {
		resp.Code = -1
		resp.Msg = "请传入taskFunc参数"
		bytes, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(bytes))
		return
	}

	strDelay := r.PostFormValue("delay")
	strDelay = strings.TrimSpace(strDelay)
	fmt.Printf("delay:%#v", strDelay)

	delay, err := strconv.Atoi(strDelay)
	if err != nil {
		resp.Code = -1
		resp.Msg = "参数delay格式错误，请传入整数"
		bytes, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(bytes))
		return
	}

	var task Task
	task.Delay = delay
	task.JsonData = strJson
	task.TaskFunc = taskFunc

	SendOneTask(task)
	fmt.Fprintf(w, task.JsonData)
}

func CreateServer() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/pre_order/process", handleProcess)
	err := http.ListenAndServe(":8666", nil)
	if err != nil {
		fmt.Println(err)
	}
}
