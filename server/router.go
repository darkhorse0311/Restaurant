package server

import (
	"log"
	"net/http"

	"github.com/reynld/carbtographer/server/auth"

	"github.com/reynld/carbtographer/server/business"
)

// loggingMiddleware logs HTTP request
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// getServerIsUp '/' endpoint cheks if server is up
func getServerIsUp(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("server is live"))
}

// routeNotFound '/*' endpoint for undefined routes
func routeNotFound(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("route not found"))
}

//////////////////
//// Restaurant handlers
//////////////////

// GetNames handler for /names endpoint
func (s *Server) GetNames(w http.ResponseWriter, r *http.Request) {
	business.GetNames(s.DB, w, r)
}

// GetItems handler for /items/{id} endpoint
func (s *Server) GetItems(w http.ResponseWriter, r *http.Request) {
	business.GetItems(s.DB, w, r)
}

// GetLocations handler for /locations/{lat}/{lon} endpoint
func (s *Server) GetLocations(w http.ResponseWriter, r *http.Request) {
	business.GetLocations(s.DB, s.Cache, w, r)
}

//////////////////
//// User handlers
//////////////////

// Login handler for /login endpoint
func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	auth.Login(s.DB, w, r)
}

// Register handler for /register endpoint
func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	auth.Register(s.DB, w, r)
}
