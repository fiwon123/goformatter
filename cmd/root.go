/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	formatter "github.com/fiwon123/goformatter/internal/goformatter"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goformatter",
	Short: "goformatter is a CLI tool to automate multiples files.",
	Long:  `goformatter is a CLI tool to automate multiples files.`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()

		filepath, _ := flags.GetString("file")
		fileType, _ := flags.GetString("type")
		dirOutput, _ := flags.GetString("output")
		// disableLog, _ := flags.GetBool("disable-logs")
		check, _ := flags.GetBool("check")
		dryRun, _ := flags.GetBool("dry-run")
		inPlace, _ := flags.GetBool("in-place")
		pretty, _ := flags.GetBool("pretty")

		// set_logger(FormatterLogger(disable_log, pretty))

		formatter.Process(filepath, dirOutput, check, dryRun, pretty, inPlace, fileType)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	flags := rootCmd.Flags()

	flags.StringP("file", "f", "", "Path to the file to format. (YAML, JSON, ...)")
	rootCmd.MarkFlagRequired("file")
	flags.StringP("type", "t", "", "Explicitly set file type in lowercase (example: 'json', 'yaml', ...).")
	flags.StringP("output", "o", "", "Change default output path.")
	flags.BoolP("disable-logs", "d", false, "Disable all logs while formatting the file.")
	flags.BoolP("check", "c", false, "Check if current file is already formatted.")
	flags.BoolP("dry-run", "r", false, "Show a preview in the CLI without generate an output file.")
	flags.BoolP("in-place", "i", false, "Overwrite the current file with the output.")
	flags.BoolP("pretty", "p", false, "CLI becomes more modern with colours, text boxes, previews becomes more readable and more.")
}
