package main

import (
	"io/ioutil"
	"os"

	"github.com/tyndyll/mocksrv/domain"
	"github.com/tyndyll/mocksrv/infrastructure/repositories"
	"gopkg.in/yaml.v2"
)

type YAMLConfig struct {
	Routes        map[string]*domain.RouteConfig      `yaml:"routes"`
	FileServerMap map[string]*domain.FileServerConfig `yaml:"files"`
	ProxyMap      map[string]*domain.ProxyConfig      `yaml:"proxy"`
}

func (config *YAMLConfig) RouteRepository() domain.RouteRepository {
	return &repositories.MemoryRouteRepository{
		RouteMap: config.Routes,
	}
}

func (config *YAMLConfig) FileServers() map[string]*domain.FileServerConfig {
	return config.FileServerMap
}

func (config *YAMLConfig) Proxy() map[string]*domain.ProxyConfig {
	return config.ProxyMap
}

func configFromYAML(path string) (domain.Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	config := &YAMLConfig{}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return config, nil
}
