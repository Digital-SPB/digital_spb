package handlers

import (
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

}
