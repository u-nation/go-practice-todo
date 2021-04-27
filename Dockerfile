# ---------------------------
# 開発環境用イメージ
# ---------------------------
# buster https://future-architect.github.io/articles/20200513/
FROM golang:1.16.0-buster as develop

ARG API_VERSION=v0.0.0
ENV API_VERSION=$API_VERSION \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on

# https://knowledge.sakura.ad.jp/15253/
WORKDIR $GOPATH/src/github.com/u-nation/go-practice-todo

# アプリケーションの依存パッケージインストール last-arg-is-destination
COPY go.mod go.sum ./
RUN go mod download

# [golangci-lint](https://golangci-lint.run/usage/install/#ci-installation)
# [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) https://dev.classmethod.jp/articles/db-migrate-with-golang-migrate/

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.35.2 && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz -C /usr/bin && \
    mv /usr/bin/migrate.linux-amd64 /usr/bin/migrate && \
    go install \
        golang.org/x/tools/cmd/godoc \
        github.com/cosmtrek/air \
        github.com/google/wire/cmd/wire && \
        rm -rf /tmp/*

# ---------------------------
# ビルドステージ
# ---------------------------
FROM develop as build
COPY . .
RUN go build -o /go/bin/app ./cmd/api/main.go

# ---------------------------
# デプロイ用イメージ
# ---------------------------
FROM gcr.io/distroless/base-debian10 as deploy
EXPOSE 5000
COPY --from=build /go/bin/app /
COPY ./migrations /migrations

ENTRYPOINT ["/app"]
