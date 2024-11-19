package repository

import (
	"context"
	"database/sql"
	"victorbetoni/trabalho-web/internal/infra/db"
)

type LoginRepository struct {
	dbConn *sql.DB
	*db.Queries
}

func NewLoginRepository(dbConn *sql.DB) *LoginRepository {
	return &LoginRepository{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (r *LoginRepository) FindLogin(ctx context.Context, username string) (string, error) {
	return r.Queries.FindLogin(ctx, username)
}
