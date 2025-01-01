# Makefile
.PHONY: build-web move-dist build-server run clean

APP_NAME = surveyX

build:
	make build-web
	make build-server

build-arm:
	make build-web
	make build-server-arm

build-web:
	cd ./web && yarn && yarn build
	rm -rf ./server/dist/ && cp -rf ./web/dist/ ./server/dist/


build-server:
	cd ./server/ && go mod tidy && CGO_ENABLED=0 go build -ldflags "-s -w" -o ../bin/$(APP_NAME)

build-server-arm:
	cd ./server/ && go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o ../bin/$(APP_NAME)

