package postgres

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"projectX/msrvs/msrv-bot-tg/config"
)

const (
	createUsers = "CREATE TABLE IF NOT EXISTS users (" +
		"\n    id SERIAL PRIMARY KEY," +
		"\n    name VARCHAR(255) NOT NULL," +
		"\n    chat_id integer NOT NULL," +
		"\n    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP," +
		"\n    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP\n    " +
		");"
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

	_, err = db.Exec(createUsers)
	if err != nil {
		panic(err)
	}
	return &PostgresDB{db: db}
}

type IRepository interface {
	AddUser(ctx context.Context, user string, chatID int64) error
	GetUser(ctx context.Context, user string) (chatID int64, err error)
	GetAllUsers(ctx context.Context) ([]int64, error)
}

type PostgresDB struct {
	db *sql.DB
}
