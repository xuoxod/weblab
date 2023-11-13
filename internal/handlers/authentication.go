package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/xuoxod/weblab/internal/models"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post Login")

	err := r.ParseForm()

	if err != nil {
		log.Println(err.Error())
		return
	}

	signinform := models.Signin{
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("password"),
	}

	fmt.Println("Signin posted")

	// Authenticate user

	// Send back JSON results
	obj := make(map[string]interface{})
	obj["title"] = "Home"
	obj["signinform"] = signinform
	obj["ok"] = true

	out, err := json.MarshalIndent(obj, "", " ")

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, rErr := w.Write(out)

	if rErr != nil {
		log.Println(err)
	}
}
