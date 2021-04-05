package domain

import "time"

type URLClient interface {
	Find(ID string) *Url
	FindOrCreateURL(destination string) (*Url, bool, error)
	AddLogClick(ID string)
	GetStatistics(u *Url) *Statistics
}

type URLService interface {
	Find(ID string) *Url
	FindOrCreateURL(destination string) (*Url, bool, error)
	GetStatistics(url *Url) *Statistics
	AddStatistics(url *Url)
	CollectStatistics()
}

type Url struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"criacao"`
	Destination string    `json:"destino"`
}

type Statistics struct {
	URL    *Url `json:"url"`
	Clicks int  `json:"clicks"`
}
