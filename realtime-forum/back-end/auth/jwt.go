package auth

type Jwt interface {
	CreateToken() error
	CheckToken() error
}
