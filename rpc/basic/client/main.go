package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 建立连接
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if err != nil {
		panic(err.Error())
	}
	// 参数
	req := "mclink"
	var resp *string
	// 同步调用
	//err = client.Call("EncryptionUtil.Encryption", req, &resp)
	//if err != nil {
	//	panic(err.Error())
	//}
	//fmt.Println(*resp)

	// 异步调用
	syncCall := client.Go("EncryptionUtil.Encryption", req, &resp, nil)
	// 阻塞，异步调用成功后解除阻塞
	replayDone := <-syncCall.Done
	fmt.Println(replayDone)
	fmt.Println(*resp)

}
