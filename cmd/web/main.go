package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/xuoxod/weblab/internal/config"
	"github.com/xuoxod/weblab/internal/driver"
	"github.com/xuoxod/weblab/internal/envloader"
	"github.com/xuoxod/weblab/internal/handlers"
	"github.com/xuoxod/weblab/internal/helpers"
	"github.com/xuoxod/weblab/internal/render"
)

// Application configuration
var app config.AppConfig

var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func initApp() {
	var devMode bool
	flag.BoolVar(&devMode, "mode", false, "Sets true for production mode or false otherwise - defaults to false")

	err := envloader.LoadEnvVars()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.InProduction = devMode
	app.DBConnection = os.Getenv("DB_URL")

	// App config middleware
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// Session middleware
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = app.InProduction
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	// Config app level config with session
	app.Session = session
}

func initDB() (*driver.DB, error) {
	var host string = os.Getenv("DB_HOST")
	var user string = os.Getenv("DB_USER")
	var password string = os.Getenv("DB_PASSWD")
	var dbname string = os.Getenv("DB_NAME")
	var port int = 5432

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := driver.ConnectSql(psqlInfo)

	if err != nil {
		log.Println("Cannot connect to database! Dying ...")
		return nil, err
	}

	fmt.Println("Connected to datastore")

	return db, nil
}

func main() {
	initApp()
	db, err := initDB()

	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()

	render.InitViews()
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandler(repo)
	helpers.NewHelpers(&app)

	mux := routes()

	log.Println("Server running on port 8080")

	_ = http.ListenAndServe(":8080", mux)
}
