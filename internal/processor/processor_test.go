package processor

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"exiforge/internal/exif"
)

// Loggerのモック
type MockLogger struct {
	LogMessages []string
	FatalCalled bool
}

func (l *MockLogger) Printf(format string, v ...interface{}) {
	l.LogMessages = append(l.LogMessages, format)
}

func (l *MockLogger) Fatal(v ...interface{}) {
	l.FatalCalled = true
}

func TestProcessFiles_NoFiles(t *testing.T) {
	tempDir := t.TempDir()

	mockExtractor := &exif.MockExtractor{}
	mockLogger := &MockLogger{}

	processor := NewFileProcessor(mockExtractor, mockLogger)

	// 存在しないパターンを指定して、ファイルが見つからないケースをテスト
	err := processor.ProcessFiles(tempDir, "*.NONEXISTENT", "NONEXISTENT", ProcessorOptions{UseModel: false})

	if err != nil {
		t.Fatalf("ProcessFiles returned error: %v", err)
	}

	// ログメッセージの確認
	if len(mockLogger.LogMessages) != 1 {
		t.Errorf("Expected 1 log message, got %d", len(mockLogger.LogMessages))
	}
}

func TestProcessFiles_WithFiles(t *testing.T) {
	tempDir := t.TempDir()

	// テストファイルを作成
	testFiles := []string{"test1.JPG", "test2.JPG", "test3.JPG"}
	for _, file := range testFiles {
		filePath := filepath.Join(tempDir, file)
		if err := os.WriteFile(filePath, []byte("dummy content"), 0644); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}

	mockExtractor := &exif.MockExtractor{
		DateToReturn:  "2025-03-08",
		ModelToReturn: "Test_Camera",
	}

	mockLogger := &MockLogger{}

	processor := NewFileProcessor(mockExtractor, mockLogger)

	// モデルなしでファイル処理
	err := processor.ProcessFiles(tempDir, "*.JPG", "JPG", ProcessorOptions{UseModel: false})

	if err != nil {
		t.Fatalf("ProcessFiles returned error: %v", err)
	}

	// 期待される移動先ディレクトリを確認
	expectedDir := filepath.Join(tempDir, "JPG", "2025-03-08")
	if _, err := os.Stat(expectedDir); os.IsNotExist(err) {
		t.Errorf("Expected directory %s was not created", expectedDir)
	}
}

func TestProcessFile_WithModel(t *testing.T) {
	tempDir := t.TempDir()

	// テストファイルを作成
	testFile := filepath.Join(tempDir, "test.JPG")
	if err := os.WriteFile(testFile, []byte("dummy content"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	mockExtractor := &exif.MockExtractor{
		DateToReturn:  "2025-03-08",
		ModelToReturn: "Test_Camera",
	}

	mockLogger := &MockLogger{}

	processor := NewFileProcessor(mockExtractor, mockLogger)

	// モデルを使用してファイル処理
	err := processor.processFile(testFile, tempDir, "JPG", true)

	if err != nil {
		t.Fatalf("processFile returned error: %v", err)
	}

	// 期待される移動先ディレクトリを確認
	expectedDir := filepath.Join(tempDir, "Test_Camera", "JPG", "2025-03-08")
	expectedFile := filepath.Join(expectedDir, "test.JPG")
	if _, err := os.Stat(expectedFile); os.IsNotExist(err) {
		t.Errorf("Expected file %s was not created", expectedFile)
	}
}

func TestProcessFile_WithError(t *testing.T) {
	tempDir := t.TempDir()

	// テストファイルを作成
	testFile := filepath.Join(tempDir, "test.JPG")
	if err := os.WriteFile(testFile, []byte("dummy content"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	mockExtractor := &exif.MockExtractor{
		ErrorToReturn: errors.New("simulated EXIF error"),
	}

	mockLogger := &MockLogger{}

	processor := NewFileProcessor(mockExtractor, mockLogger)

	// エラーケースのテスト
	err := processor.processFile(testFile, tempDir, "JPG", false)

	if err == nil {
		t.Fatalf("Expected error but got nil")
	}
}
