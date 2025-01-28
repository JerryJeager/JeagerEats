package orders

import (
	"context"
	"errors"
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
	UpdateOrderRider(ctx context.Context, orderID uuid.UUID, orderRider *models.OrderRiderUpdate) error
	UpdateOrderStatus(ctx context.Context, orderStatusUpdate *models.OrderStatusUpdate, orderID uuid.UUID) error
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

func (s *OrderServ) UpdateOrderRider(ctx context.Context, orderID uuid.UUID, orderRider *models.OrderRiderUpdate) error {
	order, err := s.repo.GetOrder(ctx, orderID)
	if err != nil {
		return err
	}
	if order.Status == models.ACCEPTED {
		return errors.New("order has already been assigned to a rider")
	}

	restaurant, err := s.repo.GetRestaurant(ctx, order.RestaurantID)
	if err != nil {
		return err
	}

	restaurantMail, err := s.repo.RestaurantOwnerMail(ctx, order.RestaurantID)
	if err != nil {
		return err
	}

	customer, err := s.repo.GetUser(ctx, order.UserID)
	if err != nil {
		return err
	}

	rider, err := s.repo.GetUser(ctx, orderRider.RiderID)
	if err != nil {
		return err
	}

	err = s.repo.UpdateOrderRider(ctx, orderID, orderRider)
	if err != nil {
		fmt.Println("error from update rider in store")
		return err
	}

	if err := sendDeliveryPickupConfirmation(rider, customer, order, restaurant); err != nil {
		fmt.Println(err)
	}
	if err := sendCustomerRiderInfo(customer, rider, order); err != nil {
		fmt.Println(err)
	}
	if err := sendRestaurantRiderMail(restaurant, rider, order, restaurantMail); err != nil {
		fmt.Println(err)
	}

	return nil
}

func (s *OrderServ) UpdateOrderStatus(ctx context.Context, orderStatusUpdate *models.OrderStatusUpdate, orderID uuid.UUID) error {
	order, err := s.repo.GetOrder(ctx, orderID)
	if err != nil {
		return nil
	}

	customer, err := s.repo.GetUser(ctx, order.UserID)
	if err != nil {
		return err
	}

	if !models.Statuses[orderStatusUpdate.Status] {
		return errors.New("invalid status used")
	}

	if err := s.repo.UpdateOrderStatus(ctx, orderStatusUpdate, orderID); err != nil {
		return err
	}

	if orderStatusUpdate.Status == models.INTRANSIT {
		if err := sendInTransitMail(customer, order); err != nil {
			fmt.Println(err)
		}
	} else if orderStatusUpdate.Status == models.DELIVERED {
		if err := sendDeliveryCompleteMail(customer, order); err != nil {
			fmt.Println(err)
		}
	}
	return nil
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
		m.SetHeader("To", rider.Email)
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
func sendDeliveryPickupConfirmation(rider, customer *models.User, order *models.Order, restaurant *models.Restaurant) error {

	email := os.Getenv("EMAIL")
	emailUsername := os.Getenv("EMAILUSERNAME")
	emailPassword := os.Getenv("EMAILPASSWORD")
	m := gomail.NewMessage()
	m.SetAddressHeader("From", email, "JeagerEats")
	m.SetHeader("To", rider.Email)
	m.SetAddressHeader("Cc", rider.Email, "JeagerEats")
	m.SetHeader("Subject", "Delivery")

	m.SetBody("text/html", emails.DeliveryPickupConfirmation(rider, customer, order, restaurant))

	d := gomail.NewDialer("smtp.gmail.com", 465, emailUsername, emailPassword)

	// Send the email to user
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
func sendCustomerRiderInfo(customer, rider *models.User, order *models.Order) error {

	email := os.Getenv("EMAIL")
	emailUsername := os.Getenv("EMAILUSERNAME")
	emailPassword := os.Getenv("EMAILPASSWORD")
	m := gomail.NewMessage()
	m.SetAddressHeader("From", email, "JeagerEats")
	m.SetHeader("To", customer.Email)
	m.SetAddressHeader("Cc", customer.Email, "JeagerEats")
	m.SetHeader("Subject", "Delivery Status Update")

	m.SetBody("text/html", emails.CustomerRiderInfoMail(customer, rider, order))

	d := gomail.NewDialer("smtp.gmail.com", 465, emailUsername, emailPassword)

	// Send the email to user
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
func sendRestaurantRiderMail(restaurant *models.Restaurant, rider *models.User, order *models.Order, restaurantOwnerMail *models.RestaurantOwnerMail) error {

	email := os.Getenv("EMAIL")
	emailUsername := os.Getenv("EMAILUSERNAME")
	emailPassword := os.Getenv("EMAILPASSWORD")
	m := gomail.NewMessage()
	m.SetAddressHeader("From", email, "JeagerEats")
	m.SetHeader("To", restaurantOwnerMail.Email)
	m.SetAddressHeader("Cc", restaurantOwnerMail.Email, "JeagerEats")
	m.SetHeader("Subject", "Rider Assigned for Pickup")

	m.SetBody("text/html", emails.RestaurantRiderMail(restaurant, rider, order))

	d := gomail.NewDialer("smtp.gmail.com", 465, emailUsername, emailPassword)

	// Send the email to user
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func sendInTransitMail(customer *models.User, order *models.Order) error {

	email := os.Getenv("EMAIL")
	emailUsername := os.Getenv("EMAILUSERNAME")
	emailPassword := os.Getenv("EMAILPASSWORD")
	m := gomail.NewMessage()
	m.SetAddressHeader("From", email, "JeagerEats")
	m.SetHeader("To", customer.Email)
	m.SetAddressHeader("Cc", customer.Email, "JeagerEats")
	m.SetHeader("Subject", "Your Order is On Its Way!")

	m.SetBody("text/html", emails.TransitMail(customer, order))

	d := gomail.NewDialer("smtp.gmail.com", 465, emailUsername, emailPassword)

	// Send the email to user
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
func sendDeliveryCompleteMail(customer *models.User, order *models.Order) error {

	email := os.Getenv("EMAIL")
	emailUsername := os.Getenv("EMAILUSERNAME")
	emailPassword := os.Getenv("EMAILPASSWORD")
	m := gomail.NewMessage()
	m.SetAddressHeader("From", email, "JeagerEats")
	m.SetHeader("To", customer.Email)
	m.SetAddressHeader("Cc", customer.Email, "JeagerEats")
	m.SetHeader("Subject", "Your Order Has Been Delivered!")

	m.SetBody("text/html", emails.DoneDeliveryMail(customer, order))

	d := gomail.NewDialer("smtp.gmail.com", 465, emailUsername, emailPassword)

	// Send the email to user
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
