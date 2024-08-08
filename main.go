package main

import (
	"fmt"
	"sync"
	"time"
)

func worker() {
	time.Sleep(1 * time.Second)
	fmt.Println("worker runned")
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker()
		}()
	}

	wg.Wait()
}
