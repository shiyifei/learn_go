package search

import (
	"encoding/json"
	"os"
)

type Feed struct {
	Name string `json:"site"`
	Url string `json:"url"`
	Type string `json:"type"`
}

const dataFile = "data/data.json"

/**
	读取并反序列化源数据文件
 */
func RetrieveFeeds() ([]*Feed, error){
	file, err := os.Open(dataFile)
	if err != nil {
		return nil,err
	}
	defer file.Close()

	var feeds []*Feed
	err = json.newDecoder(file).Decode(&feeds)
	return feeds, err
}