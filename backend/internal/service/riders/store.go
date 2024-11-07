package riders

import (
	"context"
	"errors"

	"github.com/JerryJeager/JeagerEats/config"
	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RiderStore interface {
	UpdateRider(ctx context.Context, userID uuid.UUID, rider *models.RiderUpdate) error
	UpdateRiderProfileImg(ctx context.Context, userID uuid.UUID, filePath string) error
}

type RiderRepo struct {
	client *gorm.DB
}

func NewRiderRepo(client *gorm.DB) *RiderRepo {
	return &RiderRepo{client: client}
}

func (r *RiderRepo) UpdateRider(ctx context.Context, userID uuid.UUID, rider *models.RiderUpdate) error {
	qry := config.Session.WithContext(ctx).Model(&models.Rider{}).Where("user_id = ?", userID).Updates(rider)
	if qry.RowsAffected == 0 {
		return errors.New("rider not found")
	}
	return nil
}

func (r *RiderRepo) UpdateRiderProfileImg(ctx context.Context, userID uuid.UUID, filePath string) error {
	return config.Session.WithContext(ctx).Model(&models.Rider{}).Where("user_id = ?", userID).Update("profile_img", filePath).Error
}