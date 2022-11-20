package services

import (
	"github.com/knvovk/copypass/internal/dto"
	"github.com/knvovk/copypass/internal/models"
	"github.com/knvovk/copypass/internal/storages"
	"log"
)

type AccountService struct {
	repo *storages.AccountStorage
}

func NewAccountService(repo *storages.AccountStorage) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) Create(account dto.Account) (dto.Account, error) {
	_account := accountModel(account)
	inserted, err := s.repo.Insert(_account)
	if err != nil {
		log.Printf("Operation CREATE ACCOUNT failed: %v\n", err)
		return dto.Account{}, nil
	}
	account = accountDto(inserted)
	log.Printf("Operation CREATE ACCOUNT done: %s\n", account.Id)
	return account, nil
}

func (s *AccountService) GetOne(id string) (dto.Account, error) {
	_account, err := s.repo.Find(id)
	if err != nil {
		log.Printf("Operation GET ACCOUNT failed: %v\n", err)
		log.Printf("Requested id: %s\n", id)
		return dto.Account{}, nil
	}
	account := accountDto(_account)
	log.Printf("Operation GET ACCOUNT done: %v\n", account)
	return account, nil
}

func (s *AccountService) GetMany(limit, offset int) ([]dto.Account, error) {
	_accounts, err := s.repo.FindAll(limit, offset)
	if err != nil {
		log.Printf("Operation GET ACCOUNTS failed: %v\n", err)
		return nil, err
	}
	accounts := make([]dto.Account, 0)
	for _, _account := range _accounts {
		account := accountDto(_account)
		accounts = append(accounts, account)
	}
	log.Printf("Operation GET ACCOUNTS done. Total: %d\n", len(accounts))
	return accounts, nil
}

func (s *AccountService) Update(account dto.Account) (dto.Account, error) {
	_account := accountModel(account)
	updated, err := s.repo.Update(_account)
	if err != nil {
		log.Printf("Operation UPDATE ACCOUNT failed: %v\n", err)
		return dto.Account{}, nil
	}
	account = accountDto(updated)
	log.Printf("Operation UPDATE ACCOUNT done: %v\n", account)
	return account, nil
}

func (s *AccountService) Delete(account dto.Account) error {
	if err := s.repo.Delete(account.Id); err != nil {
		log.Printf("Operation DELETE ACCOUNT failed: %v\n", err)
		return err
	}
	log.Printf("Operation DELETE ACCOUNT done: %v\n", account)
	return nil
}

func accountDto(account models.Account) dto.Account {
	return dto.Account{
		Id:          account.Id,
		UserId:      account.User.Id,
		Name:        account.Name,
		Description: account.Description,
		Url:         account.Url,
		Username:    account.Username,
		Password:    account.Password,
	}
}

func accountModel(account dto.Account) models.Account {
	return models.Account{
		Id:          account.Id,
		User:        models.User{Id: account.UserId},
		Name:        account.Name,
		Description: account.Description,
		Url:         account.Url,
		Username:    account.Username,
		Password:    account.Password,
	}
}
