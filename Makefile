# アプリケーション名
APP_NAME := Go-PasswordGen
VERSION := 1.0.0

# デフォルトのターゲット
all: build

# アプリケーションビルド
build:
	go build -o ./bin/$(APP_NAME).exe .

# アプリケーションの実行
run:
	go run .

# テストの実行
test:
	go test ./test

# 依存関係の更新
deps:
	go mod tidy

# ヘルプ情報の表示
help:
	@echo "make: ビルドタスクを実行します"
	@echo "make run: アプリケーションを実行します"
	@echo "make test: テストを実行します"
	@echo "make deps: 依存関係を更新します"

