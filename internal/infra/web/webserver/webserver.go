package webserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	return &WebServer{
		Router:        router,
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddGetHandler(path string, handler http.HandlerFunc) {
	s.Router.Get(path, handler)
}

func (s *WebServer) AddPostHandler(path string, handler http.HandlerFunc) {
	s.Router.Post(path, handler)
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	fmt.Printf("Starting webserver on port %s\n", s.WebServerPort)
	http.ListenAndServe(":"+s.WebServerPort, s.Router)
}
