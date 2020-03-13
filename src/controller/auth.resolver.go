package controller

import (
	"net/http"
	"github.com/go-chi/render"
)

func Authentication(w http.ResponseWriter, r *http.Request)  {
	render.JSON(w, r, "Hello")
}