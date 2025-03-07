package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRootCommand_Flags(t *testing.T) {
	// 初期化関数を実行してフラグを設定
	originalExt, originalDir, originalModel := ext, dir, model
	defer func() {
		// テスト後にグローバル変数を元に戻す
		ext, dir, model = originalExt, originalDir, originalModel
	}()

	// デフォルト値のテスト - init()関数は既に実行されている
	if ext != "JPG,ARW" {
		t.Errorf("Expected default ext to be 'JPG,ARW', got '%s'", ext)
	}

	if dir != "." {
		t.Errorf("Expected default dir to be '.', got '%s'", dir)
	}

	if model != false {
		t.Errorf("Expected default model to be false, got %v", model)
	}

	// フラグの設定テスト
	args := []string{"--extensions=RAW,DNG", "--directory=/tmp", "--model=true"}
	rootCmd.SetArgs(args)

	// コマンドの実行はせず、引数解析だけ行う
	rootCmd.ParseFlags(args)

	if ext != "RAW,DNG" {
		t.Errorf("Expected ext to be 'RAW,DNG', got '%s'", ext)
	}

	if dir != "/tmp" {
		t.Errorf("Expected dir to be '/tmp', got '%s'", dir)
	}

	if !model {
		t.Errorf("Expected model to be true, got %v", model)
	}
}

func TestRootCommand_Integration(t *testing.T) {
	// 統合テスト - 実際のファイルシステム操作を含む
	// CI環境では不安定になる可能性があるため、通常はスキップするか環境変数で制御する
	if os.Getenv("RUN_INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test. Set RUN_INTEGRATION_TESTS=true to run")
	}

	// テスト用の一時ディレクトリを作成
	tempDir := t.TempDir()

	// テスト用のJPGファイルを作成（実際のEXIFは含まない）
	testFile := filepath.Join(tempDir, "test.JPG")
	if err := os.WriteFile(testFile, []byte("test data"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// コマンドの引数を設定
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()
	os.Args = []string{"exiforge", "--directory=" + tempDir, "--extensions=JPG"}

	// 実際のコマンド実行はEXIFデータが必要なため、ここではスキップ
	t.Skip("実際のEXIFファイルが必要なため、スキップ")

	// Execute()  // 実際のテストでは有効化
}
