package practice

import (
	"encoding/json"
	"time"
)

type Processor struct {

}

func (o *Processor) ProcessOrder(jsonData string) {
	time.Sleep(2100 * time.Millisecond)
	var mapStr map[string]interface{}
	json.Unmarshal([]byte(jsonData), &mapStr)
	ProcessData(mapStr)
}