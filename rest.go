package main

import (
	"booking-go/data"
	"net/http"

	"github.com/go-chi/chi"
	// go_remote "github.com/mkozhukh/go-remote"
)

func initRoutes(r chi.Router, dao *data.DAO) { // , hub *go_remote.Hub) {
	r.Get("/units", func(w http.ResponseWriter, r *http.Request) {
		data, err := dao.Units.GetAll()
		sendResponse(w, data, err)
	})

	r.Post("/reservations", func(w http.ResponseWriter, r *http.Request) {
		item := data.Reservation{}
		err := parseFormObject(w, r, &item)
		if err != nil {
			format.Text(w, 500, err.Error())
			return
		}

		id, err := dao.Reservations.Add(&item)
		sendResponse(w, id, err)
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
