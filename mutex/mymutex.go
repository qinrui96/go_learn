package main

import "sync"

var l sync.RWMutex

func main() {
	l.RLock()
	defer l.RUnlock()
}
