package routes

import (
	"fmt"
	"net/http"
	"site/api/config/server"
	"site/api/pkg/controllers"
	"site/api/pkg/controllers/portfolio"

	"github.com/gin-gonic/gin"
)

type route struct {
	app           server.Application
	portfolioCtrl controllers.PortfolioController
}

func NewRoute(app server.Application) (*route, error) {
	portfolioCtrl, err := portfolio.New(portfolio.WithGorm(app.DatabaseClient))
	if err != nil {
		return nil, fmt.Errorf("error in initialize portfoilio routes: %w", err)
	}

	return &route{
		app:           app,
		portfolioCtrl: portfolioCtrl,
	}, nil
}

func (r *route) InitRoutes() {
	r.app.RouteGroup("/portfolio")
	r.app.Route("/portfolio", http.MethodGet, "", r.HandleGetAllPortfolios())
	r.app.Route("/portfolio", http.MethodPost, "", r.HandleCreatePortfolio())
}

func (r *route) HandleGetAllPortfolios() gin.HandlerFunc {
	return func(c *gin.Context) {
		portfolios, err := r.portfolioCtrl.ReadAll()
		if err != nil {
			statusError := fmt.Errorf("failed read all portfolios: %w", err)
			c.Error(statusError)
			return
		}
		c.JSON(http.StatusOK, portfolios)
	}
}

func (r *route) HandleCreatePortfolio() gin.HandlerFunc {
	return func(c *gin.Context) {
		port := controllers.PortfolioDTO{}
		err := c.Bind(&port)
		if err != nil {
			statusError := fmt.Errorf("failed bind body data in create portfolio: %w", err)
			c.Error(statusError)
			return
		}
		id, err := r.portfolioCtrl.Create(port)
		if err != nil {
			statusError := fmt.Errorf("failed read all portfolios: %w", err)
			c.Error(statusError)
			return
		}
		c.JSON(http.StatusOK, map[string]int{"id": id})
	}
}
