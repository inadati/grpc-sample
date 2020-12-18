package main

import (
	"context"
	"fmt"
	"net"

	"log"

	"google.golang.org/grpc"

	// .protoから生成されたコードをimportしている
	// 今回は筆者の$GOPATH内に作成したので適宜プロジェクトを作成したパスに合わせる
	pb "github.com/inadati/grpc-sample/protobuf/hellogrpc"
)

// gRPC server struct
type server struct {
}

// .protoで定義したGreetServerを定義している
func (s *server) GreetServer(ctx context.Context, p *pb.GreetRequest) (*pb.GreetMessage, error) {
	log.Printf("Request from: %s", p.Name)
	return &pb.GreetMessage{Msg: fmt.Sprintf("Hello, %s. ", p.Name)}, nil
}

func main() {
	// gRPC
	port := 10000
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Run server port: %d", port)
	grpcServer := grpc.NewServer()
	// メソッドを定義
	pb.RegisterHelloGrpcServer(grpcServer, &server{})
	// gRPCサーバを公開
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
