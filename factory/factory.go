package factory

import (
	bookData "github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/feature/book/data"
	bookDelivery "github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/feature/book/delivery"
	bookUsecase "github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/feature/book/usecase"
	userData "github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/feature/user/data"
	userDelivery "github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/feature/user/delivery"
	userUsecase "github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/feature/user/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	userData := userData.New(db)
	useCase := userUsecase.New(userData)
	userDelivery.New(e, useCase)

	bookData := bookData.New(db)
	bookUseCase := bookUsecase.New(bookData)
	bookDelivery.New(e, bookUseCase)
}
