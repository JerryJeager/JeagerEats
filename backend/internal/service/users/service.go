package users

import (
	"context"
	"errors"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/google/uuid"
)

type UserSv interface {
	CreateUser(ctx context.Context, user *models.User) (string, error)
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
