package main

import "fmt"

func binSearch(arr []int, findData int) int {
	low := 0
	height := len(arr) - 1
	for low <= height {
		mid := (low + height) / 2
		fmt.Println(mid)
		if arr[mid] < findData {
			height = mid - 1
		} else if arr[mid] > findData {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
func main() {
	arr := make([]int, 1024*1024, 1024*1024)
	for i := 0; i < 1024*1024; i++ {
		arr[i] = i + 1
	}
	id := binSearch(arr, 524287)
	if id != -1 {
		fmt.Println(id, arr[id])
	} else {
		fmt.Println("没有找到数据")
	}
}
