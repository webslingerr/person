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

func (c *Controller) GetPersonById(req *models.UserPrimaryKey) (models.Person, error) {
	person, err := c.store.Person.GetById(req)
	if err != nil {
		return models.Person{}, err
	}
	return person, nil
}

func (c *Controller) GetAllPeople() (models.GetListResponse, error) {
	res, err := c.store.Person.GetAll()
	if err != nil {
		return models.GetListResponse{}, err
	}
	return res, err
}

func (c *Controller) DeletePerson(req *models.UserPrimaryKey) error {
	err := c.store.Person.Delete(req)
	if err != nil {
		return err
	}
	return nil
}