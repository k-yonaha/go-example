FROM golang:1.23-alpine

# 必要なツールをインストール
RUN apk update && apk add --no-cache git curl
RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY ./src .

# go.modとgo.sumの生成
RUN go mod init github.com/k-yonaha/go-example

# 整頓
RUN go mod tidy

# コンパイル後のバイナリ実行用ポート
EXPOSE 8080

# アプリケーションを起動
CMD ["air"]