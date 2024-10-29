package riders

import "gorm.io/gorm"

type RiderStore interface {
}

type RiderRepo struct {
	client *gorm.DB
}

func NewRiderRepo(client *gorm.DB) *RiderRepo {
	return &RiderRepo{client: client}
}
