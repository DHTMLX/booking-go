package main

import (
	"booking-go/api"
	"booking-go/data"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	go_remote "github.com/mkozhukh/go-remote"
)

func initRoutes(r chi.Router, dao *data.DAO, hub *go_remote.Hub) {
	r.Get("/units", func(w http.ResponseWriter, r *http.Request) {
		data, err := dao.Units.GetAll()
		sendResponse(w, data, err)
	})

	r.Post("/doctors/reservations", func(w http.ResponseWriter, r *http.Request) {
		reservation := data.Reservation{}
		err := parseFormObject(w, r, &reservation)
		if err != nil {
			format.Text(w, 500, err.Error())
			return
		}

		id, err := dao.Reservations.Add(&reservation)
		if sendResponse(w, id, err) {
			hub.Publish("reservations", api.Reservation{
				Type: "add-reservation",
				From: getDeviceID(r),
				Data: reservation,
			})
		}
	})

	// DEMO ONLY, imitate login
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		uid, _ := strconv.Atoi(r.URL.Query().Get("id"))
		device := newDeviceID()
		token, err := createUserToken(uid, device)
		if err != nil {
			log.Println("[token]", err.Error())
		}
		w.Write(token)
	})
}

func sendResponse(w http.ResponseWriter, data interface{}, err error) bool {
	if err != nil {
		format.Text(w, 500, err.Error())
	} else {
		if data == nil {
			data = Response{}
		}
		format.JSON(w, 200, data)
	}

	return err == nil
}
