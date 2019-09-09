package main

import (
	cachedChan "cache_channel/practice"
)

func main() {
	cachedChan.ReadWrite()

	cachedChan.HowToCache()

	cachedChan.WorkerPool()
}
