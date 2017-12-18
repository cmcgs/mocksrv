package adapters

import (
	"net/http"

	"github.com/tyndyll/mocksrv/domain"
)

const (
	defaultContentType  = `text/plain`
	defaultResponseCode = 200
)

type RouteMapping interface {
	GetConfig(string) (*domain.RouteConfig, error)
}

type StrictRouteHandler struct {
	Mapping RouteMapping
}

func (handler *StrictRouteHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	config, err := handler.Mapping.GetConfig(req.URL.Path)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if config == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	response := handler.getResponse(config, req.Method)
	if response == nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if response.ContentType == "" {
		response.ContentType = defaultContentType
	}
	w.Header().Add(`Content-Type`, response.ContentType)

	if response.Code == 0 {
		response.Code = defaultResponseCode
	}
	w.WriteHeader(response.Code)
	w.Write([]byte(response.Body))
}

func (handler *StrictRouteHandler) getResponse(config *domain.RouteConfig, method string) *domain.Response {
	switch method {
	case http.MethodGet:
		return config.Get
	case http.MethodHead:
		return config.Head
	case http.MethodPost:
		return config.Post
	case http.MethodPut:
		return config.Put
	case http.MethodPatch:
		return config.Patch
	case http.MethodDelete:
		return config.Delete
	case http.MethodConnect:
		return config.Connect
	case http.MethodOptions:
		return config.Options
	case http.MethodTrace:
		return config.Trace
	}
	return nil
}
