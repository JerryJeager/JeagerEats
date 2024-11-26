package users

import (
	"context"
	"errors"
	"strings"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/JerryJeager/JeagerEats/internal/service/restaurants"
	"github.com/JerryJeager/JeagerEats/internal/utils"
	"github.com/google/uuid"
)

type UserSv interface {
	CreateUser(ctx context.Context, user *models.User) (string, error)
	Login(ctx context.Context, user *models.UserLogin) (string, string, string, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*models.User, error)
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
	} else if user.Role == models.RIDER {
		rider.ID = uuid.New()
		rider.UserID = id
	}
	if err := user.HashPassword(); err != nil {
		return "", err
	}
	return id.String(), s.repo.CreateUser(ctx, user, &restaurant, &rider)
}

func (s *UserServ) Login(ctx context.Context, user *models.UserLogin) (string, string, string, error) {
	user.Email = strings.TrimSpace(user.Email)
    user.Password = strings.TrimSpace(user.Password)
	u, err := s.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return "", "", "", errors.New("debugging: error came from get email in service")
	}
	var restaurant *models.Restaurant = &models.Restaurant{}
	if u.Role == models.VENDOR {
		restaurant, err = restaurants.GetRestaurant(ctx, u.ID)
		if err != nil {
			return "", "", "", errors.New("debugging: error came from get restaurant in service")
		}
	}
	if err := models.VerifyPassword(user.Password, u.Password); err != nil {
		return "", "", "", errors.New("debuggin: error came from verify password in service")
	}
	token, err := utils.GenerateToken(u.ID, &restaurant.ID, u.Role)
	if err != nil {
		return "", "", "", errors.New("debugging: error came from genreate token in service")
	}
	return u.ID.String(), token, u.Role, nil
}

func (s *UserServ) GetUser(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	return s.repo.GetUser(ctx, userID)
}
