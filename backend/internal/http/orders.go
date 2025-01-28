package http

import (
	"net/http"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/JerryJeager/JeagerEats/internal/service/orders"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderController struct {
	serv orders.OrderSv
}

func NewOrderController(serv orders.OrderSv) *OrderController {
	return &OrderController{serv: serv}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	user, err := GetUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	var order models.OrderCreate
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order.UserID = user.UserID

	id, err := c.serv.CreateOrder(ctx, &order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": id})

}

func (c *OrderController) AcceptOrder(ctx *gin.Context) {
	var orderID OrderIDPathParam
	if err := ctx.ShouldBindUri(&orderID); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	var orderRiderUpdate models.OrderRiderUpdate
	if err := ctx.ShouldBindJSON(&orderRiderUpdate); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	id := orderID.ID

	if err := c.serv.UpdateOrderRider(ctx, uuid.MustParse(id), &orderRiderUpdate); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusOK)
}
