package service

import (
	"context"

	"github.com/aakashkaji/empolyee-go/app/internal/domain"
	"github.com/stretchr/testify/mock"
)

type EmpolyeeMockService struct {
	mock.Mock
}

func (m *EmpolyeeMockService) CreateEmpolyee(_ context.Context, req domain.Empolyee) (*domain.Empolyee, error) {

	return nil, nil

}
