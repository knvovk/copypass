package domain

import (
	"database/sql"
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

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

type AccountRepository struct {
	db *sql.DB
}

func (r *AccountRepository) Insert(account Account) (Account, error) {
	stmt, err := r.db.Prepare(`
		INSERT INTO "account" (
			user_id, name, description, 
			url, username, password
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`)
	if err != nil {
		return account, nil
	}
	args := []any{
		account.User.Id, account.Name, account.Description,
		account.Url, account.Username, account.Password,
	}
	if err := stmt.QueryRow(args...).Scan(&account.Id); err != nil {
		return account, err
	}
	return account, nil
}

func (r *AccountRepository) Find(id string) (Account, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, user_id, name, description, url, 
			username, password
		FROM "account"
		WHERE id = $1;
	`)
	var account = Account{}
	if err != nil {
		return account, nil
	}
	args := []any{
		&account.Id, &account.User.Id, &account.Name,
		&account.Description, &account.Url,
		&account.Username, &account.Password,
	}
	if err := stmt.QueryRow(id).Scan(args...); err != nil {
		return account, err
	}
	return account, nil
}

func (r *AccountRepository) FindByUser(user User) (Account, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, user_id, name, description, url, 
			username, password
		FROM "account"
		WHERE user_id = $1;
	`)
	var account = Account{}
	if err != nil {
		return account, nil
	}
	args := []any{
		&account.Id, &account.User.Id, &account.Name,
		&account.Description, &account.Url,
		&account.Username, &account.Password,
	}
	if err := stmt.QueryRow(user.Id).Scan(args...); err != nil {
		return account, err
	}
	return account, nil
}

func (r *AccountRepository) FindByName(name string) (Account, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, user_id, name, description, url, 
			username, password
		FROM "account"
		WHERE name = $1;
	`)
	var account = Account{}
	if err != nil {
		return account, nil
	}
	args := []any{
		&account.Id, &account.User.Id, &account.Name,
		&account.Description, &account.Url,
		&account.Username, &account.Password,
	}
	if err := stmt.QueryRow(name).Scan(args...); err != nil {
		return account, err
	}
	return account, nil
}

func (r *AccountRepository) FindAll(limit, offset int) ([]Account, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, user_id, name, description, url, 
			username, password
		FROM "account"
		LIMIT $1
		OFFSET $2;
	`)
	accounts := make([]Account, 0)
	if err != nil {
		return accounts, nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(limit, offset)
	if err != nil {
		return accounts, nil
	}
	defer rows.Close()
	for rows.Next() {
		var account = Account{}
		args := []any{
			&account.Id, &account.User.Id, &account.Name,
			&account.Description, &account.Url,
			&account.Username, &account.Password,
		}
		if err := rows.Scan(args...); err != nil {
			return accounts, nil
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (r *AccountRepository) Update(account Account) (Account, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return account, err
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare(`
		UPDATE "account"
		SET name = $1, description = $2, url = $3, 
		username = $4, password = $5
		WHERE id = $6;
	`)
	if err != nil {
		return account, err
	}
	defer stmt.Close()
	args := []any{
		account.Name, account.Description, account.Url,
		account.Username, account.Password, account.Id,
	}
	if _, err := stmt.Exec(args...); err != nil {
		return account, err
	}
	stmt, err = tx.Prepare(`
		SELECT id, user_id, name, description, url, 
			username, password
		FROM "account"
		WHERE id = $1;
	`)
	if err != nil {
		return account, err
	}
	args = []any{
		&account.Id, &account.User.Id, &account.Name,
		&account.Description, &account.Url,
		&account.Username, &account.Password,
	}
	if err := stmt.QueryRow(account.Id).Scan(args...); err != nil {
		return account, err
	}
	if err := tx.Commit(); err != nil {
		return account, err
	}
	return account, nil
}

func (r *AccountRepository) Delete(id string) error {
	stmt, err := r.db.Prepare(`DELETE FROM "account" WHERE id = $1;`)
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(id); err != nil {
		return err
	}
	return nil
}
