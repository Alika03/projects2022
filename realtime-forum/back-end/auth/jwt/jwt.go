package jwt

type Jwt struct{}

func NewJwt() *Jwt {

	return &Jwt{}
}

func (j *Jwt) CreateToken(userId string) error {

	return nil
}

func (j *Jwt) CheckToken() error {
	return nil
}
