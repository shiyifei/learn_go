package search

import (
	"encoding/json"
	"os"
)

type Feed struct {
	Name string `json:"site"`
	Url string `json:"link"`
	Type string `json:"type"`
}

const dataFile = "/var/www/html/learn_go/src/sample/data/data.json"

/**
	读取并反序列化源数据文件，实际上是初始化请求源头
 */
func RetriveFeeds() ([]*Feed, error){
	file, err := os.Open(dataFile)
	if err != nil {
		return nil,err
	}
	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)   //将file文件转为结构体实例
	return feeds, err
}