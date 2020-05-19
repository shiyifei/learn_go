package practice

import "sync"

var syncMap sync.Map

func UseSyncMap() {
	syncMap.Load("order:1")
}
