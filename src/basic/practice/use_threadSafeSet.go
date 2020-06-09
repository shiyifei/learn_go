package practice

import (
	"fmt"
	"sync"
)

type threadSafeSet struct{
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{}, len(set.s))

	go func() {
		set.RLock()
		for i:=0; i<len(set.s); i++ {
			ch <- set.s[i]
			fmt.Println("item:",set.s[i])
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}


func ThreadSafeSet() {
	set := threadSafeSet{
		s:[]interface{}{"aa", "bb", "cc"},
	}

	v := <-set.Iter()
	fmt.Printf("v:%+v \n", v)
}