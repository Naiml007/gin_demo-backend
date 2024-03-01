// routes/mail.go
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

// SetupMailRoute initializes the mail route
func SetupMailRoute(r *gin.Engine) {
	r.GET("/mail", func(c *gin.Context) {
		htmlContent := `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Main website</title>
			</head>
			<body>
				<h1>Main website</h1>
				<h2>Mail Route</h2>
				<form id="mailForm">
					<label for="to">To:</label>
					<input type="email" id="to" name="to" required><br>

					<label for="subject">Subject:</label>
					<input type="text" id="subject" name="subject" required><br>

					<label for="message">Message:</label>
					<textarea id="message" name="message" required></textarea><br>

					<input type="button" value="Send Mail" onclick="sendMail()">
				</form>

				<script>
					function sendMail() {
						var to = document.getElementById('to').value;
						var subject = document.getElementById('subject').value;
						var message = document.getElementById('message').value;

						// Send data to the server for email sending
						fetch('/sendmail', {
							method: 'POST',
							headers: {
								'Content-Type': 'application/json',
							},
							body: JSON.stringify({ to: to, subject: subject, message: message }),
						})
						.then(response => response.json())
						.then(data => {
							alert('Email sent successfully!');
						})
						.catch(error => {
							alert('Error sending email: ' + error.message);
						});
					}
				</script>
			</body>
			</html>
		`

		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	})

	// Define the /sendmail endpoint inside the function block
	r.POST("/sendmail", func(c *gin.Context) {
		var requestData map[string]string
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		to := requestData["to"]
		subject := requestData["subject"]
		message := requestData["message"]

		// Send email using the SendMail function
		err := SendMail(to, subject, message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
	})
}

// Modify this part according to your SMTP server configuration
// https://app.brevo.com/ for smtp server configuration
const (
	smtpServer = "smtp-relay.brevo.com"
	smtpPort   = 587
	smtpUser   = "fromBrevo@gmail.com"
	smtpPass   = ""
)

// SendMail sends an email using gomail library
func SendMail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", smtpUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(smtpServer, smtpPort, smtpUser, smtpPass)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
