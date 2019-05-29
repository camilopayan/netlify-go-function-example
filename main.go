package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Success!",
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
