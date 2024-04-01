package config

import (
	"assessment-go-source-code-muhammad-aditya-reader/internal/delivery/http"
	"assessment-go-source-code-muhammad-aditya-reader/internal/delivery/http/route"
	"assessment-go-source-code-muhammad-aditya-reader/internal/repository"
	"assessment-go-source-code-muhammad-aditya-reader/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	sellerRepository := repository.NewSellerRepository(config.Log)

	sellerUseCase := usecase.NewSellerUseCase(config.DB, config.Log, config.Validate, sellerRepository)

	sellerController := http.NewSellerController(sellerUseCase, config.Log)

	routeConfig := route.RouteConfig{
		App:              config.App,
		SellerController: sellerController,
	}
	routeConfig.Setup()

}
