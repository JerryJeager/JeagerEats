package emails

import (
	"fmt"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
)

func RestaurantRiderMail(restaurant *models.Restaurant, rider *models.User, order *models.Order) string {
	return fmt.Sprintf(`
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
      margin: 20px 0;
    }
    .button {
      text-align: center;
      margin: 20px 0;
    }
    .button a {
      background-color: #4CAF50;
      color: white;
      padding: 10px 20px;
      text-decoration: none;
      border-radius: 5px;
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
      <h2>Rider Assigned for Pickup</h2>
    </div>
    <p>Hello %s,</p>
    <p>We’re notifying you that a rider has been assigned to pick up the delivery package for an order placed with your restaurant.</p>
    
    <div class="details">
      <strong>Rider Details:</strong>
      <ul>
        <li><strong>Name:</strong> %s</li>
        <li><strong>Phone:</strong> %s</li>
        <li><strong>Email:</strong> %s</li>
      </ul>
    </div>

    <div class="details">
      <strong>Pickup Information:</strong>
      <ul>
        <li><strong>Address:</strong> %s</li>
        <li><strong>Order Reference:</strong> %s</li>
      </ul>
    </div>

    <p>Please ensure the package is ready for pickup when the rider arrives. Thank you!</p>

    <div class="footer">
      <p>&copy; 2025 JeagerEats</p>
    </div>
  </div>
</body>
</html>

	`, restaurant.Name, rider.FirstName, rider.PhoneNumber, rider.Email, restaurant.Address, order.RefID)
}
