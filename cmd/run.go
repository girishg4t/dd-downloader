package cmd

import (
	"log"

	csv_processor "github.com/dd-downloader/pkg/csv"
	"github.com/spf13/cobra"
)

const (
	configFile = "config-file"
	csvFile    = "file"
)

var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run Datadog downloader",
	Args:  cobra.MinimumNArgs(1),
}

var cmdSynchronously = &cobra.Command{
	Use:   "sync",
	Short: "Run Datadog downloader in synchronous mode",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		yp := readYaml(cmd)
		err := yp.RunSync()
		if err != nil {
			log.Fatalf("Data is not in correct format as per Yaml for query '%s', error %v", yp.Yaml.Spec.DatadogFilter.Query, err)
		}
		log.Printf("Output is available in %s\n", yp.CsvFile)
	},
}

var cmdParallel = &cobra.Command{
	Use:   "parallel",
	Short: "Run Datadog downloader in parallel mode",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		yp := readYaml(cmd)
		done := make(chan bool)
		ch := make(chan [][]string)
		go yp.RunParallel(ch, done)
		for {
			select {
			case elem, ok := <-ch:
				{
					if ok {
						csv_processor.CsvWriter(yp.CsvFile, nil, elem)
					}
				}
			case <-done:
				log.Printf("Output is available in %s\n", yp.CsvFile)
				return
			}
		}

	},
}
