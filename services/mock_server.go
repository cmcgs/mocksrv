package services

import (
	"fmt"
	"log"
	"net/http"
)

type MockServer struct {
	*http.Server
}

func (mock *MockServer) Start(port int) {
	mock.Addr = fmt.Sprintf(":%d", port)
	go func() {
		if err := mock.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
}
