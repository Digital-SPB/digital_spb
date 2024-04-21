package entity

type Vacancy struct {
	Id   int    `json:"id" binding:"required" db:"id"`
	Name string `json:"name" binding:"required" db:"name"`
}

type VacancyEducation struct {
	Id          int `json:"id" binding:"required" db:"id"`
	VacancyId   int `json:"vacancy_id" binding:"required" db:"vacancy_id"`
	EducationId int `json:"education_id" binding:"required" db:"education_id"`
}
