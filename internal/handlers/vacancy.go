package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ApplicantVacancies(c *gin.Context) {
	vacancies, err := h.service.Vacancy.GetVacancies(context.Background())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, vacancies)
}
