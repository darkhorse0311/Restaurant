package routes

import (
	"net/http"
)

// getServerIsUp '/' endpoint cheks if server is up
func getServerIsUp(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("server is live"))
}

// routeNotFound '/*' endpoint for undefined routes
func routeNotFound(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("route not found"))
}
