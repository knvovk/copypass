package domain

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Id           string
	Username     string
	Email        string
	PasswordHash string
}

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) Insert(user User) (User, error) {
	query := `
		   INSERT INTO "user" (username, email, password_hash)
		   VALUES ($1, $2, $3)
		RETURNING id;
	`
	var id string
	args := []any{user.Username, user.Email, user.PasswordHash}
	row := r.pool.QueryRow(context.Background(), query, args...)
	if err := row.Scan(&id); err != nil {
		return User{}, err
	}
	user.Id = id
	return user, nil
}

func (r *UserRepository) Find(id string) (User, error) {
	query := `
		SELECT id, username, email, password_hash
		  FROM "user"
		 WHERE id = $1;
	`
	var user = User{}
	row := r.pool.QueryRow(context.Background(), query, id)
	args := []any{
		&user.Id, &user.Username, &user.Email, &user.PasswordHash,
	}
	if err := row.Scan(args...); err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *UserRepository) FindByUsername(username string) (User, error) {
	query := `
		SELECT id, username, email, password_hash
		  FROM "user"
		 WHERE username = $1;
	`
	var user = User{}
	row := r.pool.QueryRow(context.Background(), query, username)
	args := []any{
		&user.Id, &user.Username, &user.Email, &user.PasswordHash,
	}
	if err := row.Scan(args...); err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (User, error) {
	query := `
		SELECT id, username, email, password_hash
		  FROM "user"
		 WHERE email = $1;
	`
	var user = User{}
	row := r.pool.QueryRow(context.Background(), query, email)
	args := []any{
		&user.Id, &user.Username, &user.Email, &user.PasswordHash,
	}
	if err := row.Scan(args...); err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *UserRepository) FindAll(limit, offset int) ([]User, error) {
	query := `
		SELECT id, username, email
		  FROM "user"
		 LIMIT $1
		OFFSET $2;
	`
	rows, err := r.pool.Query(context.Background(), query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]User, 0)
	for rows.Next() {
		var user = User{}
		args := []any{&user.Id, &user.Username, &user.Email}
		if err = rows.Scan(args...); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) Update(user User) (User, error) {
	query := `
		UPDATE "user" 
		   SET username = $1, email = $2, password_hash = $3 
		 WHERE id = $4;
	`
	args := []any{user.Username, user.Email, user.PasswordHash, user.Id}
	if _, err := r.pool.Exec(context.Background(), query, args...); err != nil {
		return User{}, err
	}
	// TODO: Return Updated Instance
	return user, nil
}

func (r *UserRepository) Delete(id string) error {
	query := `DELETE FROM "user" WHERE id = $1`
	if _, err := r.pool.Exec(context.Background(), query, id); err != nil {
		return err
	}
	return nil
}
