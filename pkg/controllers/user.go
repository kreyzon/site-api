package controllers

type UserController interface {
	Login(username string, password string) (*UserDTO, error)
}

type UserDTO struct {
	Id       int
	Username string
	Fullname string
	Role     string
}
