package http


import "github.com/JerryJeager/JeagerEats/internal/service/riders"

type RiderController struct {
	serv riders.RiderSv
}

func NewRiderController(serv riders.RiderSv) *RiderController {
	return &RiderController{serv: serv}
}
