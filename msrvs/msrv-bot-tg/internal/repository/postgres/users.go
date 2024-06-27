package postgres

import (
	"context"
	"database/sql"
	"errors"
)

func (r *PostgresDB) AddUser(ctx context.Context, user string, chatID int64, group string) error {
	query := `INSERT INTO users (name, chat_id, group_id) 
			VALUES ($1, $2, (SELECT id FROM groups WHERE name = $3))`

	_, err := r.db.Exec(query, user, chatID, group)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresDB) GetUser(ctx context.Context, user string) (chatID int64, err error) {
	query := `SELECT chat_id FROM users  WHERE name = $1`

	row := r.db.QueryRowContext(ctx, query, user)
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

func (r *PostgresDB) GetUserByGroup(ctx context.Context, group string) ([]int64, error) {
	query := `
		SELECT chat_id
		FROM users
		WHERE group_id = (
			SELECT id
			FROM groups
			WHERE name = $1
		)
	`
	rows, err := r.db.QueryContext(ctx, query, group)
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
