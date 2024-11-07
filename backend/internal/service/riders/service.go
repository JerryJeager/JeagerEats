package riders

import (
	"context"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/google/uuid"
)

type RiderSv interface {
	UpdateRider(ctx context.Context, userID uuid.UUID, rider *models.RiderUpdate) error
	UpdateRiderProfileImg(ctx context.Context, userID uuid.UUID, filePath string) error
}

type RiderServ struct {
	store RiderStore
}

func NewRiderService(store RiderStore) *RiderServ {
	return &RiderServ{store: store}
}

func (s *RiderServ) UpdateRider(ctx context.Context, userID uuid.UUID, rider *models.RiderUpdate) error {
	return s.store.UpdateRider(ctx, userID, rider)
}

func (s *RiderServ) UpdateRiderProfileImg(ctx context.Context, userID uuid.UUID, filePath string) error {
	return s.store.UpdateRiderProfileImg(ctx, userID, filePath)
}
