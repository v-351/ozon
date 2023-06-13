package server

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/v-351/url-shortener/internal/service"
)

const portNumber = ":8080"

type Server struct {
	Server  *http.Server
	Service *service.Service
}

func (s *Server) Run() {
	s.Server = &http.Server{Addr: portNumber}
	r := chi.NewRouter()
	r.Get("/{shortURL}", s.getURL)
	r.Post("/", s.postURL)
	s.Server.Handler = r

	log.Printf("Starting application on port %v\n", portNumber)
	err := s.Server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) postURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("POST:", string(body), err)
		return
	}
	log.Println("POST:", string(body))
	short, err := s.Service.Put(string(body))
	if err != nil {
		log.Println(err.Error())
		fmt.Fprint(w, err.Error())
		return
	}
	log.Println(string(body), "->", short)
	fmt.Fprint(w, short)
}

func (s *Server) getURL(w http.ResponseWriter, r *http.Request) {

	shortURL := chi.URLParam(r, "shortURL")

	log.Println("GET:", shortURL)
	raw, err := s.Service.Get(shortURL)
	if err != nil {
		log.Println(err.Error())
		fmt.Fprint(w, err.Error())
		return
	}
	log.Println(shortURL, "->", raw)
	fmt.Fprint(w, raw)
}
