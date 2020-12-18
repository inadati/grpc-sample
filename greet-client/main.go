package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    // .protoから生成されたコードをimportしている
    // 今回は筆者の$GOPATH内に作成したので適宜プロジェクトを作成したパスに合わせる
    pb "github.com/inadati/grpc-sample/protobuf/hellogrpc"
)

const (
    address = "localhost:10000"
)

func main() {
    // gRPCサーバへのconnectionを作成
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    // connection終了処理
    // 関数終了後に実行される
    defer conn.Close()

    c := pb.NewHelloGrpcClient(conn)

    name := "ほげ太郎"

    // タイムアウトを設定
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    // .protoで定義したサーバのGreetServerを呼び出している
    r, err := c.GreetServer(ctx, &pb.GreetRequest{Name: name})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Response: %s", r.Msg)
}
