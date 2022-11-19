package service

import (
	"github.com/knvovk/copypass/internal/data"
	"github.com/knvovk/copypass/internal/storage"
	"github.com/sirupsen/logrus"
)

type AccountService struct {
	repo *storage.AccountStorage
	log  *logrus.Logger
}

func NewAccountService(repo *storage.AccountStorage, log *logrus.Logger) *AccountService {
	return &AccountService{repo: repo, log: log}
}

func (s *AccountService) Create(account data.Account) (data.Account, error) {
	_account := mapAccountDomain(account)
	inserted, err := s.repo.Insert(_account)
	if err != nil {
		s.log.Errorf("Operation CREATE ACCOUNT failed: %v", err)
		return data.Account{}, nil
	}
	account = mapAccountData(inserted)
	s.log.Infof("Operation CREATE ACCOUNT done: %s", account.Id)
	return account, nil
}

func (s *AccountService) GetOne(id string) (data.Account, error) {
	_account, err := s.repo.Find(id)
	if err != nil {
		s.log.Errorf("Operation GET ACCOUNT failed: %v", err)
		s.log.Errorf("Requested id: %s", id)
		return data.Account{}, nil
	}
	account := mapAccountData(_account)
	s.log.Debugf("Operation GET ACCOUNT done: %v", account)
	return account, nil
}

func (s *AccountService) GetMany(limit, offset int) ([]data.Account, error) {
	_accounts, err := s.repo.FindAll(limit, offset)
	if err != nil {
		s.log.Errorf("Operation GET ACCOUNTS failed: %v", err)
		return nil, err
	}
	accounts := make([]data.Account, 0)
	for _, _account := range _accounts {
		account := mapAccountData(_account)
		accounts = append(accounts, account)
	}
	s.log.Debugf("Operation GET ACCOUNTS done. Total: %d", len(accounts))
	return accounts, nil
}

func (s *AccountService) Update(account data.Account) (data.Account, error) {
	_account := mapAccountDomain(account)
	updated, err := s.repo.Update(_account)
	if err != nil {
		s.log.Errorf("Operation UPDATE ACCOUNT failed: %v", err)
		return data.Account{}, nil
	}
	account = mapAccountData(updated)
	s.log.Infof("Operation UPDATE ACCOUNT done: %v", account)
	return account, nil
}

func (s *AccountService) Delete(account data.Account) error {
	if err := s.repo.Delete(account.Id); err != nil {
		s.log.Errorf("Operation DELETE ACCOUNT failed: %v", err)
		return err
	}
	s.log.Infof("Operation DELETE ACCOUNT done: %v", account)
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
