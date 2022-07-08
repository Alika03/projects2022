package usecase

import (
	"back-end/auth"
	"back-end/config"
	"back-end/models"
	"back-end/utils"
	"context"
	"crypto/rsa"
	"database/sql"
	"github.com/dgrijalva/jwt-go/v4"
	"strconv"
	"time"
)

const tokenType = "Bearer"

type UseCase struct {
	userRepo auth.UserRepository
	jwtRepo  auth.JwtRepository
	key      *rsa.PrivateKey
}

func NewAuthUseCase(userRepo auth.UserRepository, jwtRepo auth.JwtRepository, key *rsa.PrivateKey) *UseCase {
	return &UseCase{
		userRepo: userRepo,
		jwtRepo:  jwtRepo,
		key:      key,
	}
}

func (u *UseCase) SignUp(ctx context.Context, username, password string) error {
	var model = &models.User{
		Id:       utils.GenerateUuid().String(),
		Username: username,
		Password: password,
	}

	memory, _ := strconv.Atoi(config.GetConfig().HashParams.Memory)
	parallelism, _ := strconv.Atoi(config.GetConfig().HashParams.Parallelism)
	iterations, _ := strconv.Atoi(config.GetConfig().HashParams.Iterations)
	saltLength, _ := strconv.Atoi(config.GetConfig().HashParams.SaltLength)
	keyLength, _ := strconv.Atoi(config.GetConfig().HashParams.KeyLength)

	hashModel := utils.NewHashPassword(uint32(memory), uint32(iterations), uint8(parallelism), uint32(saltLength), uint32(keyLength))

	hashPassword, err := hashModel.GenerateHashPassword(model.Password)
	if err != nil {
		return err
	}

	model.HashPassword = hashPassword

	return u.userRepo.CreateUser(ctx, model)
}

func (u *UseCase) SignIn(ctx context.Context, username, password string) (*models.Tokens, error) {
	userModel, err := u.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	isMatched, err := utils.ComparePasswordHash(password, userModel.HashPassword)
	if err != nil {
		return nil, err
	}
	if !isMatched {
		return nil, auth.ErrInvalidPassword
	}

	accessTokenModel := &models.AccessToken{
		Id:        utils.GenerateUuid().String(),
		UserId:    userModel.Id,
		ExpiredAt: time.Now().Add(time.Minute),
	}

	accessClaims := &jwt.StandardClaims{
		ExpiresAt: &jwt.Time{Time: accessTokenModel.ExpiredAt},
		Issuer:    accessTokenModel.UserId,
		ID:        accessTokenModel.Id,
	}
	accessToken, err := utils.NewJwt(u.key, accessClaims).CreateToken()
	if err != nil {
		return nil, err
	}

	refreshTokenModel := &models.RefreshToken{
		Id:            utils.GenerateUuid().String(),
		AccessTokenId: accessTokenModel.Id,
		ExpiredAt:     time.Now().Add(time.Minute),
	}

	refreshClaims := &jwt.StandardClaims{
		ExpiresAt: &jwt.Time{Time: refreshTokenModel.ExpiredAt},
		Issuer:    accessTokenModel.UserId,
		ID:        refreshTokenModel.Id,
	}

	refreshToken, err := utils.NewJwt(u.key, refreshClaims).CreateToken()
	if err != nil {
		return nil, err
	}

	if err = u.jwtRepo.AddAccessToken(ctx, &sql.Tx{}, accessTokenModel); err != nil {
		return nil, err
	}

	if err = u.jwtRepo.AddRefreshToken(ctx, &sql.Tx{}, refreshTokenModel); err != nil {
		return nil, err
	}

	return &models.Tokens{
		TokenType:    tokenType,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u *UseCase) VerifyAccessToken(ctx context.Context, accessToken string) (*models.User, error) {
	jwToken := utils.NewJwt(u.key, nil)

	jwTokenModel, err := jwToken.ParseToken(accessToken)
	if err != nil {
		return nil, err
	}

	stdClaim := jwTokenModel.Claims.(jwt.MapClaims)
	userId := stdClaim["iss"].(string)

	isExisted, err := u.jwtRepo.HasAccessTokenById(ctx, userId)
	if err != nil {
		return nil, err
	}

	if !isExisted {
		return nil, auth.ErrUnauthorized
	}

	return u.userRepo.GetById(ctx, userId)
}
