package models

import "gorm.io/gorm"

type Barang struct {
	gorm.Model
	NamaBarang string      `json:"nama_barang" gorm:"not null"`
	Harga      float64     `json:"harga" gorm:"not null"`
	JenisID    uint        `json:"jenis_id"`
	Jenis      JenisBarang `json:"jenis" gorm:"foreignKey:JenisID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
