package models

type Person struct {
	Id string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Gender string `json:"gender"`
	Address Address `json:"address"`
	Friends []Friends `json:"friends"`
	CardNumber string `json:"card_number"`
	Birthday string `json:"birthday"`
	Profession string `json:"profession"`
}

type Address struct {
	Street string `json:"city"`
	City string `json:"street"`
}

type Friends struct {
	Id string `json:"id"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type UpdatePerson struct {
	Id string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Gender string `json:"gender"`
	CardNumber string `json:"card_number"`
	Birthday string `json:"birthday"`
	Profession string `json:"profession"`
}

type CreatePerson struct {
	Id string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Gender string `json:"gender"`
	Address Address `json:"address"`
	Friends []Friends `json:"friends"`
	CardNumber string `json:"card_number"`
	Birthday string `json:"birthday"`
	Profession string `json:"profession"`
}