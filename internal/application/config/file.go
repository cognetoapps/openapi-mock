package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type fileConfiguration struct {
	OpenAPI     openapiConfiguration     `json:"openapi" yaml:"openapi"`
	HTTP        httpConfiguration        `json:"http" yaml:"http"`
	Application applicationConfiguration `json:"application" yaml:"application"`
	Generation  generationConfiguration  `json:"generation" yaml:"generation"`
}

type openapiConfiguration struct {
	SpecificationURL string `json:"specification_url" yaml:"specification_url"`
	urlFromEnv       bool
}

type httpConfiguration struct {
	Port            *uint16 `json:"port" yaml:"port" valid:"range(1|65535)"`
	CORSEnabled     bool    `json:"cors_enabled" yaml:"cors_enabled"`
	ResponseTimeout float64 `json:"response_timeout" yaml:"response_timeout"`
	InjectDelay     bool    `json:"inject_delay" yaml:"inject_delay"`
	DelayExpRate    float64 `json:"delay_exp_rate" yaml:"delay_exp_rate"`
	DelayMinFloat   float64 `json:"delay_min_float" yaml:"delay_min_float"`
}

type applicationConfiguration struct {
	Debug     bool   `json:"debug" yaml:"debug"`
	LogFormat string `json:"log_format" yaml:"log_format" valid:"in(tty|json)"`
	LogLevel  string `json:"log_level" yaml:"log_level" valid:"in(panic|fatal|error|warn|warning|info|debug|trace)"`
}

type generationConfiguration struct {
	DefaultMinFloat *float64 `json:"default_min_float" yaml:"default_min_float"`
	DefaultMaxFloat *float64 `json:"default_max_float" yaml:"default_max_float"`
	DefaultMinInt   *int64   `json:"default_min_int" yaml:"default_min_int"`
	DefaultMaxInt   *int64   `json:"default_max_int" yaml:"default_max_int"`
	LoripsumLength  string   `json:"loripsum_length" yaml:"loripsum_length"`
	NullProbability *float64 `json:"null_probability" yaml:"null_probability"`
	SuppressErrors  bool     `json:"suppress_errors" yaml:"suppress_errors"`
	UseExamples     string   `json:"use_examples" yaml:"use_examples" valid:"in(no|if_present|exclusively)"`
}

func loadFileConfiguration(filename string) (*fileConfiguration, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, &ErrLoadFailed{Previous: err}
	}

	var fileConfig fileConfiguration
	err = yaml.Unmarshal(data, &fileConfig)
	if err != nil {
		return nil, &ErrLoadFailed{Previous: err}
	}

	return &fileConfig, nil
}
