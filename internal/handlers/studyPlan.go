package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExamMarks struct {
	Name string `json:"name"`
	Mark int    `json:"mark"`
}

type ApplicantInput struct {
	ExamMarks []ExamMarks `json:"exam_marks"`
	Vacancy   string      `json:"vacancy"`
}

func (h *Handler) ApplicantStudyPlan(c *gin.Context) {
	var input ApplicantInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
	}

	educations, err := h.service.EducationalDirection.GetEducationalDirectionForApplicant(context.Background(), 1)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, educations)
}
