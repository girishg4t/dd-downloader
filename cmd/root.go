package cmd

import (
	"log"
	"os"
	"path"

	"github.com/dd-downloader/pkg/model"
	"github.com/dd-downloader/pkg/processor"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var rootCmd = &cobra.Command{
	Use:   "dd-downloader",
	Short: "Datadog downloader to download logs",
	Long:  `Datadog downloader uses template to download logs in csv format`,
}

func Execute() {
	rootCmd.AddCommand(cmdRun, cmdValidate, cmdGenerate)
	cmdSynchronously.Flags().StringP("config-file", "c", "", "config file path")
	cmdSynchronously.Flags().StringP("file", "f", "", "file path to save data")

	cmdParallel.Flags().StringP("config-file", "c", "", "config file path")
	cmdParallel.Flags().StringP("file", "f", "", "file path to save data")

	cmdValidate.Flags().StringP("config-file", "c", "", "config file path")
	cmdRun.AddCommand(cmdSynchronously, cmdParallel)

	cmdGenerate.AddCommand(cmdConfig)
	cmdConfig.Flags().StringP("name", "n", "", "name of the file which need to generated")

	err := cmdValidate.MarkFlagRequired("config-file")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = cmdSynchronously.MarkFlagRequired("config-file")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = cmdSynchronously.MarkFlagRequired("file")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = cmdParallel.MarkFlagRequired("config-file")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = cmdParallel.MarkFlagRequired("file")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	cobra.CheckErr(rootCmd.Execute())
}

// reads the given yaml file for processing
func readYaml(cmd *cobra.Command) processor.YamlProcessor {
	dir, _ := os.Getwd()
	configFileFlag, _ := cmd.Flags().GetString(configFile)
	yamlData, err := os.ReadFile(path.Join(dir, configFileFlag))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	y := model.YamlMapping{}

	err = yaml.Unmarshal(yamlData, &y)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	csvFileFlag, _ := cmd.Flags().GetString(csvFile)
	return processor.NewYamlProcessor(&y, csvFileFlag)
}
