package storage

import "github.com/golang-encurtador-url/domain"

type RepositoryMock struct {
	SaveInvokedCount          int
	ExistsIDInvokedCount      int
	FindByIdInvokedCount      int
	FindByURLInvokedCount     int
	AddLogClickInvokedCount   int
	FindLogClicksInvokedCount int

	SaveFn          func(domain.Url) error
	ExistsIDFn      func(id string) bool
	FindByIdFn      func(id string) *domain.Url
	FindByURLFn     func(destino string) *domain.Url
	AddLogClickFn   func(id string)
	FindLogClicksFn func(id string) int
}

func (self *RepositoryMock) Save(url domain.Url) error {
	self.SaveInvokedCount++
	return self.SaveFn(url)
}

func (self *RepositoryMock) ExistsID(id string) bool {
	self.ExistsIDInvokedCount++
	return self.ExistsIDFn(id)
}

func (self *RepositoryMock) FindByID(id string) *domain.Url {
	self.FindByIdInvokedCount++
	return self.FindByIdFn(id)
}

func (self *RepositoryMock) FindByURL(destino string) *domain.Url {
	self.FindByURLInvokedCount++
	return self.FindByURLFn(destino)
}

func (self *RepositoryMock) AddLogClick(id string) {
	self.AddLogClickInvokedCount++
	self.AddLogClickFn(id)
}

func (self *RepositoryMock) FindLogClicks(id string) int {
	self.FindLogClicksInvokedCount++
	return self.FindLogClicksFn(id)
}
