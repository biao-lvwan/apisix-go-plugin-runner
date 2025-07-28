# --- Build Stage ---
# 使用官方的 Golang 1.23 Alpine 镜像作为构建阶段的基础镜像，以获得更小的镜像。
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app

RUN apk add --no-cache make git


ENTRYPOINT ["tail", "-f", "/dev/null"]