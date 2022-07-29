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

type commentHandler struct {
	commentCases domain.CommentUseCases
}

func New(cs domain.CommentUseCases) domain.CommentHandler {
	return &commentHandler{
		commentCases: cs,
	}
}

func (cs *commentHandler) GetAllComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := cs.commentCases.GetAllComment()
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Println("User Handler", err)
				c.JSON(http.StatusNotFound, err.Error())
			} else if strings.Contains(err.Error(), "retrieve") {
				log.Println("User Handler", err)
				c.JSON(http.StatusInternalServerError, err.Error())
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all user data",
			"data":    data,
		})
	}
}

func (cs *commentHandler) PostComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := common.ExtractData(c)
		var temp CommentFormat
		err := c.Bind(&temp)

		if err != nil {
			log.Println("Tidak bisa parse data", err)
			c.JSON(http.StatusBadRequest, "tidak bisa membaca input")
		}

		data, err := cs.commentCases.PostingComment(userId, temp.ToModel())

		if err != nil {
			log.Println("tidak memproses data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "berhasil register data",
			"data":    data,
		})
	}
}

func (cs *commentHandler) DeleteComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		data, err := cs.commentCases.DeleteComment(cnv)

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
			"message": "success delete comment",
		})
	}
}
