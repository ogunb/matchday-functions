package model

type Config struct {
	ProjectID               string `yaml:"projectID"`
	Location                string `yaml:"location"`
	Queue                   string `yaml:"queue"`
	HandlerFunctionEndpoint string `yaml:"handlerFunctionEndpoint"`
	ServiceAccountEmail     string `yaml:"serviceAccountEmail"`
}
