package apiserver

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type APIServer struct {
	host   string
	server *http.Server
	routes map[string]http.HandlerFunc
}

func CreateNewServer(port string) APIServer {
	fmt.Println("Creating New Server...")
	return APIServer{
		host:   port,
		routes: make(map[string]http.HandlerFunc), // Initialize the map
		server: &http.Server{},                    // Initialize the server
	}
}

func (s *APIServer) AddRoute(route string, handleFunc http.HandlerFunc) {
	fmt.Printf("Route Added: %s", route)
	s.routes[route] = handleFunc
}

func (s *APIServer) Run() {
	fmt.Println("Running Server")
	serverMux := http.NewServeMux()

	for route, handler := range s.routes {
		serverMux.Handle(route, handler)
	}

	s.server.Addr = s.host
	s.server.Handler = serverMux
	s.server.ReadTimeout = 10 * time.Second
	s.server.WriteTimeout = 10 * time.Second

	log.Fatal(s.server.ListenAndServe())

}
