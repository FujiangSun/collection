package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	//if len(os.Args) > 0 {
	//	for i, arg := range os.Args {
	//		fmt.Printf("args[%d]=%v\n", i, arg)
	//	}
	//}

	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", true, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的间隔时间")

	flag.Parse()
	fmt.Println(name, age, married, delay)
	fmt.Println(flag.Args())
	fmt.Println(flag.NFlag())
}
