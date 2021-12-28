package client

import (
	"context"
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	//1、Dail连接
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	encryptionServiceClient := message.NewEncryptionServiceClient(conn)

	orderRequest := &message.EncryptionRequest{Str: "mclink"}
	result, err := encryptionServiceClient.Encryption(context.Background(), orderRequest)
	if result != nil {
		fmt.Println(result.Result)
	}
}
