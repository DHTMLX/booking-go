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

	SlotGap  int `json:"slotGap,omitempty"`
	SlotSize int `json:"slotSize,omitempty"`

	Slots          []Slot     `json:"slots,omitempty"`
	AvailableSlots [][2]int64 `json:"availableSlots,omitempty"`
	UsedSlots      []int64    `json:"usedSlots,omitempty"`
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
	doctors := make([]Doctor, 0)
	err := d.db.
		Preload("Review").
		Preload("Slots").
		Preload("AvailableSlots").
		Preload("UsedSlots").
		Find(&doctors).Error

	units := make([]Unit, len(doctors))
	for i, doctor := range doctors {
		slots := doctor.Slots
		if len(doctor.AvailableSlots) > 0 {
			slots = []Slot{}
		}

		booked := make(map[int64]struct{}, len(doctor.UsedSlots))
		usedSlots := make([]int64, len(doctor.UsedSlots))
		for j, slot := range doctor.UsedSlots {
			booked[slot.Date] = struct{}{}
			usedSlots[j] = slot.Date
		}

		availableSlots := make([][2]int64, 0, len(doctor.AvailableSlots))
		for _, slot := range doctor.AvailableSlots {
			if _, ok := booked[slot.Date]; ok {
				continue
			}

			availableSlots = append(availableSlots, [2]int64{
				slot.Date,
				int64(slot.Size),
			})
		}

		units[i] = Unit{
			ID:       doctor.ID,
			Title:    doctor.Title,
			Category: doctor.Category,
			SubTitle: doctor.SubTitle,
			Details:  doctor.Details,
			Preview:  doctor.Preview,
			Price:    doctor.Price,

			Review: doctor.Review,

			SlotGap:  doctor.SlotGap,
			SlotSize: doctor.SlotSize,

			Slots:          slots,
			AvailableSlots: availableSlots,
			UsedSlots:      usedSlots,
		}
	}

	return units, err
}
