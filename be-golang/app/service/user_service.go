package service

import (
	"be-golang/app/entity"
	"be-golang/app/repository"
	"be-golang/config"
)

type UserService interface {
	CreateUser(user entity.User) (entity.User, error)
	AuthenticateUser(nik, password string) (map[string]interface{}, error)
}

type UserServiceImpl struct {
	repo repository.UserRepository
	cfg  *config.Config
}

func NewUserService(repo repository.UserRepository, cfg *config.Config) UserService {
	return &UserServiceImpl{repo: repo, cfg: cfg}
}

func (s *UserServiceImpl) CreateUser(user entity.User) (entity.User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserServiceImpl) AuthenticateUser(nik, password string) (map[string]interface{}, error) {
	return s.repo.AuthenticateUser(nik, password)
}
