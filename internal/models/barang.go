package models

import "gorm.io/gorm"

type JenisBarang struct {
	gorm.Model
	NamaJenis string `json:"nama_jenis" gorm:"unique;not null"`
}
