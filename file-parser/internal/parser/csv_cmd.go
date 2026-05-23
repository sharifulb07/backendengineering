package parser

import (
	"encoding/json"
	"fmt"
	"os"
	// "file-parser/internal/models"
	"file-parser/internal/utils"
	"file-parser/pkg/formats"
	"github.com/spf13/cobra"
)

var CSVCmd = &cobra.Command{
	Use:   "csv",
	Short: "Parse CSV files",
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		toJSON, _ := cmd.Flags().GetString("to-json")

		if !utils.FileExists(file) {
			fmt.Printf("Error: File %s not found\n", file)
			return
		}

		result, err := formats.ParseCSV(file)
		if err != nil {
			fmt.Printf("Error parsing CSV: %v\n", err)
			return
		}

		utils.PrintStats(result)

		if toJSON != "" {
			data, _ := json.MarshalIndent(result, "", "  ")
			os.WriteFile(toJSON, data, 0644)
			fmt.Printf("✅ JSON exported to: %s\n", toJSON)
		}
	},
}

func init() {
	CSVCmd.Flags().StringP("file", "f", "", "CSV file path")
	CSVCmd.Flags().String("to-json", "", "Export to JSON file")
	CSVCmd.MarkFlagRequired("file")
}