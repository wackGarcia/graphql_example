package routes

import (
	"graphql_example/src/controller"

	"github.com/go-chi/chi"
)

func RouterV1(router *chi.Mux) *chi.Mux {

	router.Post("/v1/auth", controller.Authentication)
	// router.Mount("/v1/subscription",authRouter(s))

	return router
}

// func authRouter(s server.Server) http.Handler {
// 	router := chi.NewRouter()
// 	//router.Use(middlewares.Login)
// 	router.Post("/",s.GraphQL())
// 	return router
// }

