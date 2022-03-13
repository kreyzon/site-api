package routes

import (
	"net/http"
	"site/api/config/server"
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

	authGroup.Use(r.app.AuthModule.MiddlewareFunc())
	{
		r.app.Route("/auth", http.MethodPost, "/logout", r.app.AuthModule.LogoutHandler)
	}
	r.app.Route("/auth", http.MethodGet, "/refresh_token", r.app.AuthModule.RefreshHandler)
}
