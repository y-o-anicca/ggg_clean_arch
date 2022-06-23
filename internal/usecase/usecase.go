package usecase

import (
	"context"

	"github.com/y-o-anicca/ggg_clean_arch/internal/domain/model"
	"github.com/y-o-anicca/ggg_clean_arch/internal/domain/repository"
	"github.com/y-o-anicca/ggg_clean_arch/internal/util/logger"
)

type (
	Usecase interface {
		// user
		GetUser(ctx context.Context, userID int) (*model.User, error)
	}

	usecase struct {
		logger     *logger.Logger
		repository repository.Repository
	}
)

func NewUseCase(
	logger *logger.Logger, repository repository.Repository,
) Usecase {
	return &usecase{
		logger:     logger,
		repository: repository,
	}
}
