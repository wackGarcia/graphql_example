package main

import (
	"log"
	"net/http"
	"graphql_example/routes"

	"graphql_example/tools/viper"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-chi/cors"
)

func initServer()  {
	
	server := chi.NewRouter()
	
	//confiig cors domain
	confCors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	
	server.Use(confCors.Handler)

	server.Use(
		render.SetContentType(render.ContentTypeJSON), 
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.StripSlashes,
		middleware.Recoverer,
	)
	log.Println("Serve was load on port ", viper.GetEnv("SERVER_PORT"), " !")
	log.Fatal(http.ListenAndServe(viper.GetEnv("SERVER_PORT"), routes.RouterV1(server)))
}