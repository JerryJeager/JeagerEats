package restaurants


type RestaurantSv interface {
}

type RestaurantServ struct {
	store RestaurantStore
}

func NewRestaurantService(store RestaurantStore) *RestaurantServ {
	return &RestaurantServ{store: store}
}
