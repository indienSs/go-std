package postgres

import (
	"github.com/indienSs/go-std/internal/models"
	"github.com/indienSs/go-std/internal/types"
)

func (pg *Postgres) GetUsers() ([]models.User, error) {
	rows, err := pg.db.Query("SELECT id, username, email, createdAt, updatedAt FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (pg *Postgres) GetUser(id int) (models.User, error) {
	var user models.User
	err := pg.db.QueryRow("SELECT id, username, email, createdAt, updatedAt FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	
	if err != nil {
		return user, err
	}

	return user, nil
}

func (pg *Postgres) CreateUser(user *models.User) error {
	err := pg.db.QueryRow(
		"INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id, createdAt, updatedAt",
		user.Username, user.Email,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (pg *Postgres) UpdateUser(id int, user *models.User) error {
	err := pg.db.QueryRow(
		"UPDATE users SET username = $1, email = $2, updatedAt = NOW() WHERE id = $3 RETURNING createdAt, updatedAt",
		user.Username, user.Email, id,
	).Scan(&user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (pg *Postgres) DeleteUser(id int) error {
	result, err := pg.db.Exec("DELETE FROM users WHERE id = $1", id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return types.ErrNotFound
	}

	return nil
}