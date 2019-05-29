package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-gomail/gomail"
	"net/mail"
)

type Form struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type Payload struct {
	Form Form `json:"data"`
}

type Body struct {
	Payload Payload `json:"payload"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get information from the data payload.
	var body Body
	json.Unmarshal([]byte(request.Body), &body)

	// Get emails
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	host := os.Getenv("MAIL_HOST")
	senderemail := os.Getenv("MAIL_SENDER_EMAIL")
	senderidentity := os.Getenv("MAIL_SENDER_NAME")
	smtppassword := os.Getenv("MAIL_PASSWORD")
	smtpuser := os.Getenv("MAIL_USERNAME")
	myEmail := os.Getenv("MY_EMAIL")

	// Send email
	m := gomail.NewMessage()
	m.SetHeader("From", myEmail)
	m.SetHeader("To", body.Payload.Form.Email)
	m.SetHeader("Subject", fmt.Sprintf("%s, I received your email!", body.Payload.Form.Name))

	m.SetBody("text/plain", body.Payload.Form.Message)

	d := gomail.NewDialer(server, port, smtpuser, smtppassword)

	d.DialAndSend(m)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "success!",
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
