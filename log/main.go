package main

import (
	"fmt"
	"log"
	"os"
)

//logger会打印每条日志信息的日期、时间，默认输出到系统的标准错误。
//Fatal系列函数会在写入日 志信息后调用os.Exit(1)。
//Panic系列函数会在写入日志信息后panic。
func main() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志")
	log.SetPrefix("[fatal]")
	log.Fatalln("错误日志")
	log.Panicln("这是一条会触发panic的日志。")
}
