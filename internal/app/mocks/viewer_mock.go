package mocks

import (
	"github.com/golang-encurtador-url/internal/app/urlentities"
	"github.com/stretchr/testify/mock"
)

type ViewerServiceMock struct {
	mock.Mock
}

func (m *ViewerServiceMock) Find(id string) *urlentities.Url {
	args := m.Called(id)
	arg0 := args.Get(0)
	if arg0 != nil {
		return args.Get(0).(*urlentities.Url)
	}
	return nil
}

func (m *ViewerServiceMock) GetStatistics(url *urlentities.Url) *urlentities.Statistics {
	args := m.Called(url)
	return args.Get(0).(*urlentities.Statistics)
}
