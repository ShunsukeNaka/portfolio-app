# Step 1: Modules caching
# コンパイルに必要なモジュールをインストールする
FROM golang:1.23.4-alpine3.21 AS modules
WORKDIR /src
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/ \
    go mod download -x
COPY . .

# Step 2: Server building
# ソースコードをコンパイルし、/bin/serverにバイナリーを配置する
FROM modules AS build-server
RUN --mount=type=cache,target=/go/pkg/mod/ \
    go build -o /bin/server ./cmd/app

COPY --from=build-server /bin/server /bin/server

EXPOSE 8080

ENTRYPOINT [ "/bin/server" ]
