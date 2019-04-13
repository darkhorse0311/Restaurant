package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reynld/carbtographer/pkg/business"
	"github.com/reynld/carbtographer/pkg/yelp"
)

// Routes struct
type Routes struct {
	Router *mux.Router
}

// loggingMiddleware logs HTTP request
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// GetServerIsUp '/' endpoint cheks if server is up
func GetServerIsUp(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("server is live"))
}

// RouteNotFound '/*' endpoint for undefined routes
func RouteNotFound(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("route not found"))
}

// InitializeRouter maps all the routes the their handlers
func (r *Routes) InitializeRouter() {
	r.Router = mux.NewRouter()
	r.Router.Use(loggingMiddleware)

	r.Router.HandleFunc("/", GetServerIsUp).Methods("GET")
	r.Router.HandleFunc("/names", business.GetNames).Methods("GET")
	r.Router.HandleFunc("/items/{id}", business.GetItems).Methods("GET")
	r.Router.HandleFunc("/locations/{lat}/{lon}", yelp.GetLocations).Methods("GET")
	r.Router.NotFoundHandler = http.HandlerFunc(RouteNotFound)
}
