package riders

import (
	"context"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/google/uuid"
)

type RiderSv interface {
	UpdateRider(ctx context.Context, userID uuid.UUID, rider *models.RiderUpdate) error
}

type RiderServ struct {
	store RiderStore
}

func NewRiderService(store RiderStore) *RiderServ {
	return &RiderServ{store: store}
}

func (r *RiderServ) UpdateRider(ctx context.Context, userID uuid.UUID, rider *models.RiderUpdate) error {
	return r.store.UpdateRider(ctx, userID, rider)
}
