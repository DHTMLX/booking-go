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

	Review Review `gorm:"foreignkey:DoctorID"`

	SlotGap  int `json:"slot_gap"`
	SlotSize int `json:"slot_size"`

	Slots          []Slot        `gorm:"foreignkey:DoctorID"`
	UsedSlots      []Reservation `gorm:"foreignkey:DoctorID"`
	AvailableSlots AvailableSlot `gorm:"foreignkey:DoctorID"`
}

type Review struct {
	ID       int `json:"id"`
	DoctorID int `json:"-"`
	Stars    int `json:"star"`
	Count    int `json:"count"`
}

type Slot struct {
	ID       int           `json:"id"`
	DoctorID int           `json:"doctor_id"`
	From     string        `json:"from"`
	To       string        `json:"to"`
	Size     int           `json:"size"`
	Gap      int           `json:"gap,omitempty"`
	Days     pq.Int64Array `json:"days" gorm:"type:integer[]"`
	Dates    pq.Int64Array `json:"dates" gorm:"type:integer[]"`
}

type Reservation struct {
	ID            int    `json:"id"`
	DoctorID      int    `json:"doctor_id"`
	Date          int64  `json:"date"`
	ClientName    string `json:"client_name"`
	ClientEmail   string `json:"client_email"`
	ClientDetails string `json:"client_details"`
}

type AvailableSlot struct {
	ID       int           `json:"id"`
	DoctorID int           `json:"doctor_id"`
	Times    pq.Int64Array `gorm:"type:integer[]"`
	Lengths  pq.Int64Array `gorm:"type:integer[]"`
}
