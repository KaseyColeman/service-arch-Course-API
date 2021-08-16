package main

import (
	"context"
	"log"
	"module/data"
	"module/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	goHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)


func main() {
	l := log.New(os.Stdout, "college-course-api ", log.LstdFlags)
	v := data.NewValidation()

	// create the handlers
	ph := handlers.NewCourse(l, v)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	// handlers for API
	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/courses", ph.ListAll)
	getR.HandleFunc("/courses/{id:[0-9]+}", ph.ListSingle)

	putR := sm.Methods(http.MethodPut).Subrouter()
	putR.HandleFunc("/courses", ph.Update)
	putR.Use(ph.MiddlewareValidateCourse)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/courses", ph.Create)
	postR.Use(ph.MiddlewareValidateCourse)

	deleteR := sm.Methods(http.MethodDelete).Subrouter()
	deleteR.HandleFunc("/courses/{id:[0-9]+}", ph.Delete)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getR.Handle("/docs", sh)
	getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	ch := goHandlers.CORS(goHandlers.AllowedOrigins([] string{"http://localhost:3000"}))



	// create a new server
	s := http.Server{
		Addr:         ":9090",      // configure the bind address
		Handler:      ch(sm),                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}