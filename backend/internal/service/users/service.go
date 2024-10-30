package users

import (
	"context"
	"errors"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/JerryJeager/JeagerEats/internal/utils"
	"github.com/google/uuid"
)

type UserSv interface {
	CreateUser(ctx context.Context, user *models.User) (string, error)
	Login(ctx context.Context, user *models.UserLogin) (string, string, error)
}

type UserServ struct {
	repo UserStore
}

func NewUserService(repo UserStore) *UserServ {
	return &UserServ{repo: repo}
}

func (s *UserServ) CreateUser(ctx context.Context, user *models.User) (string, error) {
	if !models.IsValidRole(user.Role) {
		return "", errors.New("invalid role")
	}
	restaurant := models.Restaurant{}
	rider := models.Rider{}
	id := uuid.New()
	user.ID = id
	if user.Role == models.VENDOR {
		restaurant.ID = uuid.New()
		restaurant.UserID = id
		restaurant.Name = user.FirstName + "Restaurant"
	} else {
		rider.ID = uuid.New()
		rider.UserID = id
	}
	if err := user.HashPassword(); err != nil {
		return "", err
	}
	return id.String(), s.repo.CreateUser(ctx, user, &restaurant, &rider)
}

func (s *UserServ) Login(ctx context.Context, user *models.UserLogin) (string, string, error) {
	u, err := s.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return "", "", err
	}
	if err := models.VerifyPassword(user.Password, u.Password); err != nil {
		return "","", err
	}
	token, err := utils.GenerateToken(u.ID, u.Role)
	if err != nil {
		return "", "", err
	}
	return u.ID.String(), token, nil
}
