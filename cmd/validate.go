package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var cmdValidate = &cobra.Command{
	Use:   "validate",
	Short: "Validate Datadog downloader",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		yp := readYaml(cmd)
		configFileFlag, _ := cmd.Flags().GetString(configFile)
		data, err := yp.Validate(configFileFlag)
		if err != nil {
			log.Fatalf("Data is not in correct format as per Yaml for query '%s', error %v", yp.Yaml.Spec.DatadogFilter.Query, err)
		}
		log.Println("Output will be like this:")
		for _, out := range data {
			fmt.Printf("%s\n", strings.Join(out, ", "))
		}
	},
}
