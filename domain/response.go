package domain

type Response struct {
	ContentType string            `yaml:"content-type"`
	Body        string            `yaml:"body"`
	Code        int               `yaml:"response-code"`
	Headers     map[string]string `yaml:"headers"`
}
