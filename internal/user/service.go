package user

import (
	"context"
	"rest_api/pkg/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateUserDto) (u User, r error) {
	return
}
