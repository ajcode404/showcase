package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"main.go/config"
	"main.go/handlers"
)

func main() {

	l := log.New(os.Stdout, "[showcase-api] ", log.Flags())

	config := config.NewConfig(l)

	port := config.Port()

	host := config.Host()

	postsHandler := handlers.NewPosts(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", postsHandler.GetPosts)

	s := http.Server{
		Addr:         host + ":" + port,
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	go func() {
		l.Printf("starting server on port %s", port)
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	sig := <-sigChannel
	l.Println("Recieved terminate, graceful shutdoen", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}
