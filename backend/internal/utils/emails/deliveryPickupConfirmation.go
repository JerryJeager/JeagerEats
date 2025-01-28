package emails

import (
	"fmt"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
)

func DeliveryPickupConfirmation(rider, customer *models.User, order *models.Order, restaurant *models.Restaurant) string {
	return fmt.Sprintf(
		`
			<!DOCTYPE html>
<html>
<head>
  <style>
    body {
      font-family: Arial, sans-serif;
      line-height: 1.6;
      color: #333333;
      margin: 0;
      padding: 0;
    }
    .container {
      max-width: 600px;
      margin: 20px auto;
      padding: 20px;
      border: 1px solid #dddddd;
      border-radius: 8px;
      box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    }
    .header {
      background-color: #4CAF50;
      color: white;
      text-align: center;
      padding: 10px;
      border-radius: 8px 8px 0 0;
    }
    .details {
      margin: 2models
    }
    .footer {
      font-size: 12px;
      color: #777777;
      text-align: center;
      margin-top: 20px;
    }
  </style>
</head>
<body>
  <div class="container">
    <div class="header">
      <h2>Delivery Confirmation</h2>
    </div>
    <p>Hello %s,</p>
    <p>Thank you for confirming that you’ll take on this delivery. Below are the details for the order:</p>
    
    <div class="details">
      <strong>Delivery Details:</strong>
      <ul>
        <li><strong>Customer Name:</strong> %s</li>
        <li><strong>Delivery Address:</strong> %s</li>
        <li><strong>Customer Phone:</strong> %s</li>
        <li><strong>Order Reference:</strong> %s</li>
        <li><strong>Delivery Fee:</strong> ₦%v</li>
      </ul>
    </div>

    <div class="details">
      <strong>Pickup Information:</strong>
      <ul>
        <li><strong>Restaurant Name:</strong> %s</li>
        <li><strong>Pickup Address:</strong> %s</li>
      </ul>
    </div>

    <p>Please ensure you contact the customer if needed for clarification or updates about the delivery. Let us know if you face any issues during the process.</p>

    <p>Safe travels, and thank you for being an important part of JeagerEats!</p>

    <div class="footer">
      <p>&copy; 2025, JeagerEats</p>
    </div>
  </div>
</body>
</html>

		`,
	rider.FirstName, customer.FirstName, order.DeliveryAddress, customer.PhoneNumber, order.RefID, order.DeliveryFee, restaurant.Name, restaurant.Address)
}
