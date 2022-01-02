package repository

import (
	"gorm.io/gorm"
	"test/model"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Create(ip model.Ip) error
	LotsAccess() ([]model.Ip, error)
    LastAccess() ([]model.Ip, error)
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db:db}
}

func(r *repository) Create(ip model.Ip) error {
	if err := r.db.Create(&ip).Error; err != nil {
		return err
	}
	return nil
}

func(r *repository) LotsAccess() ([]model.Ip, error) {
	var ips []model.Ip


	return ips, nil
}

func(r *repository) LastAccess() ([]model.Ip, error) {
	var ips []model.Ip
	result := r.db.Order("id desc").Limit(10).Find(&ips)
	if result.Error != nil {
		return nil, result.Error
	}
	return ips, nil
}