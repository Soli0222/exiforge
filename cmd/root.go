/*
Copyright © 2025 Soli
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var (
	ext   string
	dir   string
	model bool
)

var rootCmd = &cobra.Command{
	Use:   "exiforge",
	Short: "Organize pictures using EXIF data",
	Long:  `Organize picture files by date and optionally by camera model based on EXIF metadata.`,
	Run: func(cmd *cobra.Command, args []string) {
		extensions := strings.Split(ext, ",")
		for _, e := range extensions {
			upperExt := strings.ToUpper(strings.TrimSpace(e))
			pattern := "*." + upperExt
			processFiles(dir, pattern, upperExt, model)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Define cobra flags matching the old command options.
	rootCmd.Flags().StringVarP(&ext, "extensions", "e", "JPG,ARW", "Comma-separated list of file extensions to process (e.g., JPG,DNG,PNG)")
	rootCmd.Flags().StringVarP(&dir, "directory", "d", ".", "Directory to process")
	rootCmd.Flags().BoolVarP(&model, "model", "m", false, "Use camera model from EXIF for file organization")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

	// Determine the output directory.
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
			// Folder structure: ${Model}/${Extension}/${Date}
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

	modelTag, err := exifData.Get(exif.Model)
	if err != nil {
		return "", err
	}

	modelStr := strings.Trim(modelTag.String(), "\"")
	modelStr = strings.ReplaceAll(modelStr, " ", "_")
	return modelStr, nil
}

func moveFile(src, dst string) error {
	return os.Rename(src, dst)
}
