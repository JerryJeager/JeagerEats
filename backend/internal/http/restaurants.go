package http

import "github.com/JerryJeager/JeagerEats/internal/service/restaurants"

type RestaurantController struct {
	serv restaurants.RestaurantSv
}

func NewRestaurantController(serv restaurants.RestaurantSv) *RestaurantController {
	return &RestaurantController{serv: serv}
}
