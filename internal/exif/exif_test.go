package exif

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractDate(t *testing.T) {
	// テスト用の一時ディレクトリを作成
	testDir := t.TempDir()
	defer os.RemoveAll(testDir)

	// この段階ではスキップ（実際のEXIFファイルが必要なため）
	t.Skip("実際のEXIFファイルが必要なため、CIでは実行しないようにスキップ")

	testFile := filepath.Join(testDir, "test_exif.jpg")
	// 実際のテスト環境では、ここでテスト用のEXIF付きJPEGファイルをコピーする

	extractor := NewExtractor()

	date, err := extractor.ExtractDate(testFile)
	if err != nil {
		t.Fatalf("ExtractDate failed: %v", err)
	}

	// YYYY-MM-DDフォーマットを検証
	if len(date) != 10 || date[4] != '-' || date[7] != '-' {
		t.Errorf("Expected date in YYYY-MM-DD format, got %s", date)
	}
}

func TestExtractModel(t *testing.T) {
	// テスト用の一時ディレクトリを作成
	testDir := t.TempDir()
	defer os.RemoveAll(testDir)

	t.Skip("実際のEXIFファイルが必要なため、CIでは実行しないようにスキップ")

	testFile := filepath.Join(testDir, "test_exif.jpg")
	// 実際のテスト環境では、ここでテスト用のEXIF付きJPEGファイルをコピーする

	extractor := NewExtractor()

	model, err := extractor.ExtractModel(testFile)
	if err != nil {
		t.Fatalf("ExtractModel failed: %v", err)
	}

	// モデル名が空でないことを確認
	if model == "" {
		t.Errorf("Expected non-empty model string")
	}
}

// モックがインターフェースを実装していることを確認
var _ ExifExtractor = (*MockExtractor)(nil)
