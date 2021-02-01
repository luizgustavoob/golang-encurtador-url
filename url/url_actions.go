package url

import (
	"math/rand"
	"net/url"
	"time"
)

const (
	size    = 5
	simbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-+"
)

var repo Repository

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetStatistics(u *Url) *Statistics {
	clicks := repo.FindLogClicks(u.ID)
	return &Statistics{u, clicks}
}

func ConfigureRepository(rep Repository) {
	repo = rep
}

func Find(id string) *Url {
	return repo.FindByID(id)
}

func FindOrCreateURL(destination string) (*Url, bool, error) {
	if u := repo.FindByURL(destination); u != nil {
		return u, false, nil
	}

	if _, err := url.ParseRequestURI(destination); err != nil {
		return nil, false, err
	}

	url := Url{generateID(), time.Now(), destination}
	repo.Save(url)
	return &url, true, nil
}

func AddLogClick(id string) {
	repo.AddLogClick(id)
}

func generateID() string {
	newID := func() string {
		id := make([]byte, size, size)
		for i := range id {
			id[i] = simbols[rand.Intn(len(simbols))]
		}
		return string(id)
	}

	for {
		if id := newID(); !repo.ExistsID(id) {
			return id
		}
	}
}
