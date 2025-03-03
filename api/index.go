package api

import (
	"booking-go/data"
	"context"
	"errors"
	"net/http"

	remote "github.com/mkozhukh/go-remote"
)

type UserID int
type DeviceID int

type Reservation struct {
	From int              `json:"-"`
	Type string           `json:"type"`
	Data data.Reservation `json:"reservation"`
}

func BuildAPI(db *data.DAO) *remote.Server {
	if remote.MaxSocketMessageSize < 32000 {
		remote.MaxSocketMessageSize = 32000
	}

	api := remote.NewServer(&remote.ServerConfig{
		WebSocket: true,
	})

	api.Events.AddGuard("reservations", guardFilter)

	api.Connect = func(r *http.Request) (context.Context, error) {
		id, _ := r.Context().Value("user_id").(int)
		if id == 0 {
			return nil, errors.New("access denied")
		}
		device, _ := r.Context().Value("device_id").(int)
		if device == 0 {
			return nil, errors.New("access denied")
		}

		return context.WithValue(
			context.WithValue(r.Context(), remote.UserValue, id),
			remote.ConnectionValue, device), nil
	}

	return api
}

func guardFilter(m *remote.Message, c *remote.Client) bool {
	tm, ok := m.Content.(Reservation)
	if !ok {
		return false
	}

	return int(tm.From) != c.ConnID
}
