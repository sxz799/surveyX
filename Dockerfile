FROM node:20

WORKDIR /go/src/github.com/sxz799/surveyX/web
COPY ./web/ .

RUN sed -i 's#//sed_tag/base#base#g' vite.config.js

RUN yarn build



# 使用官方 Golang 镜像作为基础镜像
FROM golang:1.21-alpine as builder

# 设置工作目录
WORKDIR /go/src/github.com/sxz799/surveyX/server

# 将应用的代码复制到容器中
COPY ./server/ .

COPY --from=0 /go/src/github.com/sxz799/surveyX/web/dist/ ./dist

# 编译应用程序
RUN go env -w GO111MODULE=on \
    && go env \
    && go mod tidy \
    && go build -ldflags="-s -w" -o app .



FROM alpine:latest

WORKDIR /home

RUN apk --no-cache add bash

RUN apk update && apk add tzdata 
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime 
RUN echo "Asia/Shanghai" > /etc/timezone

COPY --from=1 /go/src/github.com/sxz799/surveyX/server/app ./


EXPOSE 3000

# 运行应用程序
CMD ["./app -port 3000"]