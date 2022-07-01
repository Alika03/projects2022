package usecase

import (
	"back-end/auth"
	"back-end/config"
	"back-end/models"
	"back-end/utils"
	"context"
	"errors"

	"strconv"
)

type UseCase struct {
	repo auth.UserRepository
}

func NewAuthUseCase(repo auth.UserRepository) *UseCase {
	return &UseCase{repo: repo}
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
	hashModel := utils.HashParams{
		Memory:      uint32(memory),
		Iterations:  uint32(iterations),
		Parallelism: uint8(parallelism),
		SaltLength:  uint32(saltLength),
		KeyLength:   uint32(keyLength),
	}

	hashPassword, err := hashModel.GenerateHashPassword(model.Password)
	if err != nil {
		return err
	}

	model.HashPassword = hashPassword

	return u.repo.CreateUser(ctx, model)
}

func (u *UseCase) SignIn(ctx context.Context, username, password string) error {
	model, err := u.repo.GetByUsername(ctx, username)
	if err != nil {
		return err
	}

	isMatched, err := utils.ComparePasswordHash(password, model.HashPassword)
	if err != nil {
		return err
	}
	if !isMatched {
		return errors.New("incorrect password")
	}

	return nil
}

func (u *UseCase) ParseToken(ctx context.Context, accessToken string) error {
	return nil
}
