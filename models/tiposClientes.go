package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type TipoCliente struct {
	gorm.Model
	id        int    `gorm:"primaryKey,autoIncrement:true"`
	Descricao string `json:"descricao" validate:"nonzero"`
}

func ValidaDadosDeAluno(tipoCliente *TipoCliente) error {

	if err := validator.Validate(tipoCliente); err != nil {
		return err
	}

	return nil
}
