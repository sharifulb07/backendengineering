package service

import "github.com/sharifulb07/backendengineering.git/testing/integretiontest/projectone/internal/repository"

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService)RegisterUser(name string)error{

	return s.Repo.CreateUser(name)
}