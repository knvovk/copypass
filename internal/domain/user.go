package domain

import (
	"database/sql"
)

type User struct {
	Id           string
	Username     string
	Email        string
	PasswordHash string
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Insert(user User) (User, error) {
	stmt, err := r.db.Prepare(`
		INSERT INTO "user" (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id;
	`)
	if err != nil {
		return user, err
	}
	args := []any{user.Username, user.Email, user.PasswordHash}
	if err := stmt.QueryRow(args...).Scan(&user.Id); err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) Find(id string) (User, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, username, email, password_hash
		FROM "user"
		WHERE id = $1;
	`)
	var user = User{}
	if err != nil {
		return user, err
	}
	args := []any{
		&user.Id, &user.Username, &user.Email, &user.PasswordHash,
	}
	if err := stmt.QueryRow(id).Scan(args...); err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) FindByUsername(username string) (User, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, username, email
		FROM "user"
		WHERE username = $1;
	`)
	var user = User{}
	if err != nil {
		return user, err
	}
	args := []any{&user.Id, &user.Username, &user.Email}
	if err := stmt.QueryRow(username).Scan(args...); err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (User, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, username, email
		FROM "user"
		WHERE email = $1;
	`)
	var user = User{}
	if err != nil {
		return user, err
	}
	args := []any{&user.Id, &user.Username, &user.Email}
	if err := stmt.QueryRow(email).Scan(args...); err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) FindAll(limit, offset int) ([]User, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, username, email
		FROM "user"
		LIMIT $1
		OFFSET $2;
	`)
	users := make([]User, 0)
	if err != nil {
		return users, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(limit, offset)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var user = User{}
		args := []any{&user.Id, &user.Username, &user.Email}
		if err := rows.Scan(args...); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) Update(user User) (User, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return user, err
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare(`
		UPDATE "user"
		SET username = $1, email = $2, password_hash = $3
		WHERE id = $4;
	`)
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	args := []any{
		user.Username, user.Email,
		user.PasswordHash, user.Id,
	}
	if _, err := stmt.Exec(args...); err != nil {
		return user, err
	}
	stmt, err = tx.Prepare(`
		SELECT id, username, email, password_hash
		FROM "user"
		WHERE id = $1;
	`)
	if err != nil {
		return user, err
	}
	args = []any{
		&user.Id, &user.Username, &user.Email, &user.PasswordHash,
	}
	if err := stmt.QueryRow(user.Id).Scan(args...); err != nil {
		return user, err
	}
	if err := tx.Commit(); err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) Delete(id string) error {
	stmt, err := r.db.Prepare(`DELETE FROM "user" WHERE id = $1`)
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(id); err != nil {
		return err
	}
	return nil
}
