package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greenblat17/digital_spb/internal/entity"
)

type signInInput struct {
	EMail    string `json:"email" binding:"required" db:"user_name"`
	Password string `json:"password" binding:"required" db:"password_hash"`
}

func (h *Handler) StudentSignIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
	}

	token, err := h.service.StudentAuth.GenerateToken(context.Background(), input.EMail, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) StudentSignUp(c *gin.Context) {
	var input entity.Student
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
	}

	// id, err := h.services.Authorization.CreateAdmin(context.Background(), input)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	id, err := h.service.StudentAuth.CreateStudent(context.Background(), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
