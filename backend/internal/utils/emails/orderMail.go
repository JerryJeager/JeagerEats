package mails

import (
	"fmt"

	"github.com/JerryJeager/JeagerEats/internal/service/models"
)

func SendOrderSummary (orderItems []models.MenuOrderSummary) string {
	// Send order summary email
	emailTemplate := `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
            background-color: #f8f9fa;
        }
        .email-container {
            max-width: 600px;
            margin: 20px auto;
            background-color: #ffffff;
            border-radius: 8px;
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
            overflow: hidden;
        }
        .header {
            background-color: #007bff;
            color: #ffffff;
            text-align: center;
            padding: 20px;
        }
        .header h1 {
            margin: 0;
            font-size: 24px;
        }
        .content {
            padding: 20px;
        }
        .order-table {
            width: 100%;
            border-collapse: collapse;
            margin-top: 20px;
        }
        .order-table th, .order-table td {
            border: 1px solid #ddd;
            padding: 10px;
            text-align: left;
        }
        .order-table th {
            background-color: #007bff;
            color: #ffffff;
        }
        .order-total {
            text-align: right;
            font-weight: bold;
            font-size: 18px;
            margin-top: 10px;
        }
    </style>
</head>
<body>
    <div class="email-container">
        <div class="header">
            <h1>New Order Summary</h1>
        </div>
        <div class="content">
            <p>Dear Customer,</p>
            <p>A new order has been placed! Below is the summary of the order:</p>
            <table class="order-table">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Quantity</th>
                        <th>Price</th>
                    </tr>
                </thead>
                <tbody>
`

	// Add rows for each menu item
	var total float64
	for _, item := range orderItems {
		emailTemplate += fmt.Sprintf(`
                <tr>
                    <td>%s</td>
                    <td>%d</td>
                    <td>$%.2f</td>
                </tr>
`, item.Name, item.Quantity, item.Price)
		total += item.Price * float64(item.Quantity)
	}

	// Add total to the email template
	emailTemplate += fmt.Sprintf(`
                </tbody>
            </table>
            <p class="order-total">Total: $%.2f</p>
            <p>Thank you for chosing JeagerEats</p>
        </div>
    </div>
</body>
</html>
`, total)
	return emailTemplate
}
