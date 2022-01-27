package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"msg-board/api/domain"
	"msg-board/api/service"
	"net/http"
	"os"
	"time"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}

	frontend  string
	secretKey string
}

type application struct {
	config config
	logger *log.Logger
	mg     *MessageHandler
	//ch    CustomerHandler
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

	app.logger.Printf("ServingStarting Back end server in %s mode on port %d", app.config.env, app.config.port)

	return srv.ListenAndServe()
}

func main() {
	log.Println("Api Starting...")

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	var cfg config

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	cfg.db.dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	flag.IntVar(&cfg.port, "port", 4001, "Server port to listen on")
	flag.StringVar(&cfg.frontend, "frontend", "http://localhost:4000", "url to frontend")

	flag.Parse()

	logger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	dbClient, err := domain.GetDBClient(cfg.db.dsn)
	if err != nil {
		log.Println("failed to connect mysql " + err.Error())
	}

	messageRepositoryDb := domain.NewMessageRepository(dbClient)

	//建立各個Handlers
	//ch := CustomerHandler{service.NewCustomerService(customerRepositoryDb)}

	mg := NewMessageHandler(service.NewMessageService(messageRepositoryDb))

	app := &application{
		config: cfg,
		logger: logger,
		mg:     mg,
	}

	err = app.serve()
	if err != nil {
		app.logger.Println(err.Error())
		log.Fatal(err)
	}

}
