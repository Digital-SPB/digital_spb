package entity

type Vacancy struct {
	Id        int    `json:"id" binding:"required" db:"id"`
	Name      string `json:"name" binding:"required" db:"name"`
	Education string `json:"education" binding:"required" db:"education"`
}
