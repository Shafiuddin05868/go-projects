package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/shafiuddin05868/go-projects/femProject/internal/app"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 4040, "port to listen ton ")
	flag.Parse()
	app, err := app.NewApplication()
	if err != nil {
		panic("failed to create application" + err.Error())
	}
	app.Logger.Printf("Server is running at port, %d", port)

	http.HandleFunc("/health", HealthCheck)

	server := &http.Server {
		Addr: ":3055",
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		app.Logger.Fatal("failed to start server: " + err.Error())
	}
}


func HealthCheck (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}