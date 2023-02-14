package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"app/models"
)

type personRepo struct {
	fileName string
	file *os.File
}

func NewPersonRepo(fileName string, file *os.File) *personRepo {
	return &personRepo{
		fileName: fileName,
		file: file,
	}
}

func (p *personRepo) Update(req *models.UpdatePerson) error {
	var people []*models.Person
	err := json.NewDecoder(p.file).Decode(&people)
	if err != nil {
		return err
	}

	flag := false
	for in, val := range people {
		if val.Id == req.Id {
			people[in].FirstName = req.FirstName
			people[in].LastName = req.LastName
			people[in].Gender = req.Gender
			people[in].CardNumber = req.CardNumber
			people[in].Birthday = req.Birthday
			people[in].Profession = req.Profession
			flag = true
			fmt.Println("SUCCESS: Updated")
		}
	}

	if flag == false {
		return errors.New("There is no id with this value")
	}

	body, err := json.MarshalIndent(people, "", " ")
	err = ioutil.WriteFile(p.fileName, body, os.ModePerm)

	if err != nil {
		return err
	}
	return nil
}

func (p *personRepo) Create(req *models.CreatePerson) error {
	data, err := ioutil.ReadFile(p.fileName)
	if err != nil {
		return err 
	}
	var people []models.Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		return err
	}

	for _, val := range people {
		if val.Id == req.Id {
			return errors.New("WARNING: This ID already exists")
		}
	}
	people = append(people, models.Person{
		Id: req.Id,
		FirstName: req.FirstName,
		LastName: req.LastName,
		Gender: req.Gender,
		Address: req.Address,
		Friends: req.Friends,
		CardNumber: req.CardNumber,
		Birthday: req.Birthday,
		Profession: req.Profession,
	})
	
	body, err := json.MarshalIndent(people, "", " ")
	err = ioutil.WriteFile(p.fileName, body, os.ModePerm)

	if err != nil {
		return err
	}
	return nil
}

func (p *personRepo) GetById(req *models.UserPrimaryKey) (models.Person, error) {
	data, err := ioutil.ReadFile(p.fileName)
	if err != nil {
		return models.Person{}, err
	}
	var people []models.Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		return models.Person{}, err
	}
	for _, val := range people {
		if req.Id == val.Id {
			return val, nil
		}
	}
	return models.Person{}, errors.New("There is no person with this id")
}

func (p *personRepo) GetAll() (models.GetListResponse, error) {
	data, err := ioutil.ReadFile(p.fileName)
	if err != nil {
		return models.GetListResponse{}, err
	}
	var people []models.Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		return models.GetListResponse{}, err
	}
	return models.GetListResponse{
		Count: len(people),
		People: people,
	}, nil
}

func (p *personRepo) Delete(req *models.UserPrimaryKey) error {
	data, err := ioutil.ReadFile(p.fileName)
	if err != nil {
		return err
	}
	var people []*models.Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		return err
	}
	flag := false
	for in, val := range people {
		if val.Id == req.Id {
			people = append(people[:in], people[in+1:]...)
			flag = true
		}
	}

	if flag == false {
		return errors.New("There is no person with this id")
	}

	body, err := json.MarshalIndent(people, "", " ")
	err = ioutil.WriteFile(p.fileName, body, os.ModePerm)

	if err != nil {
		return err
	}
	return nil
}