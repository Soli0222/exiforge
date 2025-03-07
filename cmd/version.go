/*
Copyright © 2025 Soli
*/
package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// Version information
var (
	Version = "v2.1.0"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version of exiforge",
	Long:  "This command displays the current version of the exiforge CLI tool.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("exiforge %s (%s/%s)\n", Version, runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
