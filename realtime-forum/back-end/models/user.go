package models

type User struct {
	Id           string
	Username     string
	Password     string
	HashPassword string
}

type Token struct {
}
