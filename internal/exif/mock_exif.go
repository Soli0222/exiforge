package exif

// MockExtractor はExifExtractorインターフェースのモック実装です
// テスト用途に使用されます
type MockExtractor struct {
	DateToReturn  string
	ModelToReturn string
	ErrorToReturn error
}

// モックがインターフェースを実装していることを確認
var _ ExifExtractor = (*MockExtractor)(nil)

// ExtractDate はEXIF日付の抽出をモック化します
func (m *MockExtractor) ExtractDate(filename string) (string, error) {
	if m.ErrorToReturn != nil {
		return "", m.ErrorToReturn
	}
	return m.DateToReturn, nil
}

// ExtractModel はEXIFカメラモデルの抽出をモック化します
func (m *MockExtractor) ExtractModel(filename string) (string, error) {
	if m.ErrorToReturn != nil {
		return "", m.ErrorToReturn
	}
	return m.ModelToReturn, nil
}
