package invoice

import (
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
)

func SendEmailToUser(email string, subject string, name string, invoice string, total string) error {

	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		return err
	}

	invoiceHTML := `
	<html>
	<head>
		<meta charset="UTF-8">
		<title>Invoice - Reserve at Glamour Reserve</title>
		<style>
			body {
				font-family: Arial, sans-serif;
			}
			.invoice {
				width: 80%;
				margin: 0 auto;
				border: 1px solid #ccc;
				padding: 20px;
			}
			.invoice-header {
				background-color: #f0f0f0;
				padding: 10px;
			}
			.invoice-header h2 {
				margin: 0;
				color: #333;
			}
			.invoice-details {
				margin-top: 20px;
			}
			.invoice-details p {
				margin: 0;
				color: #333;
			}
			.invoice-total {
				margin-top: 20px;
			}
			.invoice-total p {
				margin: 0;
				color: #333;
			}
			.invoice-thanks {
				margin-top: 20px;
			}
			.invoice-thanks p {
				margin: 0;
				color: #333;
			}
			.invoice-footer {
				margin-top: 20px;
				text-align: center;
				color: #777;
			}
			.invoice-footer p {
				margin: 0;
			}
		</style>
	</head>
	<body>
		<div class="invoice">
			<div class="invoice-header">
				<h2>Invoice - Reserve at Glamour Reserve</h2>
				<h3>Thank you for booking at our salon</h3>
			</div>
			<div class="invoice-details">
			<p>Here is your booking information:</p>
				<p>Name: ` + name + `</p>
				<p>No. Invoice: ` + invoice + `</p>
			</div>
			<div class="invoice-total">
				<p><b>Total Amount: Rp` + total + `</b></p>
			</div>
			<div class="invoice-thanks">
				<p>We hope you enjoy your booking experience at Glamour Reserve. We will do our best to ensure you have an unforgettable beauty experience.</p>
			</div>
			<div class="invoice-thanks">
				<p>Glamour Reserve<br>
				Your beauty, our passion</p>
			</div>
			<div class="invoice-footer">
				<p>All rights reserved &copy; 2023 Glamour Reserve</p>
			</div>
		</div>
	</body>
	</html>`

	m := gomail.NewMessage()
	m.SetHeader("From", smtpUsername)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", invoiceHTML)

	d := gomail.NewDialer(smtpServer, smtpPort, smtpUsername, smtpPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
