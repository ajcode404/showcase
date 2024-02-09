package handlers

import (
	"log"
	"net/http"

	"main.go/data"
)

type Posts struct {
	l *log.Logger
}

func NewPosts(l *log.Logger) *Posts {
	return &Posts{l}
}

func (b *Posts) GetPosts(rw http.ResponseWriter, r *http.Request) {
	b.l.Println("Handle GET posts")
	lp := data.GetBlogs()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
