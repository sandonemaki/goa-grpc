package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	greeter "grpcgreeter"
	gengreeter "grpcgreeter/gen/greeter"
	genpb "grpcgreeter/gen/grpc/greeter/pb"
	genserver "grpcgreeter/gen/grpc/greeter/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// TCPリスナーを作成
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("リッスンに失敗しました: %v", err)
	}

	// オプション付きで新しいgRPCサーバーを作成
	srv := grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor),
	)

	// サービスを初期化
	svc := greeter.NewGreeterService()

	// エンドポイントを作成
	endpoints := gengreeter.NewEndpoints(svc)

	// gRPCサーバーにサービスを登録
	genpb.RegisterGreeterServer(srv, genserver.New(endpoints, nil))

	// デバッグツール用にサーバーリフレクションを有効化
	reflection.Register(srv)

	// グレースフルシャットダウンを処理
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		log.Println("gRPCサーバーをシャットダウンしています...")
		srv.GracefulStop()
	}()

	// サービスを開始
	log.Printf("gRPCサーバーが:8090でリッスンしています")
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("サービスの提供に失敗しました: %v", err)
	}
}

// ロギングインターセプターの例
func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("%sを処理中", info.FullMethod)
	return handler(ctx, req)
}
