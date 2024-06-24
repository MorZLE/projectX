package postgres

import (
	"context"
	"database/sql"
	"errors"
)

func (r *PostgresDB) AddUser(ctx context.Context, user string, chatID int64) error {
	_, err := r.db.Exec("INSERT INTO users (name, chat_id) VALUES ($1, $2)", user, chatID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresDB) GetUser(ctx context.Context, user string) (chatID int64, err error) {
	row := r.db.QueryRowContext(ctx, "SELECT chat_id FROM users WHERE name = $1", user)
	err = row.Scan(&chatID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return 0, ErrNoRows
		default:
			return 0, err
		}
	}
	return chatID, nil
}

func (r *PostgresDB) GetAllUsers(ctx context.Context) ([]int64, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT chat_id FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []int64
	for rows.Next() {
		var user int64
		if err := rows.Scan(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
