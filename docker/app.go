package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	mw "github.com/labstack/echo/middleware"
)

const (
	wwwRoot = "./"
)

var (
	httpPort = flag.Int("http", 3000, "http port number")
)

func Init() *echo.Echo {

	e := echo.New()

	e.Debug()

	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.Any("/*", echo.HandlerFunc(func(c echo.Context) (err error) {
		r := c.Request().(*standard.Request).Request
		w := c.Response().(*standard.Response).ResponseWriter

		requestPath := r.URL.Path
		fileSystemPath := wwwRoot + r.URL.Path
		endURIPath := strings.Split(requestPath, "/")[len(strings.Split(requestPath, "/"))-1]
		splitPath := strings.Split(endURIPath, ".")
		splitLength := len(splitPath)
		if splitLength > 1 && splitPath[splitLength-1] != "go" {
			f, error := os.Stat(fileSystemPath)
			if error == nil && !f.IsDir() {
				http.ServeFile(w, r, fileSystemPath)
				return
			}
		}
		http.ServeFile(w, r, wwwRoot+"index.html")
		return
	}))

	return e
}

func main() {
	flag.Parse()

	server := Init()
	server.Run(standard.New(fmt.Sprintf(`:%d`, *httpPort)))
}
