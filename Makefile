.PHONY: test test-coverage clean build

# テスト実行
test:
	go test ./...

# カバレッジレポート付きテスト
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# カバレッジ結果の表示
show-coverage: test-coverage
	open coverage.html

# ビルド
build:
	go build -o bin/exiforge

# クリーンアップ
clean:
	rm -rf bin/* coverage.out coverage.html

# 統合テスト（実際のファイルシステム操作を含むテスト）
integration-test:
	RUN_INTEGRATION_TESTS=true go test -v -tags=integration ./...

# すべてのターゲットを実行
all: clean build test show-coverage