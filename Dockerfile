# 基于官方 Golang 镜像构建
FROM golang:1.23-alpine AS builder

WORKDIR /app

# 复制 go.mod 和 go.sum，下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制源码
COPY . .

# 编译二进制文件
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o music-unlocker ./cmd/music-unlocker

# 使用更小的镜像运行
FROM alpine:latest

WORKDIR /app

# 复制编译好的二进制文件
COPY --from=builder /app/music-unlocker .

# 暴露端口（如有需要）
EXPOSE 8080

# 运行程序
ENTRYPOINT ["./music-unlocker"]
