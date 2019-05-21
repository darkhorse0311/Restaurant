package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/reynld/carbtographer/server/models"
)

// Server struct
type Server struct {
	DB     *sql.DB
	Router *mux.Router
	Cache  *redis.Client
}

// Initialize maps DB and Router instance to Server struct
func (s *Server) Initialize() {
	s.InitializeRouter()
	s.DB = models.InitializeDB()
	s.Cache = models.InitializeCache()
}

// Run runs the server router
func (s *Server) Run() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	fmt.Printf("server live on port%s\n", port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS()(s.Router)))
}

// InitializeRouter maps all the routes the their handlers
func (s *Server) InitializeRouter() {
	s.Router = mux.NewRouter()
	s.Router.Use(loggingMiddleware)

	// Default endpoint
	s.Router.HandleFunc("/", getServerIsUp).Methods("GET")

	// Restaurant endpoints
	s.Router.HandleFunc("/names", s.GetNames).Methods("GET")
	s.Router.HandleFunc("/items/{id}", s.GetItems).Methods("GET")
	s.Router.HandleFunc("/locations/{lat}/{lon}", s.GetLocations).Methods("GET")

	// User endpoints
	s.Router.HandleFunc("/login", s.Login).Methods("POST")
	s.Router.HandleFunc("/register", s.Register).Methods("POST")

	// 404 endpoint
	s.Router.NotFoundHandler = http.HandlerFunc(routeNotFound)
}
