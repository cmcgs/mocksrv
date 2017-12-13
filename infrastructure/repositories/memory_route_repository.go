package repositories

import "github.com/tyndyll/mocksrv/domain"

type MemoryRouteRepository struct {
	RouteMap map[string]*domain.RouteConfig
}

func (repo *MemoryRouteRepository) Get(route string) (*domain.RouteConfig, error) {
	config, _ := repo.RouteMap[route]
	return config, nil
}
