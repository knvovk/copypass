package service

import (
	"github.com/knvovk/copypass/internal/data"
	"github.com/knvovk/copypass/internal/storage"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	repo *storage.UserStorage
	log  *logrus.Logger
}

func NewUserService(repo *storage.UserStorage, log *logrus.Logger) *UserService {
	return &UserService{repo: repo, log: log}
}

func (s *UserService) Create(user data.User) (data.User, error) {
	passwordHash := user.Password
	_user := mapUserDomain(user)
	_user.PasswordHash = passwordHash
	inserted, err := s.repo.Insert(_user)
	if err != nil {
		s.log.Errorf("Operation CREATE USER failed: %v", err)
		return data.User{}, err
	}
	user = mapUserData(inserted, false)
	s.log.Infof("Operation CREATE USER done: %s", user.Id)
	return user, nil
}

func (s *UserService) GetOne(id string, unsafe bool) (data.User, error) {
	_user, err := s.repo.Find(id)
	if err != nil {
		s.log.Errorf("Operation GET USER failed: %v", err)
		s.log.Errorf("Requested id: %s", id)
		return data.User{}, err
	}
	user := mapUserData(_user, true)
	s.log.Debugf("Operation GET USER done: %v", user)
	return user, nil
}

func (s *UserService) GetMany(limit, offset int) ([]data.User, error) {
	_users, err := s.repo.FindAll(limit, offset)
	if err != nil {
		s.log.Errorf("Operation GET USERS failed: %v", err)
		return nil, err
	}
	users := make([]data.User, 0)
	for _, _user := range _users {
		user := mapUserData(_user, false)
		users = append(users, user)
	}
	s.log.Debugf("Operation GET USERS done. Total: %d", len(users))
	return users, nil
}

func (s *UserService) Update(user data.User) (data.User, error) {
	_user := mapUserDomain(user)
	updated, err := s.repo.Update(_user)
	if err != nil {
		s.log.Errorf("Operation UPDATE USER failed: %v", err)
		return data.User{}, err
	}
	user = mapUserData(updated, false)
	s.log.Infof("Operation UPDATE USER done: %v", user)
	return user, nil
}

func (s *UserService) Delete(user data.User) error {
	if err := s.repo.Delete(user.Id); err != nil {
		s.log.Errorf("Operation DELETE USER failed: %v", err)
		return err
	}
	s.log.Infof("Operation DELETE USER done: %v", user)
	return nil
}

func mapUserData(user storage.User, unsafe bool) data.User {
	_user := data.User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
	if unsafe {
		_user.Password = user.PasswordHash
	}
	return _user
}

func mapUserDomain(user data.User) storage.User {
	return storage.User{
		Id:           user.Id,
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.Password,
	}
}
