package mocks

import (
	"github.com/golang-encurtador-url/internal/app/urlentities"
	"github.com/stretchr/testify/mock"
)

type ShortenServiceMock struct {
	mock.Mock
}

func (m *ShortenServiceMock) FindOrCreateURL(destination string) (*urlentities.Url, bool, error) {
	args := m.Called(destination)
	arg0 := args.Get(0)
	if arg0 != nil {
		return arg0.(*urlentities.Url), args.Bool(1), args.Error(2)
	}
	return nil, args.Bool(1), args.Error(2)
}
