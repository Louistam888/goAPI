package main

import (
	"fmt"
	"net/http"

	"goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	//this enables logging of errors
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter() // pointer to a mux type like nodejs app.router
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")
	//in GO return errors even if stuff goes well
	err := http.ListenAndServe("localhost:8000", r)

	if err != nil {
		log.Error(err)
	}
}
