package models

import "gorm.io/gorm"

type Cliente struct {
	gorm.Model
	id          int         `gorm:"primaryKey,autoIncrement:true"`
	Nome        string      `json:"nome"`
	Idade       int         `json:"idade"`
	TpCliente   int         `json:"tpcliente"`
	Endereco    string      `json:"endereco"`
	TipoCliente TipoCliente `gorm:"foreignKey:TpCliente"`
}
