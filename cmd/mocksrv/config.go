package main

import "github.com/tyndyll/mocksrv/domain"

type Config interface {
	RouteRepository() domain.RouteRepository
}
