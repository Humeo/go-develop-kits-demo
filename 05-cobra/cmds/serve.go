package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	port     int
	host     string
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start the server",
		Long:  `Start the server with specified configuration`,
		Run: func(cmd *cobra.Command, args []string) {
			serve()
		},
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)

	// 添加本地标志
	serveCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the server on")
	serveCmd.Flags().StringVarP(&host, "host", "H", "localhost", "Host to run the server on")
}

func serve() {
	fmt.Printf("Starting server on %s:%d\n", host, port)
	// 这里添加服务器启动逻辑
}
