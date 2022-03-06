package controllers

type PortfolioController interface {
	Create(portfolio PortfolioDTO) (int, error)
	ReadAll() ([]*PortfolioDTO, error)
}

type PortfolioDTO struct {
	Id       int
	Title    string
	Summary  string
	Url      string
	ImageUrl string
	ImageAlt string
}
