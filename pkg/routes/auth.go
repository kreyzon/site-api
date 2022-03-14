package routes

import (
	"net/http"
	"site/api/config/server"

	"github.com/gin-gonic/gin"
)

type routeAuth struct {
	app server.Application
}

func NewRouteAuth(app server.Application) (*routeAuth, error) {
	return &routeAuth{
		app: app,
	}, nil
}

func (r *routeAuth) InitRoutes() {
	authGroup := r.app.RouteGroup("/auth")
	r.app.Route("/auth", http.MethodPost, "/login", r.app.AuthModule.LoginHandler)
	r.app.Route("/auth", http.MethodOptions, "/login", r.Preflight())

	authGroup.Use(r.app.AuthModule.MiddlewareFunc())
	{
		r.app.Route("/auth", http.MethodPost, "/logout", r.app.AuthModule.LogoutHandler)
	}
	r.app.Route("/auth", http.MethodGet, "/refresh_token", r.app.AuthModule.RefreshHandler)
}
func (r *routeAuth) Preflight() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers, content-type")
		c.JSON(http.StatusOK, struct{}{})
	}
}
