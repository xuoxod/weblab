package handlers

import (
	"log"
	"net/http"

	"github.com/xuoxod/weblab/internal/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["title"] = "Home"

	err := render.Render(w, "landing/home.jet", nil, data)

	if err != nil {
		log.Println(err.Error())
	}
}
