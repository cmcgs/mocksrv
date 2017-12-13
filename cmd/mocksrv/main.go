package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/tyndyll/mocksrv/adapters"
	"github.com/tyndyll/mocksrv/usecases"
)

var (
	echoFlag bool
	portFlag int
	jsonPath string
)

func init() {
	flag.BoolVar(&echoFlag, "echo", true, "echo request")
	flag.IntVar(&portFlag, "port", 8000, "port the server should listen on")
	flag.Parse()
}

func main() {
	config, err := configFromYAML(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}

	mapper := &usecases.RouteMapping{
		RouteRepository: config.RouteRepository(),
	}

	handler := &adapters.RouteHandler{
		Mapping: mapper,
	}

	var handlerFunc http.Handler = handler
	if echoFlag {
		handlerFunc = adapters.EchoMiddleware(handlerFunc)
	}
	http.Handle(`/`, handlerFunc)
	http.ListenAndServe(fmt.Sprintf(":%d", portFlag), nil)

	//log.Printf("%+v \n", repo.Routes["/api/users"].Get.ReponseBody)
}
