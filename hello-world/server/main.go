package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc-hello-world/hello/hello"
)

// server 实现 Greeter 服务
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello 实现 Greeter 服务的 SayHello 方法
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("收到请求: %v", req.GetName())
	return &pb.HelloReply{
		Message: "Hello, " + req.GetName() + "!",
	}, nil
}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	// 创建 gRPC 服务器
	s := grpc.NewServer()

	// 注册服务
	pb.RegisterGreeterServer(s, &server{})

	log.Println("gRPC 服务器启动在 :50051")
	
	// 启动服务器
	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}

