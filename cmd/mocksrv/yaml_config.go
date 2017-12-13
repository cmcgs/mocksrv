package main

import (
	"io/ioutil"
	"os"

	"fmt"
	"github.com/tyndyll/mocksrv/domain"
	"github.com/tyndyll/mocksrv/infrastructure/repositories"
	"gopkg.in/yaml.v2"
)

type YAMLConfig struct {
	Paths map[string]*domain.RouteConfig `yaml:"routes"`
}

func (config *YAMLConfig) RouteRepository() domain.RouteRepository {
	return &repositories.MemoryRouteRepository{
		RouteMap: config.Paths,
	}
}

func configFromYAML(path string) (Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(data))
	config := &YAMLConfig{}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return config, nil
}
