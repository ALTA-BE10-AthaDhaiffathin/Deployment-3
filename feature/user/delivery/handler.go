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

type userHandler struct {
	userUsecase domain.UserUseCase
}

func New(e *echo.Echo, us domain.UserUseCase) {
	handler := &userHandler{
		userUsecase: us,
	}
	useJWT := middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET)))
	e.POST("/login", handler.LoginUser())
	e.GET("/profile", handler.MyProfile(), useJWT)
	users := e.Group("/users")
	users.GET("/:id", handler.GetSpecificUser())
	users.GET("", handler.GetAll(), useJWT)
	users.POST("", handler.InsertUser())
	users.PUT("", handler.UpdateUser(), useJWT)
	users.DELETE("", handler.DeleteUser(), useJWT)
}

func (uh *userHandler) InsertUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertFormat
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		data, err := uh.userUsecase.AddUser(tmp.ToModel())

		if err != nil {
			log.Println("Cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    data,
		})

	}
}

func (uh *userHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userData []domain.User
		userData, err := uh.userUsecase.GetAll()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "no data found")
		}

		convertedData := []UserResponse{}
		for i := 0; i < len(userData); i++ {
			convertedData = append(convertedData, ToUserResponse(userData[i]))
		}

		res := map[string]interface{}{
			"message": "Get all data",
			"data":    convertedData,
		}
		return c.JSON(http.StatusOK, res)
	}
}

func (uh *userHandler) GetSpecificUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userData domain.User
		param := c.Param("id")
		cnv, err := strconv.Atoi(param)
		if err != nil {
			log.Println("Cant convert param")
			return c.JSON(http.StatusBadRequest, "wrong path")
		}

		userData, err = uh.userUsecase.GetSpecificUser(uint(cnv))

		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusNotFound, err.Error())
		}

		var convertedData UserResponse = ToUserResponse(userData)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Get user " + param,
			"data":    convertedData,
		})
	}
}

func (uh *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)
		var updatedData domain.User
		err := c.Bind(&updatedData)
		if err != nil {
			log.Println("error parsing data")
			return c.JSON(http.StatusBadRequest, "wrong input")
		}

		data, err := uh.userUsecase.UpdateUser(uint(id), updatedData)
		if err != nil {
			log.Println("update failed")
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "user updated",
			"data":    data,
		})
	}
}

func (uh *userHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)
		err := uh.userUsecase.DeleteUser(uint(id))
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, "User Deleted")
	}
}

func (uh *userHandler) LoginUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var loginData LoginFormat
		err := c.Bind(&loginData)
		if err != nil {
			log.Println("error parsing data", err.Error())
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		data, token, err := uh.userUsecase.LoginUser(loginData.ToModel())
		if err != nil {
			log.Println("server error", err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Login Success",
			"data":    data,
			"token":   token,
		})
	}
}

func (uh *userHandler) MyProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)
		data, err := uh.userUsecase.UserProfile(uint(id))
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get your profile",
			"data":    data,
		})
	}
}
