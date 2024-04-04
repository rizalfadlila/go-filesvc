package cmd

import (
	log2 "github.com/gofiber/fiber/v2/log"
	"github.com/rizalfadlila/go-filesvc/cmd/web"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Short: "File uploader",
	}
)

func Execute() {
	registerWebCommand()

	if err := rootCmd.Execute(); err != nil {
		log2.Fatal("error execute command: %w", err)
		os.Exit(-1)
	}
}

func registerWebCommand() {
	rootCmd.AddCommand(web.ServeWEBCmd())
}
