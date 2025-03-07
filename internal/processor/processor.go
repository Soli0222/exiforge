package processor

import (
	"fmt"
	"os"
	"path/filepath"

	"exiforge/internal/exif"

	"github.com/schollz/progressbar/v3"
)

// Logger interface for dependency injection
type Logger interface {
	Printf(format string, v ...interface{})
	Fatal(v ...interface{})
}

// ProcessorOptions defines configuration for the processor
type ProcessorOptions struct {
	UseModel bool
}

// FileProcessor handles file organization based on EXIF
type FileProcessor struct {
	exifExtractor exif.ExifExtractor
	logger        Logger
}

// NewFileProcessor creates a new file processor
func NewFileProcessor(extractor exif.ExifExtractor, logger Logger) *FileProcessor {
	return &FileProcessor{
		exifExtractor: extractor,
		logger:        logger,
	}
}

// ProcessFiles processes files matching the pattern
func (p *FileProcessor) ProcessFiles(baseDir, pattern, ext string, opts ProcessorOptions) error {
	files, err := filepath.Glob(filepath.Join(baseDir, pattern))
	if err != nil {
		return err
	}

	if len(files) == 0 {
		p.logger.Printf("No files found for pattern %s in %s", pattern, baseDir)
		return nil
	}

	// Determine output directory
	var outputDir string
	if opts.UseModel {
		outputDir = baseDir
	} else {
		outputDir = filepath.Join(baseDir, ext)
		if err := os.MkdirAll(outputDir, 0755); err != nil && !os.IsExist(err) {
			return err
		}
	}

	bar := progressbar.NewOptions(
		len(files),
		progressbar.OptionSetDescription(fmt.Sprintf("Processing %s", ext)),
	)

	for _, file := range files {
		if err := p.processFile(file, outputDir, ext, opts.UseModel); err != nil {
			p.logger.Printf("Error processing %s: %v", file, err)
		}
		bar.Add(1)
	}

	fmt.Println() // Add newline after progress bar
	return nil
}

// processFile handles a single file
func (p *FileProcessor) processFile(file, outputDir, ext string, useModel bool) error {
	var subDir string

	if useModel {
		// Folder structure: ${Model}/${Extension}/${Date}
		modelStr, err := p.exifExtractor.ExtractModel(file)
		if err != nil {
			return fmt.Errorf("extract model: %w", err)
		}

		dateStr, err := p.exifExtractor.ExtractDate(file)
		if err != nil {
			return fmt.Errorf("extract date: %w", err)
		}

		subDir = filepath.Join(modelStr, ext, dateStr)
	} else {
		// Folder structure: ${Extension}/${Date}
		dateStr, err := p.exifExtractor.ExtractDate(file)
		if err != nil {
			return fmt.Errorf("extract date: %w", err)
		}

		subDir = dateStr
	}

	destDir := filepath.Join(outputDir, subDir)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return fmt.Errorf("create directory: %w", err)
	}

	destPath := filepath.Join(destDir, filepath.Base(file))
	if err := os.Rename(file, destPath); err != nil {
		return fmt.Errorf("move file: %w", err)
	}

	return nil
}
