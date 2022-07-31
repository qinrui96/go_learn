package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var limit = make(chan struct{})

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return entries
}

func walkDir(dir string, fileSizes chan<- int64, wg *sync.WaitGroup) {
	limit <- struct{}{}
	defer func() {
		wg.Done()
		<-limit
	}()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			go walkDir(subdir, fileSizes, wg)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	wg := sync.WaitGroup{}

	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			wg.Add(1)
			go walkDir(root, fileSizes, &wg)
		}
	}()

	var reminder <-chan time.Time
	if *verbose {
		reminder = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64

	go func() {
		wg.Wait()
		close(fileSizes)
	}()
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-reminder:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
