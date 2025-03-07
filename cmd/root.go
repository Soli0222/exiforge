/*
Copyright © 2025 Soli
*/
package cmd

import (
	"log"
	"os"
	"strings"

	"exiforge/internal/exif"
	"exiforge/internal/processor"

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
	RunE: func(cmd *cobra.Command, args []string) error {
		extractor := exif.NewExtractor()
		fileProcessor := processor.NewFileProcessor(extractor, log.Default())

		extensions := strings.Split(ext, ",")
		for _, e := range extensions {
			upperExt := strings.ToUpper(strings.TrimSpace(e))
			pattern := "*." + upperExt

			opts := processor.ProcessorOptions{
				UseModel: model,
			}

			if err := fileProcessor.ProcessFiles(dir, pattern, upperExt, opts); err != nil {
				return err
			}
		}

		return nil
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
}
