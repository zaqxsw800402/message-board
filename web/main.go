package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"msg-board/web/model"
	"net/http"
	"os"
	"time"
)

type config struct {
	port int
	api  string
	db   struct {
		dsn string
	}
	frontend string
}

type application struct {
	config        config
	logger        *log.Logger
	templateCache map[string]*template.Template
	DB            *model.DB
	Session       *session.Store
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.logger.Printf("Starting HTTP server  on port %d\n", app.config.port)

	return srv.ListenAndServe()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg config

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	cfg.db.dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.frontend, "frontend", "http://localhost:4000", "url to frontend")

	flag.Parse()

	cfg.api = os.Getenv("API")

	logger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	//conn, err := mysql.OpenDb(cfg.mysql.dsn)
	conn, err := model.GetDBClient(cfg.db.dsn)
	if err != nil {
		logger.Fatal(err)
	}

	// set up store
	store := session.New()
	//session.Config{Storage: sqlite3.New()}

	tc := make(map[string]*template.Template)

	app := &application{
		config:        cfg,
		logger:        logger,
		templateCache: tc,
		DB:            model.New(conn),
		Session:       store,
	}

	err = app.serve()
	if err != nil {
		app.logger.Println(err)
		log.Fatal(err)
	}
}
