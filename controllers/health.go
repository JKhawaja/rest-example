package controllers

import (
	"github.com/JKhawaja/rest-example/controllers/app"
	"github.com/goadesign/goa"
)

// HealthController implements the health resource.
type HealthController struct {
	*goa.Controller
}

// NewHealthController creates a health controller.
func NewHealthController(service *goa.Service) *HealthController {
	return &HealthController{Controller: service.NewController("HealthController")}
}

// Healthcheck runs the healthcheck action.
func (c *HealthController) Healthcheck(ctx *app.HealthcheckHealthContext) error {
	return ctx.OK([]byte("200 - OK"))
}
