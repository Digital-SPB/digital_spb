package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greenblat17/digital_spb/internal/entity"
)

func (h *Handler) StudentSignIn(c *gin.Context) {
	fmt.Println("sign in handler")
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

	id, err := h.service.CreateStudent(context.Background(), input)
	if err!=nil {
		newErrorResponse(c, http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
