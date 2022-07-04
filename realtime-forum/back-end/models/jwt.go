package models

import "time"

type AccessToken struct {
	Id        string
	UserId    string
	ExpiredAt time.Time
}

type RefreshToken struct {
	Id            string
	AccessTokenId string
	ExpiredAt     time.Time
}

type Tokens struct {
	TokenType    string
	AccessToken  string
	RefreshToken string
}
