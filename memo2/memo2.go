package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type mome struct {
	f     Func
	cache map[string]result
}

type result struct {
	value interface{}
	err   error
}

var wg sync.WaitGroup

type Func func(Key string) (interface{}, error)

func New(f Func) *mome {
	return &mome{f, map[string]result{}}
}

func (m *mome) Get(key string) (interface{}, error) {
	v, ok := m.cache[key]
	if ok {
		return v.value, v.err
	}

	ret, err := m.f(key)
	m.cache[key] = result{ret, err}
	return ret, err
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	m := New(httpGetBody)
	URLs := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://www.baidu.com",
		"https://google.com",
		"https://www.baidu.com",
		"https://google.com",
	}
	for _, url := range URLs {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			ret, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Println(url, time.Since(start), len(ret.([]byte)))
		}(url)
	}
	wg.Wait()
}
