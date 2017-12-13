package domain

type RouteRepository interface {
	Get(url string) (*RouteConfig, error)
}
