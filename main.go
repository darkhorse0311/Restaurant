package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	r := mux.NewRouter()

	r.HandleFunc("/", getServerIsUp).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(routeNotFound)

	fmt.Println("server live on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
