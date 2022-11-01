package models

import "gorm.io/gorm"

type Dependente struct {
	gorm.Model
	id        int     `gorm:"primaryKey,autoIncrement:true"`
	Nome      string  `json:"nome"`
	Idade     int     `json:"idade"`
	IdCliente int     `json:"idCliente"`
	Cliente   Cliente `gorm:"foreignKey:IdCliente"`
}
