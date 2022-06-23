package mysql

import (
	"database/sql"

	"github.com/y-o-anicca/ggg_clean_arch/internal/domain/repository"
	"github.com/y-o-anicca/ggg_clean_arch/internal/util/logger"
)

type client struct {
	logger *logger.Logger
	db     *sql.DB
}

func NewClient(
	logger *logger.Logger, db *sql.DB,
) repository.Repository {
	return &client{
		logger: logger,
		db:     db,
	}
}
