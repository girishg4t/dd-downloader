package file_processor

import (
	"encoding/csv"
	"log"
	"os"
	"path"
	"time"

	"github.com/girishg4t/dd-downloader/pkg/model"
	"gopkg.in/yaml.v2"
)

func CsvWriter(filename string, header []string, values [][]string) {
	csvFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)

	if header != nil {
		e := csvwriter.Write(header)
		if e != nil {
			log.Fatalf("failed to write file: %s", e)
		}
	}

	for _, val := range values {
		_ = csvwriter.Write(val)
	}
	csvwriter.Flush()
}

func createFile(filePath string, content string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	_, err = file.WriteString("---\n" + content)
	if err != nil {
		log.Fatalf("Failed to write yaml file %s", err)
	}
	defer file.Close()
}

func CreateConfigYAML(name string) {
	if name == "" {
		name = "config.yaml"
	}
	dir, _ := os.Getwd()
	config := model.YamlMapping{}
	config.APIVersion = "datadog/v1"
	config.Kind = "DataDog"
	config.Spec.Auth.DdSite = "datadoghq.com"

	config.Spec.DatadogFilter.Mode = "parallel"
	config.Spec.DatadogFilter.From = int(time.Now().Add(-10 * time.Minute).UnixMilli())
	config.Spec.DatadogFilter.To = int(time.Now().UnixMilli())

	config.Spec.Mapping = []model.InnerFieldMapping{
		{
			Field:   "date",
			DdField: "date",
		},
		{
			Field:   "session_id",
			DdField: "SessionID",
		},
		{
			Field:   "req_id",
			DdField: "event.requestID",
		},
		{
			Field:   "src_id",
			DdField: "event.payload.srcID",
		},
		{
			Field:   "-",
			DdField: "data",
			InnerField: []model.InnerFieldMapping{
				{
					Field:   "event_id",
					DdField: "source.eventID",
				},
			},
		},
	}

	configYAML, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	createFile(path.Join(dir, name), string(configYAML))
}
