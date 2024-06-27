package postgres

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"projectX/msrvs/msrv-bot-tg/config"
)

var ErrNoRows = sql.ErrNoRows

func InitRepository(cnf *config.Config) IRepository {

	db, err := sql.Open("postgres", cnf.DB.Dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &PostgresDB{db: db}
}

type IRepository interface {
	AddUser(ctx context.Context, user string, chatID int64, group string) error
	GetUser(ctx context.Context, user string) (chatID int64, err error)
	GetAllUsers(ctx context.Context) ([]int64, error)
	GetUserByGroup(ctx context.Context, group string) ([]int64, error)
}

type PostgresDB struct {
	db *sql.DB
}
