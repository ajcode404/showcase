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

	postsHandler := handlers.NewPosts(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", postsHandler.GetPosts)

	s := http.Server{
		Addr:         config.Host() + ":" + config.Port(),
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	go func() {
		l.Printf("starting server on host=%s and port=%s", config.Host(), config.Port())
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)

	sig := <-sigChannel
	l.Println("Recieved terminate, graceful shutdoen", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}
