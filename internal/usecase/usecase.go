package usecase

import (
	"test/internal/repository"
	"test/model"
)

type usecase struct {
	repository repository.Repository
}

type Usecase interface {
	Create(ip model.Ip) error
	LotsAccess() ([]model.Ip, error)
	LastAccess() ([]model.Ip, error)
}


func NewUsecase(repository repository.Repository) Usecase {
	return  &usecase{repository:repository}
}

func (u *usecase) Create(ip model.Ip) error {
	return u.repository.Create(ip)
}

func (u *usecase) LotsAccess() ([]model.Ip, error) {
	return u.repository.LotsAccess()
}

func (u *usecase) LastAccess() ([]model.Ip, error) {
	return u.repository.LastAccess()
}