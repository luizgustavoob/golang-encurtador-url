package mocks

import (
	"github.com/golang-encurtador-url/internal/app/urlentities"
	"github.com/stretchr/testify/mock"
)

type RedirectServiceMock struct {
	AddStatisticsInvokedCount int
	AddStatisticsFn           func(url *urlentities.Url)
	mock.Mock
}

func (m *RedirectServiceMock) AddStatistics(url *urlentities.Url) {
	m.AddStatisticsInvokedCount++
	m.AddStatisticsFn(url)
}

func (m *RedirectServiceMock) Find(id string) *urlentities.Url {
	args := m.Called(id)
	arg0 := args.Get(0)
	if arg0 != nil {
		return arg0.(*urlentities.Url)
	}
	return nil
}
