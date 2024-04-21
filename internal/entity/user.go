package entity

type Applicant struct {
	Id int `json:"id" db:"id"`

	//2 форма
	Name       string `json:"name" binding:"required" db:"name"`
	SureName   string `json:"sure_name" binding:"required" db:"sure_name"`
	Patronymic string `json:"patronymic" binding:"required" db:"patronymic"`

	//1 форма
	EMail    string `json:"eMail" binding:"required" db:"email"`
	Password string `json:"password" binding:"required" db:"password_hash"`
}

type Exam struct {//для формы абитуриета
	Id int `json:"id" db:"id"`

	ApplicantId int    `json:"aplplicant_id" binding:"required" db:"applicant_id"`
	ExamName    string `json:"exam_name" binding:"required" db:"exam_name"`
	ExamMark    int    `json:"exam_mark" binding:"required" db:"exam_mark"`
}

type Student struct {
	Id int `json:"id" db:"id"`

	//2 форма
	Name       string `json:"name" binding:"required" db:"name"`
	SureName   string `json:"sure_name" binding:"required" db:"sure_name"`
	Patronymic string `json:"patronymic" binding:"required" db:"patronymic"`

	//1 форма
	EMail    string `json:"eMail" binding:"required" db:"email"`
	Password string `json:"password" binding:"required" db:"password_hash"`

	//3 форма
	University string `json:"university" binding:"required" db:"university"`
	Direction string `json:"direction" binding:"required" db:"direction"`
	Group string `json:"group" binding:"required" db:"group"`
}


