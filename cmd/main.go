package main

import (
	"app/config"
	"app/models"
	"log"

	// "app/storage"
	"app/controller"
	"app/storage"
)


func main() {
	cfg := config.Load()

	store, err := storage.NewFileJson(&cfg)

	if err != nil {
		panic("error while connect to json file: " + err.Error())
	}

	c := controller.NewController(&cfg, store)
	
	err = c.UpdatePerson(
		&models.UpdatePerson{
			Id: "32a803b5-3f1c-495e-8ce7-4649e4cbe3b1",
			FirstName: "Shokhrukh",
			LastName: "Safarov",
			Gender: "Male",
			CardNumber: "43545fdg45",
			Birthday: "28-05-2003",
			Profession: "Golang Developer",
		},
	)

	if err != nil {
		log.Println(err)
	}

	err = c.CreatePerson(
		&models.CreatePerson{
			Id: "546456543467",
			FirstName: "Asilbek",
			LastName: "Hamijonov",
			Gender: "male",
			Address: models.Address{
				Street: "chilonzor street",
				City: "Tashkent",
			},
			Friends: []models.Friends{
				{
					Id: "43654756745",
					Email: "asfdgg@gmail.com",
					PhoneNumber: "4356-45656-4563",
				},
			},
			CardNumber: "sddf45gfh5454",
			Birthday: "3434-456-323",
			Profession: "Motion Designer",
		},
	)

	if err != nil {
		log.Println(err)
	}
}	