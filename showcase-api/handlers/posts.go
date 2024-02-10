package handlers

import (
	"log"
	"net/http"

	"main.go/data"
)

func NewPosts(l *log.Logger) *Posts {
	return &Posts{l}
}

type Posts struct {
	l *log.Logger
}

func (b *Posts) GetPosts(rw http.ResponseWriter, r *http.Request) {
	b.l.Println("Handle GET posts")
	lp := data.GetBlogs()
	if err := lp.ToJSON(rw); err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
