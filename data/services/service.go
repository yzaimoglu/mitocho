package services

import "github.com/yzaimoglu/mitocho/config"

type Service struct {
	DB *config.Database
}

func NewService(db *config.Database) *Service {
	return &Service{DB: db}
}
