package data

import (
	"gorm.io/gorm"
)

type Unit struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	SubTitle string `json:"subtitle"`
	Details  string `json:"details"`
	Preview  string `json:"preview"`
	Price    string `json:"price"`

	Review Review `json:"review"`

	SlotGap  int `json:"slot_gap"`
	SlotSize int `json:"slot_size"`

	Slots []Slot `gorm:"foreignkey:DoctorID"`
	// UsedSlots      []Reservation `gorm:"foreignkey:DoctorID"`
	// AvailableSlots AvailableSlot `gorm:"foreignkey:DoctorID"`

	// Slots          []Schedule `json:"slots"`
	AvailableSlots []int64 `json:"availableSlots,omitempty"`
	UsedSlots      []int64 `json:"usedSlots,omitempty"`
}

type UnitsDAO struct {
	db *gorm.DB
}

func newUnitsDAO(db *gorm.DB) *UnitsDAO {
	return &UnitsDAO{db}
}

func (d *UnitsDAO) openTX() *gorm.DB {
	return d.db.Begin()
}

func (d *UnitsDAO) closeTX(tx *gorm.DB, err error) {
	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
}

func (d *UnitsDAO) GetAll() ([]Unit, error) {
	// doctors := make([]Doctor, 0)
	// err := d.db.Preload("Review").Preload("Slots").Preload("AvailableSlots").Preload("UsedSlots").Find(&doctors).Error

	// units := make([]Unit, 0, len(doctors))
	// for _, doctor := range doctors {
	// 	unit := Unit{}
	// }

	return nil, nil // Units, err
}
