package service

import(
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)



/**
	输出异常信息
 */
func FailOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s \n", msg, err)
	}
}

type JsonData struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}


func SendPostRequest(url, message string) (bool, error){
	var jsonStr = []byte(message)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 3 *time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		//panic(err)
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, errors.New(resp.Status)
	}

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	var data JsonData
	json.Unmarshal(body, &data)
	if data.Code == 1 {
		return true, nil
	} else {
		return false, errors.New(data.Msg)
	}
}