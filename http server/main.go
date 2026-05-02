package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dev-karani/http-server/internals/app"
)

func main(){
	app, err := app.NewApplication() 

	if err !=nil {
		panic(err)
	}

	app.Logger.Println("we are running out app")

	http.HandleFunc("/health", HealthCheck)
	
	server := &http.Server{
		Addr: ":8080",
		IdleTimeout: time.Minute,
		ReadTimeout: 10 *time.Second,
		WriteTimeout: 10*time.Second,
	}

	err = server.ListenAndServe()
	if err!= nil {
		app.Logger.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "status is available")
}