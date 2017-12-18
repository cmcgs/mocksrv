package domain

import "github.com/tyndyll/mocksrv/domain"

type Config interface {
	RouteRepository() domain.RouteRepository
	FileServers() map[string]*domain.FileServerConfig
	Proxy() map[string]*domain.ProxyConfig
}
