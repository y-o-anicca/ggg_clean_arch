package repository

import (
	"context"

	"github.com/y-o-anicca/ggg_clean_arch/internal/domain/model"
)

type Repository interface {
	GetUser(ctx context.Context, userID int) (*model.User, error)
}
