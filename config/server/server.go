package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"site/api/internal/logger"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	ErrNilEngine         = errors.New("engine is nil")
	ErrInvalidPort       = errors.New("invalid port")
	ErrGroupNotFound     = errors.New("router group not found")
	ErrNilGroups         = errors.New("router groups is nil")
	ErrUnsupportedMethod = errors.New("method not supported yet")
	ErrEmptyBasePath     = errors.New("empty base path")
)

const BasePath string = "/v1"

type Application struct {
	Port              int
	Engine            *gin.Engine
	Routers           map[string]*gin.RouterGroup
	BasePath          string
	globalMiddlewares []gin.HandlerFunc
	DatabaseClient    *gorm.DB
	Logger            logger.Logger
	AuthModule        *jwt.GinJWTMiddleware
}

func (app *Application) RouteGroup(name string) *gin.RouterGroup {
	if app.Engine == nil {
		log.Fatal(ErrNilEngine)
	}

	if app.Routers == nil {
		log.Fatal(ErrNilGroups)
	}

	if app.BasePath == "" {
		log.Fatal(ErrEmptyBasePath)
	}

	fullName := fmt.Sprintf("%s%s", app.BasePath, name)
	group := app.Engine.Group(fullName)

	app.Routers[fullName] = group

	return group
}

func (app *Application) RouteMiddlewaresByGroup(name string, middlewares ...gin.HandlerFunc) {
	fullName := fmt.Sprintf("%s%s", app.BasePath, name)
	group, ok := app.Routers[fullName]
	if !ok {
		log.Fatalf("%v: %s", ErrGroupNotFound, name)
	}
	group.Use(middlewares...)
}

func (app *Application) GlobalMiddleware(middleware gin.HandlerFunc) {
	if app.Routers == nil {
		log.Fatal(ErrNilGroups)
	}
	app.Engine.Use(middleware)
	app.globalMiddlewares = append(app.globalMiddlewares, middleware)
}

func (app *Application) Route(routeGroup, method, route string, handler gin.HandlerFunc, middlewares ...gin.HandlerFunc) {
	if app.Routers == nil {
		log.Fatal(ErrNilGroups)
	}
	fullName := fmt.Sprintf("%s%s", app.BasePath, routeGroup)

	group, ok := app.Routers[fullName]
	if !ok {
		log.Fatalf("%v: %s", ErrGroupNotFound, fullName)
	}

	handlers := make([]gin.HandlerFunc, 0)
	handlers = append(handlers, handler)
	handlers = append(handlers, middlewares...)
	switch method {
	case http.MethodPost:
		group.POST(route, handlers...)
	case http.MethodGet:
		group.GET(route, handlers...)
	case http.MethodPatch:
		group.PATCH(route, handlers...)
	case http.MethodDelete:
		group.DELETE(route, handlers...)
	case http.MethodPut:
		group.PUT(route, handlers...)
	case http.MethodOptions:
		group.OPTIONS(route, handlers...)
	default:
		log.Fatal(ErrUnsupportedMethod)
	}
}

// Run attaches app.Router to a http.Server and starts listening and serving HTTP requests.
// Note: this method will block the calling goroutine indefinitely unless an error happens.
func (app *Application) Run() error {
	if app.Port <= 0 {
		return ErrInvalidPort
	}
	if app.Engine == nil {
		return ErrNilEngine
	}
	app.Engine.Use(CORS())

	return app.Engine.Run(fmt.Sprintf(":%d", app.Port))
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
