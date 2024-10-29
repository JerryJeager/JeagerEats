package users

import (
	"context"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/google/uuid"
)

type UserSv interface {
	CreateUser(ctx context.Context, user *models.User) (uuid.UUID, error)
}

type UserServ struct {
	repo UserStore
}

func NewUserService(repo UserStore) *UserServ {
	return &UserServ{repo: repo}
}

func (s *UserServ) CreateUser(ctx context.Context, user *models.User) (uuid.UUID, error) {
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
	return id, s.repo.CreateUser(ctx, user, &restaurant, &rider)
}
