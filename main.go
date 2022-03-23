package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadaptor "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/oskarincon/operation-quasar-go/constants"
	"github.com/oskarincon/operation-quasar-go/handlers"
	"github.com/oskarincon/operation-quasar-go/services"
)

var adapter *fiberadaptor.FiberLambda

func init() {
	app := fiber.New()
	router := app.Group(constants.Apipath)
	// Create Routes
	handlers.SetupRoutes(router)
	// Server Init
	services.Init()
	//handlers.Init(app)
	adapter = fiberadaptor.New(app)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return adapter.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
