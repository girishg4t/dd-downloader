package util

import (
	"fmt"

	"github.com/girishg4t/dd-downloader/pkg/model"
)

func ReadHeader(m []model.InnerFieldMapping, headers *[]string) {
	for _, val := range m {
		fmt.Println("Value: ", val)
		if val.Field == "-" && val.InnerField != nil || len(val.InnerField) > 0 {
			ReadHeader(val.InnerField, headers)
			continue
		}
		*headers = append(*headers, val.Field)
	}
}
