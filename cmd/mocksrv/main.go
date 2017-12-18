package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/tyndyll/mocksrv/adapters"
	"github.com/tyndyll/mocksrv/usecases"
)

var (
	echoFlag bool
	portFlag int
)

func init() {
	flag.BoolVar(&echoFlag, "echo", false, "echo request")
	flag.IntVar(&portFlag, "port", 8000, "port the server should listen on")
	flag.Parse()
}

func main() {
	config, err := configFromYAML(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}

	for route, fileServer := range config.FileServers() {
		fs := http.FileServer(http.Dir(fileServer.Directory))
		http.Handle(route, http.StripPrefix(route, fs))
	}

	for route, proxy := range config.Proxy() {
		remoteURL, err := url.Parse(proxy.Remote)
		if err != nil {
			log.Fatalln(err)
		}
		http.Handle(route, httputil.NewSingleHostReverseProxy(remoteURL))
	}

	mapper := &usecases.RouteMapping{
		RouteRepository: config.RouteRepository(),
	}

	handler := &adapters.StrictRouteHandler{
		Mapping: mapper,
	}

	var handlerFunc http.Handler = handler
	if echoFlag {
		handlerFunc = adapters.EchoMiddleware(handlerFunc)
	}
	http.Handle(`/`, handlerFunc)

	http.ListenAndServe(fmt.Sprintf(":%d", portFlag), nil)
}
