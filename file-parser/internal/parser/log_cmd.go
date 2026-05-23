package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"file-parser/internal/models"
	"file-parser/internal/utils"
	"file-parser/pkg/formats"
	"github.com/spf13/cobra"
)

var LogCmd = &cobra.Command{
	Use:   "log",
	Short: "Parse log files",
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		format, _ := cmd.Flags().GetString("format")
		toJSON, _ := cmd.Flags().GetString("to-json")

		if !utils.FileExists(file) {
			fmt.Printf("Error: File %s not found\n", file)
			return
		}

		var result *models.ParseResult
		var err error

		switch format {
		case "nginx":
			result, err = formats.ParseNginxLog(file)
		default:
			fmt.Println("Unsupported log format. Use: nginx")
			return
		}

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		utils.PrintStats(result)

		if toJSON != "" {
			data, _ := json.MarshalIndent(result, "", "  ")
			os.WriteFile(toJSON, data, 0644)
			fmt.Printf("✅ Exported to: %s\n", toJSON)
		}
	},
}

func init() {
	LogCmd.Flags().StringP("file", "f", "", "Log file path")
	LogCmd.Flags().String("format", "nginx", "Log format (nginx)")
	LogCmd.Flags().String("to-json", "", "Export to JSON")
	LogCmd.MarkFlagRequired("file")
}