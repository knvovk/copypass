package service

import (
	"github.com/knvovk/copypass/internal/data"
	"github.com/knvovk/copypass/internal/storage"
	"log"
)

type AccountService struct {
	repo *storage.AccountStorage
}

func NewAccountService(repo *storage.AccountStorage) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) Create(account data.Account) (data.Account, error) {
	_account := mapAccountDomain(account)
	inserted, err := s.repo.Insert(_account)
	if err != nil {
		log.Printf("Operation CREATE ACCOUNT failed: %v\n", err)
		return data.Account{}, nil
	}
	account = mapAccountData(inserted)
	log.Printf("Operation CREATE ACCOUNT done: %s\n", account.Id)
	return account, nil
}

func (s *AccountService) GetOne(id string) (data.Account, error) {
	_account, err := s.repo.Find(id)
	if err != nil {
		log.Printf("Operation GET ACCOUNT failed: %v\n", err)
		log.Printf("Requested id: %s\n", id)
		return data.Account{}, nil
	}
	account := mapAccountData(_account)
	log.Printf("Operation GET ACCOUNT done: %v\n", account)
	return account, nil
}

func (s *AccountService) GetMany(limit, offset int) ([]data.Account, error) {
	_accounts, err := s.repo.FindAll(limit, offset)
	if err != nil {
		log.Printf("Operation GET ACCOUNTS failed: %v\n", err)
		return nil, err
	}
	accounts := make([]data.Account, 0)
	for _, _account := range _accounts {
		account := mapAccountData(_account)
		accounts = append(accounts, account)
	}
	log.Printf("Operation GET ACCOUNTS done. Total: %d\n", len(accounts))
	return accounts, nil
}

func (s *AccountService) Update(account data.Account) (data.Account, error) {
	_account := mapAccountDomain(account)
	updated, err := s.repo.Update(_account)
	if err != nil {
		log.Printf("Operation UPDATE ACCOUNT failed: %v\n", err)
		return data.Account{}, nil
	}
	account = mapAccountData(updated)
	log.Printf("Operation UPDATE ACCOUNT done: %v\n", account)
	return account, nil
}

func (s *AccountService) Delete(account data.Account) error {
	if err := s.repo.Delete(account.Id); err != nil {
		log.Printf("Operation DELETE ACCOUNT failed: %v\n", err)
		return err
	}
	log.Printf("Operation DELETE ACCOUNT done: %v\n", account)
	return nil
}

func mapAccountData(account storage.Account) data.Account {
	return data.Account{
		Id:          account.Id,
		UserId:      account.User.Id,
		Name:        account.Name,
		Description: account.Description,
		Url:         account.Url,
		Username:    account.Username,
		Password:    account.Password,
	}
}

func mapAccountDomain(account data.Account) storage.Account {
	return storage.Account{
		Id:          account.Id,
		User:        storage.User{Id: account.UserId},
		Name:        account.Name,
		Description: account.Description,
		Url:         account.Url,
		Username:    account.Username,
		Password:    account.Password,
	}
}
