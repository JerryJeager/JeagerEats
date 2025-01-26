package orders

import (
	"context"
	"fmt"
	"os"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	mails "github.com/JerryJeager/JeagerEats/internal/utils/emails"
	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
)

type OrderSv interface {
	CreateOrder(ctx context.Context, order *models.OrderCreate) (string, error)
}

type OrderServ struct {
	repo OrderStore
}

func NewOrderService(repo OrderStore) *OrderServ {
	return &OrderServ{repo: repo}
}

func (s *OrderServ) CreateOrder(ctx context.Context, order *models.OrderCreate) (string, error) {
	var items []models.Item
	var orderDetails models.Order
	var menuOrderSum []models.MenuOrderSummary
	id := uuid.New()
	orderDetails.ID = id
	orderDetails.UserID = order.UserID
	orderDetails.RestaurantID = order.RestaurantID
	orderDetails.TotalPrice = order.TotalPrice

	for _, item := range order.Items {
		item.ID = uuid.New()
		item.OrderID = id
		items = append(items, item)
		if menu, err := s.repo.GetMenu(ctx, item.MenuID); err != nil {
			return "", err
		} else {
			menuOrderSum = append(menuOrderSum, models.MenuOrderSummary{Name: menu.Name, Quantity: item.Quantity, Price: item.PricePerItem})
		}

	}
	err := s.repo.CreateOrder(ctx, &orderDetails, &items)
	if err != nil {
		return "", err
	}

	// Send email to user
	user, err := s.repo.GetUser(ctx, order.UserID)
	if err != nil{
		fmt.Println("failed to get user")
	}

	err = sendEmail(user.Email, &menuOrderSum)
	if err != nil {
		fmt.Print(err)
	}
	return id.String(), nil
}

func sendEmail(userEmail string, menuSummary *[]models.MenuOrderSummary) error {

	email := os.Getenv("EMAIL")
	emailUsername := os.Getenv("EMAILUSERNAME")
	emailPassword := os.Getenv("EMAILPASSWORD")
	m := gomail.NewMessage()
	m.SetAddressHeader("From", email, "JeagerEats")
	m.SetHeader("To", userEmail)
	m.SetAddressHeader("Cc", userEmail, "JeagerEats")
	m.SetHeader("Subject", "Your Order Has Been Placed")

	m.SetBody("text/html", mails.SendOrderSummary(*menuSummary))

	d := gomail.NewDialer("smtp.gmail.com", 465, emailUsername, emailPassword)

	// Send the email to user
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
