package services

import (
	"admin-backend/models"
	"admin-backend/repository"
	"admin-backend/utils"
	"context"
	"time"
)

type userService struct {
	ur             repository.UserRepository
	contextTimeout time.Duration
}

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	CreateAccessToken(user *models.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *models.User, secret string, expiry int) (refreshToken string, err error)
	ExtractIDFromToken(requestToken string, secret string) (string, error)
	GetUserByID(ctx context.Context, id int) (*models.User, error)
}

func NewService(ur repository.UserRepository, contextTimeout time.Duration) *userService {
	return &userService{
		ur:             ur,
		contextTimeout: contextTimeout,
	}
}

func (s *userService) CreateUser(ctx context.Context, user *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	return s.ur.Create(ctx, user)
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	return s.ur.GetUserByEmail(ctx, email)
}

func (su *userService) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()
	return su.ur.GetUserByID(ctx, id)
}

func (su *userService) CreateAccessToken(user *models.User, secret string, expiry int) (accessToken string, err error) {
	return utils.CreateAccessToken(user, secret, expiry)
}

func (su *userService) CreateRefreshToken(user *models.User, secret string, expiry int) (refreshToken string, err error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}

func (rtu *userService) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return utils.ExtractIDFromToken(requestToken, secret)
}
