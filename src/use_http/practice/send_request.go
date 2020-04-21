package practice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

func SendRequest() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	headers := resp.Header

	for k, v := range headers {
		fmt.Printf("k=%v, v=%v \n", k, v)
	}

	fmt.Printf("resp status:%s, statusCode:%d \n", resp.Status, resp.StatusCode)
	fmt.Printf("resp Proto:%s \n", resp.Proto)
	fmt.Printf("resp Content length:%d \n", resp.ContentLength)
	fmt.Printf("resp transfer encoding:%v\n", resp.TransferEncoding)
	fmt.Printf("resp Uncompressed %t \n", resp.Uncompressed)
	fmt.Println("type of response body:", reflect.TypeOf(resp.Body))
	buf := bytes.NewBuffer(make([]byte, 0, 512))
	length, _ := buf.ReadFrom(resp.Body)
	fmt.Println(len(buf.Bytes()))
	fmt.Println(length)
	fmt.Println(string(buf.Bytes()))
}

func HttpDo() {
	//设置发起请求的超时时间
	timeout := time.Duration(5*time.Second)
	client := &http.Client{
		Timeout: timeout,
	}
	loginUrl := "https://www.maimaiche.com/loginRegister/login.do"
	req, err := http.NewRequest("POST", loginUrl, strings.NewReader("mobile=69d0ca7a43a0716285742b2f80700b4416d19cc055306c950fa7d835fae521ddd0b261915a8c7aa7fb259ab5f5bc04b9e571a4adcc4e38d57897fd92e2dcf4928ab1565d32f309025b1232a6582f1a5b4726d3626e9b335034b5bc0ee96d1d191e3e0b8858294c01199305cd490b11df03e422232f9b4c692d0a04cc4df6b363&isRemberPwd=1"))
	if err != nil {
		log.Println("in httpDo(),err:", err)
		return
	}
	//req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	//req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	//req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	agent := "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36"
	//agent := "PostmanRuntime/7.24.1"

	//如果不设置User-Agent，有时可能获取不到正确结果
	req.Header.Set("User-Agent", agent)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("in HttpDo(), body:", string(body))

	fmt.Println(req.Header.Get("Content-Type"))

	type Result struct {
		Msg    string `json:"msg"`
		Status string `json:"status"`
		Obj    string `json:"obj"`
	}

	result := &Result{}
	json.Unmarshal(body, result)
	if result.Status == "1" {
		fmt.Println(result.Msg)
	} else {
		fmt.Println("login error")
	}

	fmt.Printf("%+v \n", result)
}

func HttpPost() {
	loginUrl := "https://www.maimaiche.com/loginRegister/login.do"
	resp, err := http.Post(loginUrl, "application/x-www-form-urlencoded", strings.NewReader("mobile=69d0ca7a43a0716285742b2f80700b4416d19cc055306c950fa7d835fae521ddd0b261915a8c7aa7fb259ab5f5bc04b9e571a4adcc4e38d57897fd92e2dcf4928ab1565d32f309025b1232a6582f1a5b4726d3626e9b335034b5bc0ee96d1d191e3e0b8858294c01199305cd490b11df03e422232f9b4c692d0a04cc4df6b363&isRemberPwd=1"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("in HttpPost(),resp:", string(body))
}

func HttpPostForm() {
	postParam := url.Values{
		"mobile":      {"69d0ca7a43a0716285742b2f80700b4416d19cc055306c950fa7d835fae521ddd0b261915a8c7aa7fb259ab5f5bc04b9e571a4adcc4e38d57897fd92e2dcf4928ab1565d32f309025b1232a6582f1a5b4726d3626e9b335034b5bc0ee96d1d191e3e0b8858294c01199305cd490b11df03e422232f9b4c692d0a04cc4df6b363"},
		"isRemberPwd": {"1"},
	}
	loginUrl := "https://www.maimaiche.com/loginRegister/login.do"
	resp, err := http.PostForm(loginUrl, postParam)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("in HttpPostForm(),resp:", string(body))
}
