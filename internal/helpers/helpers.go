package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/xuoxod/weblab/internal/config"
)

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of ", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	debug.PrintStack()
	trace := fmt.Sprintf("\n%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func IsAuthenticated(r *http.Request) bool {
	exists := app.Session.Exists(r.Context(), "user_id")

	return exists
}

func IsAdmin(r *http.Request) bool {
	isAdmin := app.Session.Exists(r.Context(), "admin_id")
	isAuthed := app.Session.Exists(r.Context(), "user_id")

	return isAdmin && isAuthed
}
