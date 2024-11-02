package gateways

import (
	httpclient "fiber-v3-template/src/infra/http_client"
	"fiber-v3-template/src/services"

	"github.com/gofiber/fiber/v3"
)

type HTTPGateway struct {
	NumberService  services.INumberService
	CarbonProvider httpclient.ICarbonProvider
}

func NewHTTPGateway(app *fiber.App, numberService services.INumberService, carbonProv httpclient.ICarbonProvider) {
	gateway := &HTTPGateway{
		NumberService:  numberService,
		CarbonProvider: carbonProv,
	}

	RouteMain(gateway, app)
	RouteCarbon(gateway, app)
}
