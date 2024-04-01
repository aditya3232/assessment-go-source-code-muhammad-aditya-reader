package route

import (
	"assessment-go-source-code-muhammad-aditya-reader/internal/delivery/http"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type RouteConfig struct {
	App              *fiber.App
	Log              *logrus.Logger
	SellerController *http.SellerController
}

func (c *RouteConfig) Setup() {
	c.App.Use(c.recoverPanic)
	c.SetupGuestRoute()
}

func (c *RouteConfig) recoverPanic(ctx *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("panic occured: %v", r)
			c.Log.WithError(err).Error("Panic occured")
			ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
				"error":   err.Error(),
			})
		}
	}()

	return ctx.Next()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Get("/api/sellers", c.SellerController.List)
	c.App.Post("/api/sellers", c.SellerController.Create)
	c.App.Put("/api/sellers/:sellerId", c.SellerController.Update)
	c.App.Get("/api/sellers/:sellerId", c.SellerController.Get)
	c.App.Delete("/api/sellers/:sellerId", c.SellerController.Delete)
}
