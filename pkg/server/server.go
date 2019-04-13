package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/reynld/carbtographer/pkg/routes"
)

// Server struct
type Server struct {
	DB     *sql.DB
	Routes routes.Routes
}

// InitializeServer maps DB and Router instance to Server struct
func (s *Server) InitializeServer() {
	s.Routes = routes.Routes{}
	s.Routes.InitializeRouter()
}

// Run runs the server router
func (s *Server) Run() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	fmt.Printf("server live on port%s\n", port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS()(s.Routes.Router)))
}
