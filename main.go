package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tuyentv96/go-cloudfunction/handler"
)

var e *echo.Echo

func init() {
	// Echo instance
	e = echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	log.Println("initial echo")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Init article handler
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200,"hello world")
	})
	handler.NewArticleHandler(e)

	e.ServeHTTP(w, r)
}

//func main() {
//	r := mux.NewRouter()
//	r.HandleFunc("/", Handler)
//
//	log.Fatal(http.ListenAndServe("localhost:8081", r))
//}
