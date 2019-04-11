package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/bobbydeveaux/ngis-status/app/common"
	"github.com/bobbydeveaux/ngis-status/app/home"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	router := mux.NewRouter()
	http.Handle("/", httpInterceptor(router))

	router.HandleFunc("/", home.GetHomePage).Methods("GET")

	fileServer := http.StripPrefix("/dist/", http.FileServer(http.Dir("web/dist")))
	http.Handle("/dist/", fileServer)

	http.ListenAndServe(":8181", nil)
}

func httpInterceptor(router http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		startTime := time.Now()

		router.ServeHTTP(w, req)

		finishTime := time.Now()
		elapsedTime := finishTime.Sub(startTime)

		switch req.Method {
		case "GET":
			// We may not always want to StatusOK, but for the sake of
			// this example we will
			common.LogAccess(w, req, elapsedTime)
		case "POST":
			// here we might use http.StatusCreated
		}

	})
}
