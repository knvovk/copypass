package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/knvovk/copypass/internal/domain"
)

var _ domain.AccountRepository = (*accountRepository)(nil)

func NewAccountRepository(pool *pgxpool.Pool) domain.AccountRepository {
	return &accountRepository{pool: pool}
}

type accountRepository struct {
	pool *pgxpool.Pool
}

func (r *accountRepository) Insert(a domain.Account) (domain.Account, error) {
	query := `
		INSERT INTO "account" (user_id, name, description, 
							   url, username, password)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`
	var id string
	args := []any{
		a.User.Id, a.Name, a.Description, a.Url, a.Username, a.Password,
	}
	row := r.pool.QueryRow(context.Background(), query, args...)
	if err := row.Scan(&id); err != nil {
		return domain.Account{}, nil
	}
	a.Id = id
	return a, nil
}

func (r *accountRepository) Find(id string) (domain.Account, error) {
	query := `
		SELECT id, user_id, name, description, url, username, password
		  FROM "account"
		 WHERE id = $1;
	`
	var a = domain.Account{}
	row := r.pool.QueryRow(context.Background(), query, id)
	args := []any{
		&a.Id, &a.User.Id, &a.Name, &a.Description,
		&a.Url, &a.Username, &a.Password,
	}
	if err := row.Scan(args...); err != nil {
		return domain.Account{}, nil
	}
	return a, nil
}

func (r *accountRepository) FindByUser(user domain.User) (domain.Account, error) {
	query := `
		SELECT id, user_id, name, description, url, username, password
		  FROM "account"
		 WHERE user_id = $1;
	`
	var a = domain.Account{}
	row := r.pool.QueryRow(context.Background(), query, user.Id)
	args := []any{
		&a.Id, &a.User.Id, &a.Name, &a.Description,
		&a.Url, &a.Username, &a.Password,
	}
	if err := row.Scan(args...); err != nil {
		return domain.Account{}, nil
	}
	return a, nil
}

func (r *accountRepository) FindByName(name string) (domain.Account, error) {
	query := `
		SELECT id, user_id, name, description, url, username, password
		  FROM "account"
		 WHERE name = $1;
	`
	var a = domain.Account{}
	row := r.pool.QueryRow(context.Background(), query, name)
	args := []any{
		&a.Id, &a.User.Id, &a.Name, &a.Description,
		&a.Url, &a.Username, &a.Password,
	}
	if err := row.Scan(args...); err != nil {
		return domain.Account{}, nil
	}
	return a, nil
}

func (r *accountRepository) FindAll(limit, offset int) ([]domain.Account, error) {
	query := `
		SELECT id, user_id, name, description, url, username, password
		  FROM "account"
		 LIMIT $1
		OFFSET $2;
	`
	rows, err := r.pool.Query(context.Background(), query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	accounts := make([]domain.Account, 0)
	for rows.Next() {
		var a = domain.Account{}
		args := []any{
			&a.Id, &a.User.Id, &a.Name, &a.Description,
			&a.Url, &a.Username, &a.Password,
		}
		if err := rows.Scan(args...); err != nil {
			return nil, err
		}
		accounts = append(accounts, a)
	}
	return accounts, nil
}

func (r *accountRepository) Update(a domain.Account) (domain.Account, error) {
	query := `
		UPDATE "account"
		   SET name = $1, description = $2, url = $3, 
		   	   username = $4, password = $5
		 WHERE id = $6;
	`
	args := []any{
		&a.Name, &a.Description, &a.Url, &a.Username, &a.Password, &a.Id,
	}
	if _, err := r.pool.Exec(context.Background(), query, args...); err != nil {
		return domain.Account{}, nil
	}
	// TODO: Return Updated Instance
	return a, nil
}

func (r *accountRepository) Delete(id string) error {
	query := `DELETE FROM "account" WHERE id = $1;`
	if _, err := r.pool.Exec(context.Background(), query, id); err != nil {
		return err
	}
	return nil
}
