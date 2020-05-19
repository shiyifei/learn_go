package practice

import "sync"

var syncMap sync.Map

func UseSyncMap() {
	ret := syncMap.Load("order:1")
}
