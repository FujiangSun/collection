package serve

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"net"
)

// EncryptionServiceImpl 加密服务实现类
type EncryptionServiceImpl struct {
}

// Encryption 加密程序
func (es *EncryptionServiceImpl) Encryption(ctx context.Context, request *message.EncryptionRequest) (*message.EncryptionResult, error) {
	md5Str := ToMd5(request.GetStr())
	res := message.EncryptionResult{Result: md5Str}
	return &res, nil
}

// ToMd5 简单封装的 md5
func ToMd5(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}

func main() {
	// 创建一个 grpc server
	server := grpc.NewServer()

	message.RegisterEncryptionServiceServer(server, new(EncryptionServiceImpl))

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}
	_ = server.Serve(lis)
}
