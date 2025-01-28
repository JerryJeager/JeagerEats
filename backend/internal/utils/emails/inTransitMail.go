package emails

import (
	"fmt"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
)

func TransitMail(customer *models.User, order *models.Order) string {
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
      background-color: #007BFF;
      color: white;
      text-align: center;
      padding: 10px;
      border-radius: 8px 8px 0 0;
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
      <h2>Your Order is On Its Way!</h2>
    </div>
    <p>Hello %s,</p>
    <p>Great news! The rider has picked up your order and is now en route to your delivery address. Below are the details:</p>
    
    <div class="details">
      <strong>Order Details:</strong>
      <ul>
        <li><strong>Order Reference:</strong> %s</li>
        <li><strong>Delivery Address:</strong> %s</li>
      </ul>
    </div>


    <p>Please ensure you're available at the delivery location. Contact the rider if you need to make arrangements.</p>

    <p>Thank you for choosing JeagerEats. Enjoy your meal!</p>

    <div class="footer">
      <p>&copy; 2025 JeagerEats</p>
    </div>
  </div>
</body>
</html>

	`, customer.FirstName, order.RefID, order.DeliveryAddress)
}
