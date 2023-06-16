package cmd

import (
	"log"

	file_processor "github.com/dd-downloader/pkg/csv"
	"github.com/spf13/cobra"
)

const (
	configFileName = "name"
)

var cmdGenerate = &cobra.Command{
	Use:   "generate",
	Short: "Generates the sample Datadog file",
	Args:  cobra.MinimumNArgs(1),
}

var cmdConfig = &cobra.Command{
	Use:   "config",
	Short: "Generates the sample Datadog config file",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		configFileNameFlag, _ := cmd.Flags().GetString(configFileName)
		file_processor.CreateConfigYAML(configFileNameFlag)
		log.Printf("generated sample config file at location %s\n", configFileNameFlag)
	},
}
