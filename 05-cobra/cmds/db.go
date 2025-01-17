package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	dbCmd = &cobra.Command{
		Use:   "db",
		Short: "Database operations",
		Long:  `Perform various database operations`,
	}

	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Run database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running database migrations...")
		},
	}

	backupCmd = &cobra.Command{
		Use:   "backup",
		Short: "Backup database",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Backing up database...")
		},
	}
)

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(migrateCmd)
	dbCmd.AddCommand(backupCmd)
}
