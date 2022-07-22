package delivery

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/config"
	"github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/domain"
	"github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/feature/common"
	"github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/feature/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type bookHandler struct {
	bookUsecase domain.BookUseCase
}

func New(e *echo.Echo, bs domain.BookUseCase) {
	handler := &bookHandler{
		bookUsecase: bs,
	}
	useJWT := middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET)))
	books := e.Group("/books")
	books.POST("", handler.AddBook(), useJWT)
	books.GET("", handler.GetAllBooks())
	books.GET("/:id", handler.GetSpecificBook())
	books.PUT("/:id", handler.UpdateBook(), useJWT)
	books.DELETE("/:id", handler.DeleteBook(), useJWT)
}

func (bh *bookHandler) AddBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		var bookData InsertFormat
		err := c.Bind(&bookData)
		if err != nil {
			log.Println("failed parsing data", err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		convertToDomain := bookData.ToDomain()
		convertToDomain.User_ID = common.ExtractData(c)
		convertToDomain, err = bh.bookUsecase.AddBook(convertToDomain)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "book created",
			"data":    convertToDomain,
		})
	}
}

func (bh *bookHandler) GetAllBooks() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := bh.bookUsecase.GetAll()
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all books",
			"data":    data,
		})
	}
}

func (bh *bookHandler) GetSpecificBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		cnv, err := strconv.Atoi(id)
		if err != nil {
			log.Println("error convert id", err.Error())
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		data, err := bh.bookUsecase.GetSpecificBook(uint(cnv))
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "get user " + id,
			"data":    data,
		})
	}
}

func (bh *bookHandler) UpdateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		cnv, err := strconv.Atoi(id)
		if err != nil {
			log.Println("error convert id", err.Error())
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		var updatedData domain.Book
		err = c.Bind(&updatedData)
		if err != nil {
			log.Println("Wrong input", err.Error())
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		userID := common.ExtractData(c)
		data, err := bh.bookUsecase.UpdateBook(uint(userID), uint(cnv), updatedData)
		if err != nil {
			log.Println("server error", err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update data book",
			"data":    data,
		})
	}
}

func (bh *bookHandler) DeleteBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		cnv, err := strconv.Atoi(id)
		if err != nil {
			log.Println("error convert id", err.Error())
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		userID := common.ExtractData(c)
		err = bh.bookUsecase.DeleteBook(uint(userID), uint(cnv))
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, "success delete book data")
	}
}
