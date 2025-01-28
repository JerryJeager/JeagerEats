package emails

import (
	"fmt"
	"os"
	// "strconv"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
	"github.com/google/uuid"
)

func SendDeliveryAcceptMail(restaurant *models.Restaurant, order *models.Order, riderID uuid.UUID) string {
	var acceptLink string
	// deliveryFee, _ := strconv.ParseFloat(order.DeliveryFee, 64)
	if os.Getenv("ENVIRONMENT") == "production" {
		acceptLink = fmt.Sprintf("https://jeager-eats.vercel.app/rider/%s/accept/%s", riderID, order.ID)
	} else {
		acceptLink = fmt.Sprintf("http://localhost:3000/rider/%s/accept/%s", riderID, order.ID)
	}
	return fmt.Sprintf(`
		<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>New Delivery Request</title>
  <style>
    body {
      font-family: Arial, sans-serif;
      background-color: #f4f4f4;
      margin: 0;
      padding: 20px;
    }
    .container {
      max-width: 600px;
      background: #ffffff;
      padding: 20px;
      border-radius: 10px;
      box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
      margin: auto;
    }
    .header {
      font-size: 20px;
      font-weight: bold;
      text-align: center;
      margin-bottom: 20px;
    }
    .details {
      font-size: 16px;
      margin-bottom: 20px;
      line-height: 1.5;
    }
    .button {
      display: block;
      width: 100%%;
      text-align: center;
      background-color: #28a745;
      color: white;
      padding: 12px;
      font-size: 16px;
      font-weight: bold;
      border-radius: 5px;
      text-decoration: none;
      margin-bottom: 10px;
    }
    .button.decline {
      background-color: #dc3545;
    }
    .footer {
      text-align: center;
      font-size: 14px;
      color: #666;
      margin-top: 20px;
    }
  </style>
</head>
<body>
  <div class="container">
    <div class="header">🚴 New Delivery Request</div>
    <div class="details">
      <p>Hi <strong>JeagerEats Rider</strong>,</p>
      <p>A new delivery request is available for you:</p>
      <p><strong>📍 Pickup Location:</strong> %s, %s</p>
      <p><strong>🏠 Delivery Address:</strong> %s</p>
      <p><strong>💰 Delivery Fee:</strong> ₦%.2f</p>
      <p>If you're available to take this order, click below:</p>
    </div>
    <a href="%s" class="button">✅ Accept Delivery</a>
    <div class="footer">
      Thank you, <br>
      <strong>JeagerEats Team</strong>
    </div>
  </div>
</body>
  `, restaurant.Name, restaurant.Address, order.DeliveryAddress, float64(order.DeliveryFee), acceptLink)
}
