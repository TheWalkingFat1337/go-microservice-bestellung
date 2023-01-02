package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Bestellung struct {
	gorm.Model
	Artikelbez   string         `json:"artikelbez"`
	Artikelnr    int            `json:"artikelnr"`
	Kundennr     int            `json:"kundennr"`
	Bestellmenge int            `json:"menge"`
	Bestelldatum datatypes.Date `json:"datum"`
}
