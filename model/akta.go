package model

import (
	"encoding/json"
)

// Akta ...
type Akta struct {
	ID         int    `json:"id" gorm:"primary_key"`
	AktaNumber string `json:"akta_number"`
	AktaDate   string `json:"akta_date"`
	AktaAuthor string `json:"akta_author"`
	FullName   string `json:"full_name"`
	FatherName string `json:"father_name"`
	MotherName string `json:"mother_name"`
	Gender     string `json:"gender"`
	BirthPlace string `json:"birth_place"`
	BirthDate  string `json:"birth_date"`
}

// FromJSON ...
func FromJSON(str string) (akta Akta, err error) {
	err = json.Unmarshal([]byte(str), &akta)

	return
}

// ToJSON ...
func (a Akta) ToJSON() (string, error) {
	data, err := json.Marshal(&a)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// NewAkta ...
func NewAkta(id int, aktaNum, aktaDate, aktaAuth, fullname, fathername, mothername, gender, birthplace, birthdate string) (Akta, error) {
	return Akta{
		id,
		aktaNum,
		aktaDate,
		aktaAuth,
		fullname,
		fathername,
		mothername,
		gender,
		birthplace,
		birthdate,
	}, nil
}
