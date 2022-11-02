package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model    `json:"-"`
	ID            int          `json:"id"`
	Nome          string       `json:"nome" validate:"nonzero"`
	Idade         int          `json:"idade"`
	TipoClienteID int          `json:"tipoclienteid"`
	Endereco      string       `json:"endereco"`
	Dependentes   []Dependente `json:"-" gorm:"foreignKey:ClienteID"`
}

func ValidaDadosCliente(cliente *Cliente) error {

	if err := validator.Validate(cliente); err != nil {
		return err
	}

	return nil
}
