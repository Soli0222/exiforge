package main

import (
	"os"
	"path/filepath"
	"testing"

	"exiforge/internal/exif"
	"exiforge/internal/processor"
)

// Loggerのモック
type mockLogger struct {
	logMessages []string
	fatalCalled bool
}

func (l *mockLogger) Printf(format string, v ...interface{}) {
	l.logMessages = append(l.logMessages, format)
}

func (l *mockLogger) Fatal(v ...interface{}) {
	l.fatalCalled = true
}

func TestMainProcessingLogic(t *testing.T) {
	// テスト用の一時ディレクトリ作成
	testDir := t.TempDir()
	testFile := filepath.Join(testDir, "test.JPG")
	if err := os.WriteFile(testFile, []byte("dummy content"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// モック作成
	extractor := &exif.MockExtractor{
		DateToReturn:  "2025-03-08",
		ModelToReturn: "Test_Camera",
	}

	logger := &mockLogger{}

	// FileProcessorを生成してテスト
	proc := processor.NewFileProcessor(extractor, logger)

	// プロセッサを直接呼び出す
	err := proc.ProcessFiles(testDir, "*.JPG", "JPG", processor.ProcessorOptions{UseModel: false})
	if err != nil {
		t.Fatalf("ProcessFiles failed: %v", err)
	}

	// 処理結果の検証
	expectedDir := filepath.Join(testDir, "JPG", "2025-03-08")
	if _, err := os.Stat(expectedDir); os.IsNotExist(err) {
		t.Errorf("Expected directory %s was not created", expectedDir)
	}
}

// cmd.Execute をモック化して統合テスト
func TestCmdIntegration(t *testing.T) {
	// skipping this test in normal runs as it would require EXIF files
	t.Skip("統合テストは別途実行します")
}
