package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type App struct {
	Router *gin.Engine
}

func NewApp() *App {
	app := &App{
		Router: gin.New(),
	}

	app.initializeMiddleware()

	app.setupHTMLRendering()

	app.registerRoutes()

	return app
}

func (a *App) Run(addr string) error {
	return a.Router.Run(addr)
}

func (a *App) initializeMiddleware() {
	a.Router.Use(gin.Logger())
	a.Router.Use(gin.Recovery())
}

func (a *App) registerRoutes() {
	a.Router.GET("/labs", func(c *gin.Context) {
		c.HTML(http.StatusOK, "labs.html", nil)
	})

	a.Router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})
}

func (a *App) setupHTMLRendering() {
	a.Router.LoadHTMLGlob("views/**/*.html")
}
