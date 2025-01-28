package emails

import (
	"fmt"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
)

func CustomerRiderInfoMail(customer, rider *models.User, order *models.Order) string {
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
      border-radius: 8px 8px User
    }
    .details {
      margin: 20px 0;
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
      <h2>Delivery Status Update</h2>
    </div>
    <p>Hello %s,</p>
    <p>We’re excited to let you know that a rider has accepted your delivery request. Here are the details of the rider:</p>
    
    <div class="details">
      <strong>Rider Details:</strong>
      <ul>
        <li><strong>Name:</strong> %s</li>
        <li><strong>Phone:</strong> %s</li>
        <li><strong>Email:</strong> %s</li>
      </ul>
    </div>

    <div class="details">
      <strong>Delivery Details:</strong>
      <ul>
        <li><strong>Delivery Address:</strong> %s</li>
        <li><strong>Order Reference:</strong> %s</li>
      </ul>
    </div>

    <p>The rider will contact you shortly if needed. You can also contact them directly using the details above. We hope you enjoy your meal!</p>

    <div class="footer">
      <p>&copy; 2025 JeagerEats</p>
    </div>
  </div>
</body>
</html>

`,
	customer.FirstName, rider.FirstName, rider.PhoneNumber, rider.Email, order.DeliveryAddress, order.RefID)
}
