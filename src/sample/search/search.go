package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	feeds, err := RetriveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	results := make(char *Result)

	var wg sync.WaitGroup

	wg.Add(len(feeds))

	for _,feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			wg.Done()
		}(matcher, feed)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	Display(results)

}

func Register(feedType string, matcher Matcher) {
	if _,exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}
	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}