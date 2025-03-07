package exif

import (
	"os"
	"strings"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

// Extractor handles EXIF metadata extraction
type Extractor struct{}

// NewExtractor creates a new EXIF extractor
func NewExtractor() *Extractor {
	return &Extractor{}
}

// ExtractDate returns date from EXIF in YYYY-MM-DD format
func (e *Extractor) ExtractDate(filename string) (string, error) {
	exifData, err := e.getExifData(filename)
	if err != nil {
		return "", err
	}

	dateTime, err := exifData.Get(exif.DateTimeOriginal)
	if err != nil {
		return "", err
	}

	dateStr := strings.Trim(dateTime.String(), "\"")
	parsedTime, err := time.Parse("2006:01:02 15:04:05", dateStr)
	if err != nil {
		return "", err
	}

	return parsedTime.Format("2006-01-02"), nil
}

// ExtractModel returns camera model from EXIF
func (e *Extractor) ExtractModel(filename string) (string, error) {
	exifData, err := e.getExifData(filename)
	if err != nil {
		return "", err
	}

	modelTag, err := exifData.Get(exif.Model)
	if err != nil {
		return "", err
	}

	modelStr := strings.Trim(modelTag.String(), "\"")
	modelStr = strings.ReplaceAll(modelStr, " ", "_")
	return modelStr, nil
}

// getExifData opens file and extracts EXIF data
func (e *Extractor) getExifData(filename string) (*exif.Exif, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return exif.Decode(file)
}
