package models

import (
	"gorm.io/datatypes"
)

type Bestellung struct {
	Id           int            `json:"id" gorm:"primary_key"`
	Artikelbez   string         `json:"artikelbez"`
	Artikelnr    int            `json:"artikelnr"`
	KundenID     int            `json:"kunde_id"`
	Bestellmenge int            `json:"menge"`
	Bestelldatum datatypes.Date `json:"datum"`
}
