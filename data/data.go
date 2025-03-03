package data

import "github.com/lib/pq"

type Doctor struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	SubTitle string `json:"subtitle"`
	Details  string `json:"details"`
	Preview  string `json:"preview"`
	Price    string `json:"price"`

	Review Review `json:"review" gorm:"foreignkey:DoctorID"`

	SlotGap  int `json:"slotGap"`
	SlotSize int `json:"slotSize"`

	Slots          []Slot          `json:"slots" gorm:"foreignkey:DoctorID"`
	AvailableSlots []AvailableSlot `gorm:"foreignkey:DoctorID"`
	UsedSlots      []Reservation   `gorm:"foreignkey:DoctorID"`
}

type Review struct {
	ID       int `json:"-"`
	DoctorID int `json:"-"`
	Stars    int `json:"star"`
	Count    int `json:"count"`
}

type Slot struct {
	ID       int           `json:"-"`
	DoctorID int           `json:"-"`
	From     string        `json:"from"`
	To       string        `json:"to"`
	Size     int           `json:"size,omitempty"`
	Gap      int           `json:"gap,omitempty"`
	Days     pq.Int32Array `json:"days,omitempty" gorm:"type:integer[]"`
	Dates    pq.Int64Array `json:"dates,omitempty" gorm:"type:integer[]"`
}

type AvailableSlot struct {
	ID       int
	DoctorID int
	Date     int64
	Size     int
}

type ReservationForm struct {
	ClientName    string `json:"name"`
	ClientEmail   string `json:"email"`
	ClientDetails string `json:"details"`
}

type Reservation struct {
	ID       int   `json:"id"`
	DoctorID int   `json:"doctor" gorm:"uniqueIndex:idx_reservation"`
	Date     int64 `json:"date" gorm:"uniqueIndex:idx_reservation"`

	ReservationForm `json:"form"`
}
