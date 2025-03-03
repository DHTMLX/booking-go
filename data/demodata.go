package data

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func dataDown(tx *gorm.DB) {
	must(tx.Exec("DELETE FROM `doctors`").Error)
	must(tx.Exec("DELETE FROM `reviews`").Error)
	must(tx.Exec("DELETE FROM `slots`").Error)
	must(tx.Exec("DELETE FROM `available_slots`").Error)
	must(tx.Exec("DELETE FROM `reservations`").Error)
}

func dataUp(tx *gorm.DB) {
	now := time.Now().UTC()
	y, m, d := now.Date()
	day := int(now.Weekday())

	// days, hours, minutes
	getDate := func(params ...int) int64 {
		var days, hours, minutes int

		switch len(params) {
		case 3:
			minutes = params[2]
			fallthrough
		case 2:
			hours = params[1]
			fallthrough
		case 1:
			days = params[0]
		}

		return time.Date(y, m, d+days, hours, minutes, 0, 0, time.UTC).UnixMilli()
	}

	// next weekDay, hours, minutes
	getNextDate := func(params ...int) int64 {
		if len(params) > 0 {
			params[0] = (7 + params[0] - day) % 7
		}

		return getDate(params...)
	}

	units := []Doctor{
		{
			Title:    "Dr. Conrad Hubbard",
			Category: "Psychiatrist",
			SubTitle: "2 years of experience",
			Details:  "Desert Springs Hospital (Schroeders Avenue 90, Fannett, Ethiopia)",
			Preview:  "https://snippet.dhtmlx.com/codebase/data/booking/01/img/11.jpg",
			Price:    "$45",
			Review: Review{
				Count: 1245,
				Stars: 4,
			},
			SlotGap:  20,
			SlotSize: 20,
			Slots: []Slot{
				{
					From: "09:00",
					To:   "17:00",
					Size: 40,
					Gap:  10,
					Days: pq.Int32Array{1, 2, 3, 4, 5},
				},
				{
					From: "11:00",
					To:   "19:00",
					Days: pq.Int32Array{0, 6},
				},
			},
			UsedSlots: []Reservation{
				{
					Date: getNextDate(5, 9, 50),
					ReservationForm: ReservationForm{
						ClientName:  "James King",
						ClientEmail: "james.king@booking.demo",
					},
				},
				{
					Date: getNextDate(6, 13, 40),
					ReservationForm: ReservationForm{
						ClientName:    "Lily Wilson",
						ClientEmail:   "lily.wilson@booking.demo",
						ClientDetails: `I'll be early.`,
					},
				},
			},
		},
		{
			Title:    "Dr. Debra Weeks",
			Category: "Allergist",
			SubTitle: "7 years of experience",
			Details:  "Silverstone Medical Center (Vanderbilt Avenue 13, Chestnut, New Zealand)",
			Preview:  "https://snippet.dhtmlx.com/codebase/data/booking/01/img/03.jpg",
			Price:    "$120",
			Review: Review{
				Count: 6545,
				Stars: 4,
			},
			SlotGap:  5,
			SlotSize: 45,
			Slots: []Slot{
				{
					From: "11:00",
					To:   "19:00",
					Size: 40,
					Dates: pq.Int64Array{
						getDate(1),
						getDate(2),
						getDate(3),
					},
				},
			},
			UsedSlots: []Reservation{
				{
					Date: getDate(2, 14),
					ReservationForm: ReservationForm{
						ClientName:  "Olivia Johnson",
						ClientEmail: "olivia.johnson@booking.demo",
					},
				},
			},
		},
		{
			Title:    "Dr. Barnett Mueller",
			Category: "Ophthalmologist",
			SubTitle: "6 years of experience",
			Details:  "Navy Street 1, Kiskimere, United States",
			Preview:  "https://snippet.dhtmlx.com/codebase/data/booking/01/img/02.jpg",
			Price:    "$35",
			Review: Review{
				Count: 184,
				Stars: 3,
			},
			SlotSize: 25,
			Slots: []Slot{
				{
					From: "15:00",
					To:   "19:00",
					Days: pq.Int32Array{1, 2, 3},
				},
				{
					From: "09:00",
					To:   "11:00",
				},
			},
			UsedSlots: []Reservation{
				{
					Date: getNextDate(6, 10, 15),
					ReservationForm: ReservationForm{
						ClientName:  "Mia Adams",
						ClientEmail: "mia.adams@booking.demo",
					},
				},
			},
		},
		{
			Title:    "Dr. Myrtle Wise",
			Category: "Ophthalmologist",
			SubTitle: "4 years of experience",
			Details:  "Prescott Place 5, Freeburn, Bulgaria",
			Preview:  "https://snippet.dhtmlx.com/codebase/data/booking/01/img/01.jpg",
			Price:    "$40",
			Review: Review{
				Count: 829,
				Stars: 5,
			},
			SlotGap:  5,
			SlotSize: 25,
			Slots: []Slot{
				{
					From: "07:00",
					To:   "11:00",
					Days: pq.Int32Array{1, 3, 5},
				},
				{
					From: "15:00",
					To:   "19:00",
					Days: pq.Int32Array{2, 4},
				},
			},
		},
		{
			Title:    "Dr. Browning Peck",
			Category: "Dentist",
			SubTitle: "11 years of experience",
			Details:  "Seacoast Terrace 174, Belvoir, Mauritania",
			Preview:  "https://snippet.dhtmlx.com/codebase/data/booking/01/img/12.jpg",
			Price:    "$175",
			Review: Review{
				Count: 391,
				Stars: 5,
			},
			SlotGap:  10,
			SlotSize: 60,
			Slots: []Slot{ // slots will be ignored
				{
					From:  "09:00",
					To:    "11:00",
					Dates: pq.Int64Array{getDate(1)},
				},
			},
			AvailableSlots: []AvailableSlot{
				{
					Date: getDate(0, 14),
					Size: 40,
				},
				{
					Date: getDate(0, 19, 35),
					Size: 25,
				},
				{
					Date: getDate(1, 9, 55),
					Size: 35,
				},
				{
					Date: getDate(1, 15, 25),
					Size: 20,
				},
			},
			UsedSlots: []Reservation{
				{
					Date: getDate(1, 15, 25),
					ReservationForm: ReservationForm{
						ClientName:  "William Green",
						ClientEmail: "william.green@booking.demo",
					},
				},
			},
		},
	}

	err := tx.Create(units).Error
	if err != nil {
		panic(err)
	}
}
