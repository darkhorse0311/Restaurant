package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

// InitializeRouter maps all the routes the their handlers
func (s *Server) InitializeRouter() {
	s.Router = mux.NewRouter()
	s.Router.Use(loggingMiddleware)

	s.Router.HandleFunc("/", getServerIsUp).Methods("GET")
	s.Router.HandleFunc("/names", s.GetNames).Methods("GET")
	s.Router.HandleFunc("/items/{id}", s.GetItems).Methods("GET")
	s.Router.HandleFunc("/locations/{lat}/{lon}", s.Locations).Methods("GET")
	s.Router.HandleFunc("/signin", s.Signin).Methods("POST")
	s.Router.HandleFunc("/register", s.Register).Methods("POST")
	s.Router.NotFoundHandler = http.HandlerFunc(routeNotFound)
}
