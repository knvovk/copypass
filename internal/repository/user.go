package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/knvovk/copypass/internal/domain"
)

var _ domain.UserRepository = (*repository)(nil)

func NewUserRepository(pool *pgxpool.Pool) domain.UserRepository {
	return &repository{pool: pool}
}

type repository struct {
	pool *pgxpool.Pool
}

func (r *repository) Insert(user domain.User) (*domain.User, error) {
	query := `
		   INSERT INTO "user" (username, email, password_hash)
		   VALUES ($1, $2, $3)
		RETURNING id;
	`
	var id string
	args := []any{user.Username, user.Email, user.PasswordHash}
	row := r.pool.QueryRow(context.Background(), query, args...)
	if err := row.Scan(&id); err != nil {
		return nil, err
	}
	user.Id = id
	return &user, nil
}

func (r *repository) Find(id string) (*domain.User, error) {
	query := `
		SELECT id, username, email, password_hash
		  FROM "user"
		 WHERE id = $1;
	`
	var user = new(domain.User)
	row := r.pool.QueryRow(context.Background(), query, id)
	args := []any{&user.Id, &user.Username, &user.Email, &user.PasswordHash}
	err := row.Scan(args...)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) FindByUsername(username string) (*domain.User, error) {
	query := `
		SELECT id, username, email, password_hash
		  FROM "user"
		 WHERE username = $1;
	`
	var user = new(domain.User)
	row := r.pool.QueryRow(context.Background(), query, username)
	args := []any{&user.Id, &user.Username, &user.Email, &user.PasswordHash}
	err := row.Scan(args...)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (*domain.User, error) {
	query := `
		SELECT id, username, email, password_hash
		  FROM "user"
		 WHERE email = $1;
	`
	var user = new(domain.User)
	row := r.pool.QueryRow(context.Background(), query, email)
	args := []any{&user.Id, &user.Username, &user.Email, &user.PasswordHash}
	err := row.Scan(args...)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) FindAll(limit, offset int) ([]*domain.User, error) {
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
	users := make([]*domain.User, 0)
	for rows.Next() {
		var user = new(domain.User)
		args := []any{&user.Id, &user.Username, &user.Email}
		if err = rows.Scan(args...); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *repository) Update(user domain.User) error {
	query := `
		UPDATE "user" 
		   SET username = $1, email = $2, password_hash = $3 
		 WHERE id = $4;
	`
	args := []any{user.Username, user.Email, user.PasswordHash, user.Id}
	if _, err := r.pool.Exec(context.Background(), query, args...); err != nil {
		return err
	}
	return nil
}

func (r *repository) Delete(id string) error {
	query := `DELETE FROM "user" WHERE id = $1`
	if _, err := r.pool.Exec(context.Background(), query, id); err != nil {
		return err
	}
	return nil
}
