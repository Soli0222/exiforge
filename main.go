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
	flag.Parse()

	extensions := strings.Split(*extFlag, ",")
	for _, ext := range extensions {
		processFiles(*dirFlag, "*."+strings.ToUpper(ext), strings.ToUpper(ext))
	}
}

func processFiles(baseDir, pattern, dir string) {
	files, err := filepath.Glob(filepath.Join(baseDir, pattern))
	if err != nil {
		log.Fatal(err)
	}

	if len(files) == 0 {
		log.Printf("No files found for pattern %s in %s", pattern, baseDir)
		return
	}

	outputDir := filepath.Join(baseDir, dir)
	if err := os.Mkdir(outputDir, 0755); err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	bar := progressbar.NewOptions(
		len(files),
		progressbar.OptionSetDescription("Processing "+dir),
	)

	for _, file := range files {
		date, err := extractDate(file)
		if err != nil {
			log.Printf("Error processing %s: %v", file, err)
			bar.Add(1)
			continue
		}

		destDir := filepath.Join(outputDir, date)
		if err := os.Mkdir(destDir, 0755); err != nil && !os.IsExist(err) {
			log.Fatal(err)
		}

		if err := moveFile(file, filepath.Join(destDir, filepath.Base(file))); err != nil {
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

func moveFile(src, dst string) error {
	return os.Rename(src, dst)
}
