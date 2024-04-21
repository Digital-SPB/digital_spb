package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greenblat17/digital_spb/internal/entity"
)

func (h *Handler) ApplicantSignIn(c *gin.Context) {
	fmt.Println("sign in handler")
}

func (h *Handler) ApplicantSignUp(c *gin.Context) { //регистрация
	var input entity.Applicant
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	//fmt.Println(len(input.Exams))
	id, err := h.service.ApplicantAuth.CreateApplicant(context.Background(), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
