# 前端-缓存
FROM node:20 AS web-cached
WORKDIR /survey-x/web
COPY ./web/package.json ./web/yarn.lock ./
RUN yarn install

# 前端-编译
FROM web-cached AS web-builder
COPY ./web/ ./
RUN yarn build

# 后端-缓存
FROM golang:1.23.4-alpine AS server-cached
WORKDIR /survey-x/server
COPY ./server/go.mod ./server/go.sum ./
RUN go env -w GOPROXY=https://goproxy.io && \
    go mod tidy  && \
    go mod download

# 后端-编译   
FROM server-cached AS server-builder
COPY ./server/ ./
COPY --from=web-builder /survey-x/web/dist/ ./dist
RUN go build -ldflags="-s -w" -o app .

# alpine镜像
FROM alpine:latest
WORKDIR /home
RUN apk --no-cache add bash
RUN apk add --no-cache tzdata && \
  cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
  apk del tzdata && \
  echo "Asia/Shanghai" > /etc/timezone
COPY --from=server-builder /survey-x/server/app ./


EXPOSE 65534

# 运行应用程序
CMD ["./app"]