package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/xuoxod/weblab/internal/render"
)

func main() {
	fmt.Println("Awesome Web App!")
	render.InitViews()
	mux := routes()

	log.Println("Server running on port 8080")

	_ = http.ListenAndServe(":8080", mux)
}
