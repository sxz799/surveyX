# Makefile
.PHONY: build-web move-dist build-server run clean

APP_NAME = surveyX

build:
	make build-web
	make move-dist
	make build-server

build-arm:
	make build-web
	make move-dist
	make build-server-arm

build-web:
ifeq ($(shell uname -s), Darwin)
	cd ./web && sed -i '' 's#//sed_tag/base#base#g' vite.config.js && yarn && yarn build && sed -i '' 's#base#//sed_tag/base#g' vite.config.js
else
	cd ./web && sed -i 's#//sed_tag/base#base#g' vite.config.js && yarn && yarn build && sed -i 's#base#//sed_tag/base#g' vite.config.js
endif
move-dist:
	rm -rf ./server/dist/ && mv ./web/dist/ ./server/

build-server:
	cd ./server/ && go mod tidy && CGO_ENABLED=0 go build -ldflags "-s -w" -o $(APP_NAME)

build-server-arm:
	cd ./server/ && go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o $(APP_NAME)