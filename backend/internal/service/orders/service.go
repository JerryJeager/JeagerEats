package orders

import (
	"context"
	"fmt"
	"os"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/JerryJeager/JeagerEats/internal/utils"
	"github.com/JerryJeager/JeagerEats/internal/utils/emails"
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
	orderDetails.Status = models.PENDING
	orderDetails.DeliveryAddress = order.DeliveryAddress
	orderDetails.DeliveryFee = order.DeliveryFee 
	orderDetails.RefID = utils.GenerateCode()

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
	riders, err := s.repo.GetRiders(ctx)
	if err != nil {
		return "", err
	}
	restaurant, err := s.repo.GetRestaurant(ctx, order.RestaurantID)
	if err != nil {
		return "", err
	}
	err = s.repo.CreateOrder(ctx, &orderDetails, &items)
	if err != nil {
		return "", err
	}

	// Send email to user
	user, err := s.repo.GetUser(ctx, order.UserID)
	if err != nil {
		fmt.Println("failed to get user")
	}

	err = sendCustomerOrderEmail(user.Email, &menuOrderSum)
	if err != nil {
		fmt.Print(err)
	}

	//Send email to riders
	if len(*riders) > 0 {
		err = sendRiderOrderEmail(riders, restaurant, &orderDetails)
		if err != nil {
			fmt.Print(err)
		}
	}
	return id.String(), nil
}

func sendCustomerOrderEmail(userEmail string, menuSummary *[]models.MenuOrderSummary) error {

	email := os.Getenv("EMAIL")
	emailUsername := os.Getenv("EMAILUSERNAME")
	emailPassword := os.Getenv("EMAILPASSWORD")
	m := gomail.NewMessage()
	m.SetAddressHeader("From", email, "JeagerEats")
	m.SetHeader("To", userEmail)
	m.SetAddressHeader("Cc", userEmail, "JeagerEats")
	m.SetHeader("Subject", "Your Order Has Been Placed")

	m.SetBody("text/html", emails.SendOrderSummary(*menuSummary))

	d := gomail.NewDialer("smtp.gmail.com", 465, emailUsername, emailPassword)

	// Send the email to user
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func sendRiderOrderEmail(riders *[]models.User, restaurant *models.Restaurant, order *models.Order) error {

	for _, rider := range *riders {
		email := os.Getenv("EMAIL")
		emailUsername := os.Getenv("EMAILUSERNAME")
		emailPassword := os.Getenv("EMAILPASSWORD")
		m := gomail.NewMessage()
		m.SetAddressHeader("From", email, "JeagerEats")
		m.SetHeader("To")
		m.SetAddressHeader("Cc", rider.Email, "JeagerEats")
		m.SetHeader("Subject", "New Delivery Request")

		m.SetBody("text/html", emails.SendDeliveryAcceptMail(restaurant, order, rider.ID))

		d := gomail.NewDialer("smtp.gmail.com", 465, emailUsername, emailPassword)

		// Send the email to user
		if err := d.DialAndSend(m); err != nil {
			return err
		}
	}
	return nil
}
