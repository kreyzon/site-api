package controllers

type PortfolioController interface {
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
