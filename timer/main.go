package main

func main() {
	//5秒后往通道写数据
	//timer1 := time.NewTimer(5 * time.Second)
	//t1 := time.Now()
	//fmt.Printf("t1:%v\n", t1)
	//t2 := <-timer1.C
	//fmt.Printf("t2:%v\n", t2)

	//实现延时功能
	//(1)
	//time.Sleep(2 * time.Second)
	//(2)
	//timer3 := time.NewTimer(2 * time.Second)
	//<-timer3.C
	//fmt.Println("2秒时间到")
	//（3）
	//<-time.After(3 * time.Second)
	//fmt.Println("3秒时间到")

	//停止定时器
	//timer4 := time.NewTimer(3 * time.Second)
	//go func() {
	//	<-timer4.C
	//	fmt.Println("定时器执行了")
	//}()
	//b := timer4.Stop()
	//if b {
	//	fmt.Println("定时器关闭了")
	//}

	//重置定时器
	//timer5 := time.NewTimer(3 * time.Second)
	//timer5.Reset(1 * time.Second)
	//fmt.Println(time.Now())
	//fmt.Println(<-timer5.C)
	//for {
	//
	//}
}
