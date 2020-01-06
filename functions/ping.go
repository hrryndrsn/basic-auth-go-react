package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	//must be compatble with the Lamda server response
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "hello world",
	}, nil
}

func main() {
	lambda.Start(handler)
}
