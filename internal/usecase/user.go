package usecase

import (
	"context"

	"github.com/y-o-anicca/ggg_clean_arch/internal/domain/model"
)

func (u *usecase) GetUser(ctx context.Context, userID int) (*model.User, error) {
	user, err := u.repository.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
