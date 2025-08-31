package config

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	HTTPListenAddr string

	//handlers
	NotFoundHandler         http.Handler //custom page for not found 404
	MethodNotAllowedHandler http.Handler //custome page for method not match

	//server properties
	HttpServer     *http.Server
	WriteTimeout   time.Duration
	ReadTimeout    time.Duration
	IdleTimeout    time.Duration
	HandlerTimeout time.Duration
}

func (server *Server) createMuxRouter() *mux.Router {
	router := mux.NewRouter()

	router.NotFoundHandler = server.NotFoundHandler
	router.MethodNotAllowedHandler = server.MethodNotAllowedHandler

	router.SkipClean(true)

	return router
}

func (server *Server) StartServer() {
	router := server.createMuxRouter()

	server.HttpServer = &http.Server{
		Addr:         server.HTTPListenAddr,
		Handler:      router,
		WriteTimeout: server.WriteTimeout,
		ReadTimeout:  server.ReadTimeout,
		IdleTimeout:  server.IdleTimeout,
	}
	server.HttpServer.ListenAndServe()
}
