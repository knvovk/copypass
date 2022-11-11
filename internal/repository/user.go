package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/knvovk/copypass/internal/domain"
)

var _ domain.UserRepository = (*userRepository)(nil)

func NewUserRepository(pool *pgxpool.Pool) domain.UserRepository {
	return &userRepository{pool: pool}
}

type userRepository struct {
	pool *pgxpool.Pool
}

func (r *userRepository) Insert(u domain.User) (domain.User, error) {
	query := `
		   INSERT INTO "user" (username, email, password_hash)
		   VALUES ($1, $2, $3)
		RETURNING id;
	`
	var id string
	args := []any{u.Username, u.Email, u.PasswordHash}
	row := r.pool.QueryRow(context.Background(), query, args...)
	if err := row.Scan(&id); err != nil {
		return domain.User{}, err
	}
	u.Id = id
	return u, nil
}

func (r *userRepository) Find(id string) (domain.User, error) {
	query := `
		SELECT id, username, email, password_hash
		  FROM "user"
		 WHERE id = $1;
	`
	var user = domain.User{}
	row := r.pool.QueryRow(context.Background(), query, id)
	args := []any{
		&user.Id, &user.Username, &user.Email, &user.PasswordHash,
	}
	if err := row.Scan(args...); err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *userRepository) FindByUsername(username string) (domain.User, error) {
	query := `
		SELECT id, username, email, password_hash
		  FROM "user"
		 WHERE username = $1;
	`
	var user = domain.User{}
	row := r.pool.QueryRow(context.Background(), query, username)
	args := []any{
		&user.Id, &user.Username, &user.Email, &user.PasswordHash,
	}
	if err := row.Scan(args...); err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *userRepository) FindByEmail(email string) (domain.User, error) {
	query := `
		SELECT id, username, email, password_hash
		  FROM "user"
		 WHERE email = $1;
	`
	var user = domain.User{}
	row := r.pool.QueryRow(context.Background(), query, email)
	args := []any{
		&user.Id, &user.Username, &user.Email, &user.PasswordHash,
	}
	if err := row.Scan(args...); err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *userRepository) FindAll(limit, offset int) ([]domain.User, error) {
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
	users := make([]domain.User, 0)
	for rows.Next() {
		var user = domain.User{}
		args := []any{&user.Id, &user.Username, &user.Email}
		if err = rows.Scan(args...); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) Update(u domain.User) (domain.User, error) {
	query := `
		UPDATE "user" 
		   SET username = $1, email = $2, password_hash = $3 
		 WHERE id = $4;
	`
	args := []any{u.Username, u.Email, u.PasswordHash, u.Id}
	if _, err := r.pool.Exec(context.Background(), query, args...); err != nil {
		return domain.User{}, err
	}
	// TODO: Return Updated Instance
	return u, nil
}

func (r *userRepository) Delete(id string) error {
	query := `DELETE FROM "user" WHERE id = $1`
	if _, err := r.pool.Exec(context.Background(), query, id); err != nil {
		return err
	}
	return nil
}
