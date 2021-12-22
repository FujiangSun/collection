package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//file, err := os.Create("xxx.log")
	//defer file.Close()
	//if err != nil {
	//	fmt.Println("create file failed")
	//}
	//for i := 0; i < 5; i++ {
	//	file.WriteString("ab\n")
	//	file.Write([]byte("cd\n"))
	//}
	file2, err := os.Open("a.txt")
	defer file2.Close()
	if err != nil {
		fmt.Println("open file failed")
		return
	}
	var buf [128]byte
	var content []byte
	for {
		n, err := file2.Read(buf[:])
		if err != nil {
			fmt.Println("read failed")
			return
		}
		if err == io.EOF {
			break //读取结束
		}
		content = append(content, buf[:n]...)

	}
	fmt.Println(string(content))
}
