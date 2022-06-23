package mysql

import (
	"context"
	"time"

	"github.com/y-o-anicca/ggg_clean_arch/internal/domain/model"
)

type UserDTO struct {
	ID      int
	Name    string
	Created time.Time
}

func (c client) GetUser(ctx context.Context, userID int) (*model.User, error) {
	stmt := "SELECT * FROM users WHERE id = ?"

	userDTO := UserDTO{}
	err := c.db.QueryRow(stmt, userID).Scan(&userDTO)
	if err != nil {
		// c.logger.Error()
		return nil, err
	}

	return &model.User{
		ID:   userDTO.ID,
		Name: userDTO.Name,
	}, nil
}
