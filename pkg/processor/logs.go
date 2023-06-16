package processor

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/dd-downloader/pkg/model"
)

//TODO: need to find simple approach for below logic

// check each logs for inner fields and get all the values in the pointer
// since there can be nested object we need to use recursion
func getAllValues(data []interface{}, innerField []model.InnerFieldMapping, ddVal *[]string) {
	for _, fl := range innerField {
		for _, log := range data {
			if fl.InnerField == nil || len(fl.InnerField) == 0 {
				val := getValue(fl, log)
				*ddVal = append(*ddVal, fmt.Sprint(val))
			} else {
				var innerObj []interface{}
				_ = json.Unmarshal([]byte(log.(map[string]interface{})[fl.DdField].(string)), &innerObj)
				getAllValues(innerObj, fl.InnerField, ddVal)
			}
		}
	}
}

// get the value inside the nested struct
func getValue(fl model.InnerFieldMapping, log interface{}) string {
	field_dep := strings.Split(fl.DdField, ".")
	n := len(field_dep)
	val, ok := log.(map[string]interface{})
	if !ok {
		return ""
	}
	var outerObj = deepSearch(val, field_dep)
	return getValueBasedOnType(field_dep[n-1], outerObj)
}

// Inner log object can be json string or a map
// check each key and find type to get map
func deepSearch(val map[string]interface{}, keys []string) map[string]interface{} {
	if len(keys) == 1 {
		return val
	}
	var out map[string]interface{} = val
	for i := 0; i < len(keys)-1; i++ {
		var innerObj map[string]interface{}
		if out[keys[i]] != nil {
			switch reflect.TypeOf(out[keys[i]]).Kind() {
			case reflect.String:
				err := json.Unmarshal([]byte(out[keys[i]].(string)), &innerObj)
				if err != nil {
					continue
				}
				out = innerObj
				continue
			case reflect.Map:
				out = out[keys[i]].(map[string]interface{})
			}
		}

	}
	return out
}

// Since csv uses string to write, we convert log value based on type
func getValueBasedOnType(key string, obj map[string]interface{}) string {
	if obj[key] == nil {
		return ""
	}
	switch reflect.TypeOf(obj[key]).Kind() {
	case reflect.Float64:
		return fmt.Sprintf("%f", obj[key].(float64))
	case reflect.Int64:
		return fmt.Sprintf("%d", obj[key].(int64))
	default:
		return fmt.Sprintf("%s", obj[key])
	}
}
