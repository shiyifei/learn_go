package practice

import (
	"fmt"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r * http.Request) {
	input := r.ParseForm()
	fmt.Printf("%#v", input)
	fmt.Println("PATH:", r.URL.Path)
	fmt.Println("SCHEME", r.URL.Scheme)
	fmt.Println("METHOD:", r.Method)
	fmt.Println()

	var task Task
	task.Delay = 16
	task.JsonData = `{"pre_order_id":22, "userid":123}`
	task.TaskFunc = "ProcessOrder"

	SendOneTask(task)

	fmt.Fprintf(w, "<h1>"+task.JsonData+"</h1>")

}

func CreateServer() {
	//defer wg.Done()
	http.HandleFunc("/", handleIndex)
	err := http.ListenAndServe(":8666", nil)
	if err != nil {
		fmt.Println(err)
	}
}
