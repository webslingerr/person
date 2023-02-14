package storage

import (
	"app/config"
	"os"
)

type Store struct {
	Person *personRepo
}

func NewFileJson(cfg *config.Config) (*Store, error) {
	personFile, err := os.Open(cfg.Path + cfg.PersonFileName)
	if err != nil {
		return nil, err
	}

	return &Store{
		Person: NewPersonRepo(cfg.Path+cfg.PersonFileName, personFile),
	}, nil
}