package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
)

// Post represents a blog post.
type Post struct {
	ID   int
	Body string
}

type server struct {
	sync.Mutex
	posts []*Post
}

func (s *server) save(p *Post) {
	s.Lock()
	defer s.Unlock()
	p.ID = len(s.posts) + 1
	s.posts = append(s.posts, p)
}

func (s *server) routes() {
	http.HandleFunc("/posts", s.handleIndex())
	http.HandleFunc("/posts/new", s.handleNew())
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, s.posts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (s *server) handleNew() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := &Post{Body: r.FormValue("body")}
		s.save(p)
		http.Redirect(w, r, fmt.Sprintf("/posts#%d", p.ID), http.StatusFound)
	}
}

func main() {
	srv := server{}
	srv.routes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
