package main

import (
	"fmt"
	"sync"
)

var x int64
var lock sync.Mutex
var wg sync.WaitGroup

func add() {
	for i := 0; i < 10; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
