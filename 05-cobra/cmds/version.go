package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Version   = "1.0.0"
	BuildTime = "2025-01-17"
	GitCommit = "abc123"

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Long:  `All software has versions. This is mycli's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s\n", Version)
			fmt.Printf("Build Time: %s\n", BuildTime)
			fmt.Printf("Git Commit: %s\n", GitCommit)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
