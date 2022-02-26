package routes

import (
	"net/http"
	"site/api/config/server"
	"site/api/pkg/controllers"

	"github.com/gin-gonic/gin"
)

type route struct {
	app server.Application
	// portfolioCtrl controllers.PortfolioController
}

func NewRoute(app server.Application) (*route, error) {
	return &route{
		app: app,
	}, nil
}

func (r *route) InitRoutes() {
	r.app.RouteGroup("/portfolio")
	r.app.Route("/portfolio", http.MethodGet, "", r.HandleGetAllPortfolios())
}

func (r *route) HandleGetAllPortfolios() gin.HandlerFunc {
	return func(c *gin.Context) {
		portfolios := []controllers.PortfolioDTO{
			{
				Id:       1,
				Title:    "Teste",
				Summary:  "Um teste de sumário",
				Url:      "",
				ImageUrl: "",
				ImageAlt: "Texto Alternativo de imagem",
			},
			{
				Id:       1,
				Title:    "Teste",
				Summary:  "Um teste de sumário",
				Url:      "",
				ImageUrl: "",
				ImageAlt: "Texto Alternativo de imagem",
			},
			{
				Id:       1,
				Title:    "Teste",
				Summary:  "Um teste de sumário",
				Url:      "",
				ImageUrl: "",
				ImageAlt: "Texto Alternativo de imagem",
			},
			{
				Id:       1,
				Title:    "Teste",
				Summary:  "Um teste de sumário",
				Url:      "",
				ImageUrl: "",
				ImageAlt: "Texto Alternativo de imagem",
			}, {
				Id:       1,
				Title:    "Teste",
				Summary:  "Um teste de sumário",
				Url:      "",
				ImageUrl: "",
				ImageAlt: "Texto Alternativo de imagem",
			},
			{
				Id:       1,
				Title:    "Teste",
				Summary:  "Um teste de sumário",
				Url:      "",
				ImageUrl: "",
				ImageAlt: "Texto Alternativo de imagem",
			},
			{
				Id:       1,
				Title:    "Teste",
				Summary:  "Um teste de sumário",
				Url:      "",
				ImageUrl: "",
				ImageAlt: "Texto Alternativo de imagem",
			},
			{
				Id:       1,
				Title:    "Teste",
				Summary:  "Um teste de sumário",
				Url:      "",
				ImageUrl: "",
				ImageAlt: "Texto Alternativo de imagem",
			},
		}
		c.JSON(http.StatusOK, portfolios)
	}
}
