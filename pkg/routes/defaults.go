package routes

import (
	"encoding/json"
	"net/http"
	"os"
)

// getServerIsUp '/' endpoint cheks if server is up
func getServerIsUp(w http.ResponseWriter, req *http.Request) {
	port := os.Getenv("PORT")

	res, _ := json.Marshal(struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{Status: 200, Message: "server live on port: " + port})

	w.Write(res)
}

// routeNotFound '/*' endpoint for undefined routes
func routeNotFound(w http.ResponseWriter, req *http.Request) {
	res, _ := json.Marshal(struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{Status: 404, Message: "route not found"})

	w.Write(res)
}
