package client

import (
	"github.com/golang-encurtador-url/domain"
	"github.com/golang-encurtador-url/internal/infrastructure/storage"
)

type UrlClientMock struct {
	RepositoryMock              storage.RepositoryMock
	FindInvokedCount            int
	AddLogClickInvokedCount     int
	FindOrCreateURLInvokedCount int

	FindFn            func(ID string) *domain.Url
	AddLogClickFn     func(ID string)
	FindOrCreateURLFn func(repositoryMock *storage.RepositoryMock, destination string, invokedCount *int) (*domain.Url, bool, error)
}

func (self *UrlClientMock) Find(ID string) *domain.Url {
	self.FindInvokedCount++
	return self.FindFn(ID)
}

func (self *UrlClientMock) AddLogClick(ID string) {
	self.AddLogClickInvokedCount++
	self.AddLogClickFn(ID)
}

func (self *UrlClientMock) FindOrCreateURL(destination string) (*domain.Url, bool, error) {
	self.FindOrCreateURLInvokedCount++
	return self.FindOrCreateURLFn(&self.RepositoryMock, destination, &self.FindOrCreateURLInvokedCount)
}
