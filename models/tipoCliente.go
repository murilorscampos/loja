package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type TipoCliente struct {
	gorm.Model `json:"-"`
	ID         int       `json:"id"`
	Descricao  string    `json:"descricao" validate:"nonzero"`
	Clientes   []Cliente `json:"-" gorm:"foreignKey:TipoClienteID"`
}

func ValidaDadosTipoCliente(tipoCliente *TipoCliente) error {

	if err := validator.Validate(tipoCliente); err != nil {
		return err
	}

	return nil
}
