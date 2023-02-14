package controller

import (
	"app/models"
)

func (c *Controller) UpdatePerson(req *models.UpdatePerson) error {
	err := c.store.Person.Update(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) CreatePerson(req *models.CreatePerson) error {
	err := c.store.Person.Create(req)
	if err != nil {
		return err
	}
	return nil
}