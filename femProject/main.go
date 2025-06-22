package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/shafiuddin05868/go-projects/femProject/internal/app"
	"github.com/shafiuddin05868/go-projects/femProject/internal/routes"
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

	r := routes.SetUpRoutes(app)

	server := &http.Server {
		Addr: fmt.Sprintf(":%d", port),
		Handler: r,
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		app.Logger.Fatal("failed to start server: " + err.Error())
	}
}
