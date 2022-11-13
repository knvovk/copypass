package domain

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Account struct {
	Id          string
	User        User
	Name        string
	Description string
	Url         string
	Username    string
	Password    string
}

func NewAccountRepository(pool *pgxpool.Pool) *AccountRepository {
	return &AccountRepository{pool: pool}
}

type AccountRepository struct {
	pool *pgxpool.Pool
}

func (r *AccountRepository) Insert(account Account) (Account, error) {
	query := `
		INSERT INTO "account" (user_id, name, description, 
							   url, username, password)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`
	var id string
	args := []any{
		account.User.Id, account.Name, account.Description,
		account.Url, account.Username, account.Password,
	}
	row := r.pool.QueryRow(context.Background(), query, args...)
	if err := row.Scan(&id); err != nil {
		return Account{}, nil
	}
	account.Id = id
	return account, nil
}

func (r *AccountRepository) Find(id string) (Account, error) {
	query := `
		SELECT id, user_id, name, description, url, username, password
		  FROM "account"
		 WHERE id = $1;
	`
	var a = Account{}
	row := r.pool.QueryRow(context.Background(), query, id)
	args := []any{
		&a.Id, &a.User.Id, &a.Name, &a.Description,
		&a.Url, &a.Username, &a.Password,
	}
	if err := row.Scan(args...); err != nil {
		return Account{}, nil
	}
	return a, nil
}

func (r *AccountRepository) FindByUser(user User) (Account, error) {
	query := `
		SELECT id, user_id, name, description, url, username, password
		  FROM "account"
		 WHERE user_id = $1;
	`
	var account = Account{}
	row := r.pool.QueryRow(context.Background(), query, user.Id)
	args := []any{
		&account.Id, &account.User.Id, &account.Name, &account.Description,
		&account.Url, &account.Username, &account.Password,
	}
	if err := row.Scan(args...); err != nil {
		return Account{}, nil
	}
	return account, nil
}

func (r *AccountRepository) FindByName(name string) (Account, error) {
	query := `
		SELECT id, user_id, name, description, url, username, password
		  FROM "account"
		 WHERE name = $1;
	`
	var a = Account{}
	row := r.pool.QueryRow(context.Background(), query, name)
	args := []any{
		&a.Id, &a.User.Id, &a.Name, &a.Description,
		&a.Url, &a.Username, &a.Password,
	}
	if err := row.Scan(args...); err != nil {
		return Account{}, nil
	}
	return a, nil
}

func (r *AccountRepository) FindAll(limit, offset int) ([]Account, error) {
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
	accounts := make([]Account, 0)
	for rows.Next() {
		var account = Account{}
		args := []any{
			&account.Id, &account.User.Id, &account.Name, &account.Description,
			&account.Url, &account.Username, &account.Password,
		}
		if err := rows.Scan(args...); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (r *AccountRepository) Update(account Account) (Account, error) {
	query := `
		UPDATE "account"
		   SET name = $1, description = $2, url = $3, 
		   	   username = $4, password = $5
		 WHERE id = $6;
	`
	args := []any{
		&account.Name, &account.Description, &account.Url,
		&account.Username, &account.Password, &account.Id,
	}
	if _, err := r.pool.Exec(context.Background(), query, args...); err != nil {
		return Account{}, nil
	}
	// TODO: Return Updated Instance
	return account, nil
}

func (r *AccountRepository) Delete(id string) error {
	query := `DELETE FROM "account" WHERE id = $1;`
	if _, err := r.pool.Exec(context.Background(), query, id); err != nil {
		return err
	}
	return nil
}
