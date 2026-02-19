# gRPC Hello World 示例

这是第一个 gRPC 学习项目，演示了基本的 gRPC 服务端和客户端实现。

## 项目结构

```
hello-world/
├── hello/
│   └── hello.proto          # Protocol Buffers 定义文件
├── server/
│   └── main.go              # 服务端实现
├── client/
│   └── main.go              # 客户端实现
├── go.mod                   # Go 模块文件
└── README.md                # 本文件
```

## 使用步骤

### 1. 生成 Go 代码

在 `hello-world` 目录下运行：

```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       hello/hello.proto
```

这会在 `hello/` 目录下生成：
- `hello.pb.go` - 消息类型的 Go 代码
- `hello_grpc.pb.go` - 服务接口的 Go 代码

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 运行服务端

在一个终端窗口中运行：

```bash
go run server/main.go
```

你应该看到：
```
gRPC 服务器启动在 :50051
```

### 4. 运行客户端

在另一个终端窗口中运行：

```bash
go run client/main.go
```

你应该看到：
```
服务器响应: Hello, World!
```
哦，原来你说的是gRPC呀。在gRPC里，message确实是用protobuf定义数据结构的，就像定义请求和响应的格式，然后service里声明远程调用的方法，指定输入输出是哪个message。你是刚开始用protobuf写gRPC的服务定义吗？
## 代码说明

### .proto 文件

- `syntax = "proto3"` - 使用 Protocol Buffers 版本 3
- `service Greeter` - 定义 gRPC 服务
- `rpc SayHello` - 定义 RPC 方法
- `message HelloRequest/HelloReply` - 定义请求和响应消息

### 服务端

- 实现 `GreeterServer` 接口
- 监听 `:50051` 端口
- 处理客户端请求并返回响应

### 客户端

- 连接到 `localhost:50051`
- 创建客户端存根
- 调用 `SayHello` 方法
- 打印服务器响应

## 下一步

- 尝试修改消息内容
- 添加更多 RPC 方法
- 学习流式传输（Streaming）
- 添加错误处理

