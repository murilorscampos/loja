package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Dependente struct {
	gorm.Model `json:"-"`
	ID         int    `json:"id"`
	Nome       string `json:"nome" validate:"nonzero"`
	Idade      int    `json:"idade"`
	ClienteID  int    `json:"clienteid"`
}

func ValidaDadosDependente(dependente *Dependente) error {

	if err := validator.Validate(dependente); err != nil {
		return err
	}

	return nil
}
