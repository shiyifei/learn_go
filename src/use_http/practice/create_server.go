package practice

import (
	"fmt"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r * http.Request) {
	r.ParseForm()
	fmt.Println("PATH:", r.URL.Path)
	fmt.Println("SCHEME", r.URL.Scheme)
	fmt.Println("METHOD:", r.Method)
	fmt.Println()

	fmt.Fprintf(w, "<h1>This is home page,what are you doing now?</h1>")

}

func CreateServer() {
	http.HandleFunc("/", handleIndex)
	err := http.ListenAndServe(":8666", nil)
	if err != nil {
		fmt.Println(err)
	}
}