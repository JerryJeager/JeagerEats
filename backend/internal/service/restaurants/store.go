package restaurants


import "gorm.io/gorm"

type RestaurantStore interface {
}

type RestaurantRepo struct {
	client *gorm.DB
}

func NewRestaurantRepo(client *gorm.DB) *RestaurantRepo {
	return &RestaurantRepo{client: client}
}
