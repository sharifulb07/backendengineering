package main

import (
	"fmt"
	"os"

	"file-parser/internal/parser"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myparser",
	Short: "MyParser - A powerful file parser CLI",
}

func main() {
	rootCmd.AddCommand(parser.CSVCmd)
	rootCmd.AddCommand(parser.LogCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}