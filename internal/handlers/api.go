package handlers

import (
	"github.com/go-chi/chi"                      // Router package to create routes
	chimiddle "github.com/go-chi/chi/middleware" // Middleware utilities from chi with alias 'chimiddle'
	// Your own internal middleware package
)

func Handler(r *chi.Mux) {
	//middleware to regex out trailing /s from req
	r.Use(chimiddle.StripSlashes) //this prebuilt chi function regexes out trailing /s
	r.Route("/account", func(router chi.Router) {
		//additional middleware for /account route - any req has to pass thorugh auth function first
		router.Use(middleware.Authorization)
		router.Get("coins", getCoinBalance)
	})
}

//1:0135