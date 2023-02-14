package controller

import (
	"app/config"
	"app/storage"
)

type Controller struct {
	store *storage.Store
	cfg *config.Config
}

func NewController(cfg *config.Config, store *storage.Store) *Controller {
	return &Controller{
		cfg: cfg,
		store: store,
	}
}	