package models

import (
	"gorm.io/datatypes"
)

type Bestellung struct {
	Id           int            `json:"id" gorm:"primary_key"`
	Artikelbez   string         `json:"artikelbez"`
	Artikelnr    int            `json:"artikelnr"`
	Kundennr     int            `json:"kundennr"`
	Bestellmenge int            `json:"menge"`
	Bestelldatum datatypes.Date `json:"datum"`
}
