package main

import (
	"fmt"
	"time"
)

//ticker 时间到了多次执行
func main() {
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	go func() {
		for {
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				ticker.Stop()
			}
		}
	}()
	for {

	}
}
