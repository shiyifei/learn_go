package practice

/**
	如何保证Map在多线程使用情况下的安全性，这是个很重要的问题
	本示例讲解了线程安全的Map的用法 sync.Map
 */

/**
	func (m *Map) Delete(key interface{})
	func (m *Map) Load(key interface{}) (value interface{}, ok bool)
	func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
	func (m *Map) Range(f func(key, value interface{}) bool)
	func (m *Map) Store(key, value interface{})
 */

import (
	"fmt"
	"sync"
)

var syncMap sync.Map

func UseSyncMap() {
	val, ok := syncMap.Load("order:1")
	fmt.Println("val=", val, "ok=", ok)

	syncMap.Store("order:1", 22)
	syncMap.Store("order:2", 12)
	syncMap.Store("order:3", 15)

	//已存值的情况下不允许修改值了
	oldval, loaded := syncMap.LoadOrStore("order:3", 18)
	fmt.Println("key=order:3,已保存的值:",oldval,"之前是否已经有存值:",loaded)

	oldval, loaded = syncMap.LoadOrStore("order:4", 19)
	fmt.Println("key=order:4,已保存的值:",oldval,"之前是否已经有存值:",loaded)

	//return false 会导致只输出一个元素
	//return true 才会迭代输出所有的元素
	syncMap.Range(func(k, v interface{}) bool {
		fmt.Printf("k:%s,v:%d \n", k.(string), v.(int))
		return false
	})
}
