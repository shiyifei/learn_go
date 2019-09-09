package search

import (
	"fmt"
	"log"
)

type Result struct {
	Field string
	Content string
}

/**
	声明了一个接口类型，需要实现Search方法
 */
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

/**
	将最终匹配到的结果写入到Result信道中
 */
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}
	for _,result := range searchResults {
		results <- result
	}
}

/**
	读取并打印信道中的结果
 */
func Display(results chan *Result) {
	for result := range results {
		fmt.Println("%s:\n %s:\n\n", result.Field, result.Content)
	}
}