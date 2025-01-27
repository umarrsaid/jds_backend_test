package router

import (
	"be-golang/app/handler"
	"be-golang/app/middleware"
	"be-golang/app/repository"
	"be-golang/app/service"
	"be-golang/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitRouters(app *fiber.App, db *gorm.DB, cfg *config.Config) {
	v1 := app.Group("/api/v1")
	userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepository(db, cfg), cfg))
	productHandler := handler.NewProductHandler(service.NewProductService(repository.NewProductRepository(db, cfg), cfg))

	v1.Post("/register", userHandler.CreateUser)
	v1.Post("/login", userHandler.LoginUser)
	v1.Get("/private-claim", middleware.AuthMiddleware(), userHandler.PrivateClaim)

	v1.Get("/product", middleware.AuthMiddleware(), productHandler.GetProducts)
	v1.Get("/product/aggregation", middleware.AuthMiddleware("admin"), productHandler.GetAggregatedProducts)
	v1.Get("/product/private-claim", middleware.AuthMiddleware(), productHandler.PrivateClaim)
}
