package storages

import (
	"database/sql"
	"github.com/knvovk/copypass/internal/models"
)

func NewAccountStorage(db *sql.DB) *AccountStorage {
	return &AccountStorage{db: db}
}

type AccountStorage struct {
	db *sql.DB
}

func (s *AccountStorage) Insert(account models.Account) (models.Account, error) {
	stmt, err := s.db.Prepare(`
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

func (s *AccountStorage) Find(id string) (models.Account, error) {
	stmt, err := s.db.Prepare(`
		SELECT id, user_id, name, description, url, 
			username, password
		FROM "account"
		WHERE id = $1;
	`)
	var account = models.Account{}
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

func (s *AccountStorage) FindByUser(user models.User) (models.Account, error) {
	stmt, err := s.db.Prepare(`
		SELECT id, user_id, name, description, url, 
			username, password
		FROM "account"
		WHERE user_id = $1;
	`)
	var account = models.Account{}
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

func (s *AccountStorage) FindByName(name string) (models.Account, error) {
	stmt, err := s.db.Prepare(`
		SELECT id, user_id, name, description, url, 
			username, password
		FROM "account"
		WHERE name = $1;
	`)
	var account = models.Account{}
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

func (s *AccountStorage) FindAll(limit, offset int) ([]models.Account, error) {
	stmt, err := s.db.Prepare(`
		SELECT id, user_id, name, description, url, 
			username, password
		FROM "account"
		LIMIT $1
		OFFSET $2;
	`)
	accounts := make([]models.Account, 0)
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
		var account = models.Account{}
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

func (s *AccountStorage) Update(account models.Account) (models.Account, error) {
	tx, err := s.db.Begin()
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

func (s *AccountStorage) Delete(id string) error {
	stmt, err := s.db.Prepare(`DELETE FROM "account" WHERE id = $1;`)
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(id); err != nil {
		return err
	}
	return nil
}
