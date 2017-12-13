package adapters

import (
	//"encoding/json"
	//"io/ioutil"
	"log"
	"net/http"
)

func EchoMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// TODO: Check this...
		//data, err := ioutil.ReadAll(req.Body)
		//if err != nil {
		//	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		//}

		log.Printf("%+v", req)
		next.ServeHTTP(w, req)
	})
}
