package data

import (
	"gorm.io/gorm"
)

type ReservationsDAO struct {
	db *gorm.DB
}

func newReservationsDAO(db *gorm.DB) *ReservationsDAO {
	return &ReservationsDAO{db}
}

func (d *ReservationsDAO) openTX() *gorm.DB {
	return d.db.Begin()
}

func (d *ReservationsDAO) closeTX(tx *gorm.DB, err error) {
	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
}

func (d *ReservationsDAO) GetAll() ([]Reservation, error) {
	reservations := make([]Reservation, 0)
	err := d.db.Find(&reservations).Error
	return reservations, err
}

func (d *ReservationsDAO) Add(reservation *Reservation) (int, error) {
	err := d.db.Create(&reservation).Error
	if err != nil {
		return 0, err
	}

	return reservation.ID, nil
}
