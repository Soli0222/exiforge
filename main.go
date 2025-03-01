package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/schollz/progressbar/v3"
)

func main() {
	extFlag := flag.String("ext", "JPG,DNG", "Comma-separated list of file extensions to process (e.g., JPG,DNG,PNG)")
	dirFlag := flag.String("dir", ".", "Directory to process (default is current directory)")
	modelFlag := flag.Bool("model", false, "Use camera model from exif for file organization")
	flag.Parse()

	extensions := strings.Split(*extFlag, ",")
	for _, ext := range extensions {
		processFiles(*dirFlag, "*."+strings.ToUpper(ext), strings.ToUpper(ext), *modelFlag)
	}
}

func processFiles(baseDir, pattern, ext string, useModel bool) {
	files, err := filepath.Glob(filepath.Join(baseDir, pattern))
	if err != nil {
		log.Fatal(err)
	}

	if len(files) == 0 {
		log.Printf("No files found for pattern %s in %s", pattern, baseDir)
		return
	}

	// 出力先の基点ディレクトリは、modelフラグが有効な場合はbaseDirそのものとする
	var outputDir string
	if useModel {
		outputDir = baseDir
	} else {
		outputDir = filepath.Join(baseDir, ext)
		if err := os.Mkdir(outputDir, 0755); err != nil && !os.IsExist(err) {
			log.Fatal(err)
		}
	}

	bar := progressbar.NewOptions(
		len(files),
		progressbar.OptionSetDescription("Processing "+ext),
	)

	for _, file := range files {
		var subDir string
		if useModel {
			// ${機種名}/${拡張子}/${日付} の構造とする
			modelStr, err := extractModel(file)
			if err != nil {
				log.Printf("Error extracting model from %s: %v", file, err)
				bar.Add(1)
				continue
			}

			dateStr, err := extractDate(file)
			if err != nil {
				log.Printf("Error extracting date from %s: %v", file, err)
				bar.Add(1)
				continue
			}

			subDir = filepath.Join(modelStr, ext, dateStr)
		} else {
			// modelフラグが無い場合は従来通り、拡張子ディレクトリ直下に日付フォルダを作成するなど、必要に応じて実装してください
			subDir, err = extractDate(file)
			if err != nil {
				log.Printf("Error processing %s: %v", file, err)
				bar.Add(1)
				continue
			}
		}

		destDir := filepath.Join(outputDir, subDir)
		if err := os.MkdirAll(destDir, 0755); err != nil {
			log.Fatal(err)
		}

		destPath := filepath.Join(destDir, filepath.Base(file))
		if err := moveFile(file, destPath); err != nil {
			log.Printf("Failed to move %s: %v", file, err)
		}

		bar.Add(1)
	}

	println()
}

func extractDate(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	exifData, err := exif.Decode(file)
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

func extractModel(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	exifData, err := exif.Decode(file)
	if err != nil {
		return "", err
	}

	model, err := exifData.Get(exif.Model)
	if err != nil {
		return "", err
	}

	// 機種名から余分な引用符を除去し、スペースをアンダースコアへ置換して整形する
	modelStr := strings.Trim(model.String(), "\"")
	modelStr = strings.ReplaceAll(modelStr, " ", "_")
	return modelStr, nil
}

func moveFile(src, dst string) error {
	return os.Rename(src, dst)
}
