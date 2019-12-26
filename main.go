package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gitlab.com/tuyentv96/go-cloudfunction/handler"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Init article handler
	handler.NewArticleHandler(e)

	e.ServeHTTP(w, r)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Handler)

	log.Fatal(http.ListenAndServe("localhost:8081", r))
}
