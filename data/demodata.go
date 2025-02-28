package data

import (
	"log"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func dataDown(tx *gorm.DB) {
	log.Print("dowwwn")
	must(tx.Exec("DELETE FROM `doctors`").Error)
	must(tx.Exec("DELETE FROM `available_slots`").Error)
	must(tx.Exec("DELETE FROM `slots`").Error)
	must(tx.Exec("DELETE FROM `reviews`").Error)
	must(tx.Exec("DELETE FROM `reservations`").Error)
}

// AvailableSlots []AvailableSlot `json:"available_slots"`
func dataUp(tx *gorm.DB) {
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
			// SlotGap:  20,
			// SlotSize: 20,
			Slots: []Slot{
				{
					From:  "09:00",
					To:    "17:00",
					Size:  40,
					Gap:   40,
					Days:  pq.Int64Array{1, 2, 3, 4, 5},
					Dates: pq.Int64Array{10000000000, 2, 3, 4, 5},
					// Dates []int  `json:"dates"`
				},
			},
			AvailableSlots: AvailableSlot{
				Times:   pq.Int64Array{28937489273, 2397428934, 23784293874},
				Lengths: pq.Int64Array{28937489273, 2397428934, 23784293874},
			},
			UsedSlots: []Reservation{
				{
					ClientName:    "eduard",
					ClientEmail:   "blabla@gmail.com",
					ClientDetails: "blabla BLA",
					Date:          2000000000,
				},
			},
		},
		// },
		// {
		// 	Title:    "Dr. Debra Weeks",
		// 	Category: "Allergist",
		// 	SubTitle: "7 years of experience",
		// 	Details:  "Silverstone Medical Center (Vanderbilt Avenue 13, Chestnut, New Zealand)",
		// 	Preview:  "https://snippet.dhtmlx.com/codebase/data/booking/01/img/03.jpg",
		// 	Price:    "$120",
		// 	Review: Review{
		// 		Count: 6545,
		// 		Stars: 4,
		// 	},
		// 	SlotGap:  5,
		// 	SlotSize: 45,
		// 	Slots: append(
		// 		[]Slot{
		// 			// mon, wed 7:00-15:00
		// 			RecurringSchedule(7*60, 15*60, "MO,WE"),
		// 			// tue, thu 12:00-20:00
		// 			RecurringSchedule(12*60, 20*60, "TU,TH"),
		// 			// sat-sun 20:00-4:00
		// 			RecurringSchedule(20*60, 4*60, "SA"), // or RecurringSchedule(20*60, 28*60, "SA")
		// 		},
		// 		// next wed 18:00-22:00
		// 		genSchedule(18*60, 22*60, nextWeekDay(3), 1)...,
		// 	),
		// 	UsedSlots: genSlots(
		// 		newSlots(nextWeekDay(1), 7*60+50),            // next mon 7:50
		// 		newSlots(nextWeekDay(2), 13*60+40),           // next tue 13:40
		// 		newSlots(nextWeekDay(3), 11*60+10),           // next wed 11:10
		// 		newSlots(nextWeekDay(4), 14*60+30, 17*60+50), // next thu 14:30 17:50
		// 		newSlots(nextWeekDay(4, 1), 17*60+50),        // after next thu 17:50
		// 		newSlots(nextWeekDay(0), 2*60+40),            // next SUN 2:40; or newSlots(nextWeekDay(6), 24*60+2*60+40)
		// 	),
		// },
		// {
		// 	Title:    "Dr. Barnett Mueller",
		// 	Category: "Ophthalmologist",
		// 	SubTitle: "6 years of experience",
		// 	Details:  "Navy Street 1, Kiskimere, United States",
		// 	Preview:  "https://snippet.dhtmlx.com/codebase/data/booking/01/img/02.jpg",
		// 	Price:    "$35",
		// 	Review: Review{
		// 		Count: 184,
		// 		Stars: 3,
		// 	},
		// 	SlotGap:  0,
		// 	SlotSize: 25,
		// 	Slots: []Slot{
		// 		// mon, wed, fri 9:00-17:00
		// 		RecurringSchedule(9*60, 17*60, "MO,WE,FR"),
		// 		// sat, sun 15:00-19:00
		// 		RecurringSchedule(15*60, 19*60, "SA,SU"),
		// 	},
		// 	UsedSlots: genSlots(
		// 		newSlots(nextWeekDay(1), 13*60+10),    // after next mon 13:10
		// 		newSlots(nextWeekDay(1, 1), 12*60+45), // after next mon 12:45
		// 		newSlots(nextWeekDay(3), 9*60+25),     // next wed 9:25
		// 		newSlots(nextWeekDay(5), 11*60+55),    // next fri 11:55
		// 		newSlots(nextWeekDay(5, 1), 11*60+30), // after next fri 11:30
		// 		newSlots(nextWeekDay(6), 16*60+10),    // next sat 16:10
		// 		newSlots(nextWeekDay(0), 17*60),       // next sun 17:00
		// 	),
		// },
		// {
		// 	Title:    "Dr. Myrtle Wise",
		// 	Category: "Ophthalmologist",
		// 	SubTitle: "4 years of experience",
		// 	Details:  "Prescott Place 5, Freeburn, Bulgaria",
		// 	Preview:  "https://snippet.dhtmlx.com/codebase/data/booking/01/img/01.jpg",
		// 	Price:    "$40",
		// 	Review: Review{
		// 		Count: 829,
		// 		Stars: 5,
		// 	},
		// 	SlotGap:  5,
		// 	SlotSize: 25,
		// 	Slots: append(
		// 		[]Slot{
		// 			// tue, thu 7:00-15:00
		// 			RecurringSchedule(7*60, 15*60, "TU,TH"),
		// 			// sat, sun 11:00-15:00
		// 			RecurringSchedule(11*60, 15*60, "SA,SU"),
		// 		},
		// 		// next fri, sat 4:00-8:00
		// 		genSchedule(4*60, 8*60, nextWeekDay(5), 2)...,
		// 	),
		// 	UsedSlots: genSlots(
		// 		newSlots(nextWeekDay(2), 7*60, 10*60),    // next tue 7:00, 10:00
		// 		newSlots(nextWeekDay(4), 9*60+30),        // next thu 9:30
		// 		newSlots(nextWeekDay(5), 7*60+30),        // next fri 7:30
		// 		newSlots(nextWeekDay(6), 11*60+30, 5*60), // next sat 11:30, 5:00
		// 		newSlots(nextWeekDay(0), 12*60),          // next sun 12:00
		// 	),
		// },
		// {
		// 	Title:    "Dr. Browning Peck",
		// 	Category: "Dentist",
		// 	SubTitle: "11 years of experience",
		// 	Details:  "Seacoast Terrace 174, Belvoir, Mauritania",
		// 	Preview:  "https://snippet.dhtmlx.com/codebase/data/booking/01/img/12.jpg",
		// 	Price:    "$175",
		// 	Review: Review{
		// 		Count: 391,
		// 		Stars: 5,
		// 	},
		// 	SlotGap:  10,
		// 	SlotSize: 60,
		// 	Slots: []Slot{
		// 		// thu, fri, sat, sun 9:00-17:00
		// 		RecurringSchedule(9*60, 17*60, "TH,FR,SA,SU"),
		// 	},
		// 	UsedSlots: genSlots(
		// 		newSlots(nextWeekDay(4), 11*60+20),       // next thu 11:20
		// 		newSlots(nextWeekDay(5), 14*60+50),       // next fri 14:50
		// 		newSlots(nextWeekDay(6), 9*60, 13*60+20), // next sat 9:00, 13:20
		// 		newSlots(nextWeekDay(0), 14*60+50),       // next sun 14:50
		// 	),
		// },
	}

	err := tx.Create(units).Error
	if err != nil {
		panic(err)
	}
}
