package usecases

import "github.com/tyndyll/mocksrv/domain"

type RouteMapping struct {
	RouteRepository domain.RouteRepository
}

func (mapping *RouteMapping) GetConfig(route string) (*domain.RouteConfig, error) {
	return mapping.RouteRepository.Get(route)
}
