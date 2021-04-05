package domain

type Repository interface {
	ExistsID(id string) bool
	FindByID(id string) *Url
	FindByURL(url string) *Url
	Save(url Url) error
	AddLogClick(id string)
	FindLogClicks(id string) int
}
