package entity

type EducatitionalDirection struct {
	Id            int     `json:"id" db:"id" binding:"required"`
	Name          string  `json:"Направление"`
	Group         string  `json:"Конкурсн. группа"`
	CountBudget   int     `json:"КЦП_Б"`
	CountContract int     `json:"КЦП_К"`
	Price         int     `json:"Цена"`
	Subject1      string  `json:"Предмет 1"`
	Subject2      string  `json:"Предмет 2"`
	Subject3      string  `json:"Предмет 3"`
	Value1        int     `json:"Мин. балл 1"`
	Value2        int     `json:"Мин. балл 2"`
	Value3        int     `json:"Мин. балл 3"`
	Sum           int     `json:"Балл"`
	CompetiveB    float32 `json:"Конкурс_Б"`
	CompetiveK    float32 `json:"Конкурс_К"`
}
