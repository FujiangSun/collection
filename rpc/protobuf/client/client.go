package client

import (
	"../message/message"
	"fmt"
	"net/rpc"
	"testing"
)

func TestClient(t *testing.T) {
	// 建立连接
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if err != nil {
		panic(err.Error())
	}

	req := message.EncryptionRequest{Str: "mclink"}
	resp := message.EncryptionResult{}
	// 同步调用
	err = client.Call("EncryptionService.Encryption", &req, &resp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(resp.GetResult())

}
