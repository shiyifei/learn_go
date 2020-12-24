/**
 * 本示例演示了如何通过加锁来防止有竞态条件出现，多个协程可以同时读写一个map
 * @Author:shiyf
 * @Date: 2020/12/24 23:51
 **/

package practice

import (
	"fmt"
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}
func (p *UserAges) Add(name string, age int) {
	p.Lock()
	defer p.Unlock()
	p.ages[name] = age
}

/**
	如果没有Lock()和Unlock()会有读写冲突的发生，竟态条件会有出现
	加入锁，也就意味着当前时刻只能有一个协程来读或写map
 */
func (p *UserAges) Get(name string) int {
	p.Lock()
	defer p.Unlock()
	if ret, ok := p.ages[name]; ok {
		return ret
	}
	return -1
}


func MultiReadWrite() {
	var wg sync.WaitGroup
	wg.Add(3)
	ages := make(map[string]int)
	userAges := &UserAges{ages:ages}
	go func(p *UserAges){
		defer wg.Done()
		p.Add("zhangfei", 20)
	}(userAges)

	go func(p *UserAges) {
		defer wg.Done()
		v := p.Get("zhangfei")
		fmt.Println("age:", v)
	}(userAges)

	go func(p *UserAges) {
		defer wg.Done()
		v := p.Get("zhangfei")
		fmt.Println("age:", v)
	}(userAges)
	wg.Wait()
}