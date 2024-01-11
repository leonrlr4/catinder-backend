# 使用 golang 最新版作為基礎映像
FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/server

RUN go install github.com/cosmtrek/air@latest

# 確定 $GOPATH 是被設置好的，並且 `air` 可以被執行。
RUN echo $GOPATH
RUN air -v

# 使用 Air 啟動應用
CMD ["air"]
