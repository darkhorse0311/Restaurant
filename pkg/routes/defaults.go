package routes

import (
	"net/http"
)

// GetServerIsUp '/' endpoint cheks if server is up
func GetServerIsUp(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("server is live"))
}

// RouteNotFound '/*' endpoint for undefined routes
func RouteNotFound(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("route not found"))
}
