package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/CloudyKit/jet"
	"github.com/justinas/nosurf"
	"github.com/xuoxod/weblab/internal/forms"
	"github.com/xuoxod/weblab/internal/helpers"
	"github.com/xuoxod/weblab/internal/models"
	"github.com/xuoxod/weblab/internal/render"
)

func (m *Respository) Authenticate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post Login")

	obj := make(map[string]interface{})

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

	// Validate form
	form := forms.New(r.PostForm)
	form.Required("email", "password")
	form.IsEmail("email")

	if !form.Valid() {
		obj["ok"] = false
		obj["form"] = form
		obj["signinform"] = signinform

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
		return
	}

	// Authenticate user
	u, p, s, err := m.DB.Authenticate(signinform.Email, signinform.Password)

	if err != nil {
		fmt.Println("Authentication Error:\t", err.Error())

		obj["ok"] = false
		obj["msg"] = "Invalid Login Credentials"
		obj["type"] = "error"

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
		return
	}

	var user models.User
	var profile models.Profile
	var preferences models.Preferences

	user = u
	profile = p
	preferences = s

	// Put user in session
	m.App.Session.Put(r.Context(), "user_id", user)
	m.App.Session.Put(r.Context(), "profile", profile)
	m.App.Session.Put(r.Context(), "preferences", preferences)

	if user.AccessLevel == 1 {
		m.App.Session.Put(r.Context(), "admin_id", user)
	}

	// Send back JSON results
	obj["title"] = "Home"
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

func (m *Respository) PostRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post Register")
	err := r.ParseForm()

	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	registration := models.Registration{
		FirstName:       r.Form.Get("fname"),
		LastName:        r.Form.Get("lname"),
		Email:           r.Form.Get("email"),
		Phone:           r.Form.Get("phone"),
		PasswordCreate:  r.Form.Get("pwd1"),
		PasswordConfirm: r.Form.Get("pwd2"),
	}

	fmt.Println("Registration posted")

	form := forms.New(r.PostForm)
	form.Required("fname", "lname", "email", "phone", "pwd1", "pwd2")
	form.MinLength("fname", 2, r)
	form.MinLength("lname", 2, r)
	form.IsEmail("email")
	form.PasswordsMatch("pwd1", "pwd2", r)

	if !form.Valid() {
		fmt.Println(form.Errors)

		// vars := make(jet.VarMap)
		// vars.Set("title", "Registration")

		data := make(map[string]string)
		data["title"] = "Registration'"

		obj := make(map[string]interface{})
		obj["csrftoken"] = nosurf.Token(r)
		obj["registrationform"] = registration
		obj["form"] = form

		vars := make(jet.VarMap)
		vars.Set("csrftoken", nosurf.Token(r))

		err := render.Render(w, r, "landing/register.jet", vars, obj)

		if err != nil {
			log.Println(err.Error())
		}
	} else {
		// Send user sms to confirm that it's them

		// Create new user in the database
		// ERROR: duplicate key value violates unique constraint "users_un" (SQLSTATE 23505)
		userId, err := m.DB.CreateUser(registration)

		if err != nil {
			fmt.Println("Account with that email already exists")
			sErr := err.Error()
			uniqueErr := strings.HasSuffix(sErr, "(SQLSTATE 23505)")

			if uniqueErr {
				fmt.Println("Record already exists")
				var registrationErrData models.RegistrationErrData

				regErrData := make(map[string]string)
				regErrData["title"] = "Home"
				regErrData["error"] = "Registration Error"
				regErrData["type"] = "error"
				regErrData["msg"] = "Account already exists"

				registrationErrData.Data = regErrData
				m.App.Session.Put(r.Context(), "reg-error", registrationErrData)
			}

			/* vars := make(jet.VarMap)
			vars.Set("title", "Registration")

			data := make(map[string]interface{})
			data["csrftoken"] = nosurf.Token(r)
			data["registrationform"] = registration
			data["form"] = form */

			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return
		}

		if userId > 0 {
			fmt.Println("User created successfully")
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
