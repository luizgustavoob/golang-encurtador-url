package urlentities

import "time"

type (
	Url struct {
		ID          string    `json:"id"`
		CreatedAt   time.Time `json:"criacao"`
		Destination string    `json:"destino"`
	}

	Statistics struct {
		URL    *Url `json:"url"`
		Clicks int  `json:"clicks"`
	}
)
