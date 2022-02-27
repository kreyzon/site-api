package portfolio

import (
	"fmt"
	"site/api/pkg/controllers"
	"site/api/pkg/models"

	"gorm.io/gorm"
)

type controller struct {
	db *gorm.DB
}

type Option func(*controller) error

func WithGorm(db *gorm.DB) Option {
	return func(c *controller) error {
		c.db = db
		return nil
	}
}

func New(options ...Option) (controllers.PortfolioController, error) {
	c := new(controller)
	for _, opt := range options {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *controller) ReadAll() ([]*controllers.PortfolioDTO, error) {
	portfoliosDB := []models.Portfolio{}
	// Get all records
	result := c.db.Find(&portfoliosDB)
	if result.Error != nil {
		return nil, fmt.Errorf("error in find all portfolios: %w", result.Error)
	}
	portfoliosDTO := []*controllers.PortfolioDTO{}
	for _, p := range portfoliosDB {
		portfoliosDTO = append(portfoliosDTO, toDTO(p))
	}
	return portfoliosDTO, nil
}

func toDTO(portfolioDB models.Portfolio) *controllers.PortfolioDTO {
	return &controllers.PortfolioDTO{
		Id:       int(portfolioDB.ID),
		Title:    portfolioDB.Title,
		Summary:  portfolioDB.Summary,
		Url:      portfolioDB.Url,
		ImageUrl: portfolioDB.ImageUrl,
		ImageAlt: portfolioDB.ImageAlt,
	}
}
