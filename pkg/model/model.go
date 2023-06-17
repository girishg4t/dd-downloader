package model

type Envelope map[string]interface{}

type Auth struct {
	DdSite   string `yaml:"dd_site"`
	DdAPIKey string `yaml:"dd_api_key"`
	DdAppKey string `yaml:"dd_app_key"`
}

type DataDogFilter struct {
	Query string `yaml:"query"`
	From  int    `yaml:"from"`
	To    int    `yaml:"to"`
}

type Spec struct {
	Auth          `yaml:"auth"`
	DatadogFilter DataDogFilter       `yaml:"datadog_filter"`
	Mapping       []InnerFieldMapping `yaml:"mapping"`
}

type Interval struct {
	From int
	To   int
}

type InnerFieldMapping struct {
	Field      string              `yaml:"field"`
	DdField    string              `yaml:"dd_field"`
	InnerField []InnerFieldMapping `yaml:"inner_field,omitempty"`
}
type YamlMapping struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Spec       Spec   `yaml:"spec"`
}
