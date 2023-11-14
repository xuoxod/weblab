package render

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/CloudyKit/jet"
	"github.com/justinas/nosurf"
	"github.com/xuoxod/weblab/internal/config"
	"github.com/xuoxod/weblab/pkg/utils"
)

// may also use an absolute path:
var root, _ = os.Getwd()
var View = jet.NewHTMLSet(filepath.Join(root, "views"))
var app *config.AppConfig

func NewRenderer(a *config.AppConfig) {
	app = a
}

func InitViews() {
	View.SetDevelopmentMode(true)
	View.AddGlobal("appver", "0.0.3")
	View.AddGlobal("copyright", utils.CopyrightDate())
	View.AddGlobal("appname", "Awesome Web App")
	View.AddGlobal("appdate", fmt.Sprintf("%v", utils.DateTimeStamp()))
}

func AddDefaultData(r *http.Request) {
	View.AddGlobal("flash", app.Session.PopString(r.Context(), "flash"))
	View.AddGlobal("error", app.Session.PopString(r.Context(), "error"))
	View.AddGlobal("warning", app.Session.PopString(r.Context(), "warning"))
	View.AddGlobal("csrftoken", app.Session.PopString(r.Context(), nosurf.Token(r)))

	if app.Session.Exists(r.Context(), "user_id") {
		View.AddGlobal("isAuthenticated", 1)
	}

	if app.Session.Exists(r.Context(), "admin_id") {
		View.AddGlobal("isAdmin", 1)
	}
}

func Render(w http.ResponseWriter, r *http.Request, tmpl string, vars jet.VarMap, data map[string]interface{}) error {
	view, err := View.GetTemplate(tmpl)
	AddDefaultData(r)

	if vars == nil {
		vars = make(jet.VarMap)
	}

	if data == nil {
		data = make(map[string]interface{})
	}

	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = view.Execute(w, vars, data)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
