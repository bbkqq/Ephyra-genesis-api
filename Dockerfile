FROM golang:1.24-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0
ENV GO111MODULE=on

#ENV GOPROXY https://goproxy.cn,direct
# 编译项目

WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod tidy
COPY . .
RUN sh ./build.sh
RUN mv ./output /app

FROM alpine:latest
# 安装bash工具
RUN apk update --no-cache && apk add --no-cache bash && apk add --no-cache tzdata
# 设置时区
#RUN ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
#ENV TZ Asia/Shanghai
# 将编译好的二进制放到新的文件夹里
COPY --from=builder /app /app
# 设置环境变量
ENV GO_ENV=prod
WORKDIR /app
# 运行服务
CMD ["bash", "bootstrap.sh"]
