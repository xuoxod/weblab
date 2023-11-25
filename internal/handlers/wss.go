package handlers

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/xuoxod/weblab/internal/render"
)

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["title"] = "Home"
	data["csrf_token"] = nosurf.Token(r)

	err := render.Render(w, r, "landing/home.jet", nil, data)

	if err != nil {
		log.Println(err.Error())
	}
}
