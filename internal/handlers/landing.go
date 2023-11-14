package handlers

import (
	"log"
	"net/http"

	"github.com/xuoxod/weblab/internal/config"
	"github.com/xuoxod/weblab/internal/driver"
	"github.com/xuoxod/weblab/internal/render"
)

// Repo the repository used by the handlers
var Repo *Respository

// Repository the Repository type
type Respository struct {
	App *config.AppConfig
	// DB  repository.DatabaseRepo
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Respository {
	return &Respository{
		App: a,
		// DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// NewHandler sets the repository for the handlers
func NewHandler(r *Respository) {
	Repo = r
	render.InitViews()
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["title"] = "Home"

	err := render.Render(w, "landing/home.jet", nil, data)

	if err != nil {
		log.Println(err.Error())
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["title"] = "About"

	err := render.Render(w, "landing/about.jet", nil, data)

	if err != nil {
		log.Println(err.Error())
	}
}
