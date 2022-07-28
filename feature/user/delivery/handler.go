package delivery

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/AltaProject/AltaSocialMedia/domain"
	"github.com/AltaProject/AltaSocialMedia/feature/common"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUsecase domain.UserUseCases
}

func New(us domain.UserUseCases) domain.UserHandler {
	return &userHandler{
		userUsecase: us,
	}
}

func (us *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var temp RegisterFormat
		err := c.Bind(&temp)

		if err != nil {
			log.Println("Tidak bisa parse data", err)
			c.JSON(http.StatusBadRequest, "tidak bisa membaca input")
		}

		data, err := us.userUsecase.Register(temp.ToModel())

		if err != nil {
			log.Println("tidak memproses data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "berhasil register data",
			"data":    data,
			"token":   common.GenerateToken(data.ID),
		})
	}
}

func (us *userHandler) GetSpecificUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)

		data, err := us.userUsecase.GetSpecificUser(id)

		if err != nil {
			log.Println("data tidak ditemukan", err)
			return c.JSON(http.StatusNotFound, "data tidak ditemukan")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "data ditemukan",
			"data":    data,
		})
	}
}

func (us *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userLogin LoginFormat
		errLog := c.Bind(&userLogin)
		if errLog != nil {
			return c.JSON(http.StatusBadRequest, "email atau password salah")
		}
		username, token, err := us.userUsecase.Login(userLogin.Email, userLogin.Password)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "login gagal")
		}
		data := map[string]interface{}{
			"username": username,
			"token":    token,
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "login berhasil",
			"data":    data,
		})
	}
}

func (us *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp RegisterFormat
		err := c.Bind(&tmp)
		idCon := c.Param("id")
		id, _ := strconv.Atoi(idCon)

		if err != nil {
			log.Println("Cannot Parse Data", err)
			c.JSON(http.StatusBadRequest, "error read update")
		}
		data, err := us.userUsecase.UpdateUser(tmp.ToModel(), id)

		if err != nil {
			log.Println("Cannot Update Content", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    data,
			"message": "Update user Success",
		})

	}
}

func (us *userHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		data, err := us.userUsecase.DeleteUser(cnv)

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, err.Error())
			} else {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}

		if !data {
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete user",
		})
	}
}
