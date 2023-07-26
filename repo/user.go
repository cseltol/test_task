package repo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v5"
	"test_task/model"
)

func Create(c *pgx.Conn, u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeUserCreation(); err != nil {
		return err
	}

	return c.QueryRow(
		context.Background(),
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email, u.EncryptedPassword,
	).Scan(&u.ID)
}

func FindByID(c *pgx.Conn, id uint32) (*model.User, error) {
	u := &model.User{}

	if err := c.QueryRow(
		context.Background(),
		"SELECT id, email, encrypted_password FROM users WHERE id = ($1)",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("not found")
		}

		return nil, err
	}

	return u, nil
}

func FindByEmail(c *pgx.Conn, email string) (*model.User, error) {
	u := &model.User{}

	if err := c.QueryRow(
		context.Background(),
		"SELECT id, email, encrypted_password FROM users WHERE email = ($1)",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("not found")
		}
		return nil, err
	}

	return u, nil
}
