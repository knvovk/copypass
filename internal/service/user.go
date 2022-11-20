package service

import (
	"github.com/knvovk/copypass/internal/data"
	"github.com/knvovk/copypass/internal/storage"
	"log"
)

type UserService struct {
	repo *storage.UserStorage
}

func NewUserService(repo *storage.UserStorage) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user data.User) (data.User, error) {
	passwordHash := user.Password
	_user := mapUserDomain(user)
	_user.PasswordHash = passwordHash
	inserted, err := s.repo.Insert(_user)
	if err != nil {
		log.Printf("Operation CREATE USER failed: %v\n", err)
		return data.User{}, err
	}
	user = mapUserData(inserted, false)
	log.Printf("Operation CREATE USER done: %s\n", user.Id)
	return user, nil
}

func (s *UserService) GetOne(id string, unsafe bool) (data.User, error) {
	_user, err := s.repo.Find(id)
	if err != nil {
		log.Printf("Operation GET USER failed: %v\n", err)
		log.Printf("Requested id: %s\n", id)
		return data.User{}, err
	}
	user := mapUserData(_user, true)
	log.Printf("Operation GET USER done: %v\n", user)
	return user, nil
}

func (s *UserService) GetMany(limit, offset int) ([]data.User, error) {
	_users, err := s.repo.FindAll(limit, offset)
	if err != nil {
		log.Printf("Operation GET USERS failed: %v\n", err)
		return nil, err
	}
	users := make([]data.User, 0)
	for _, _user := range _users {
		user := mapUserData(_user, false)
		users = append(users, user)
	}
	log.Printf("Operation GET USERS done. Total: %d\n", len(users))
	return users, nil
}

func (s *UserService) Update(user data.User) (data.User, error) {
	_user := mapUserDomain(user)
	updated, err := s.repo.Update(_user)
	if err != nil {
		log.Printf("Operation UPDATE USER failed: %v\n", err)
		return data.User{}, err
	}
	user = mapUserData(updated, false)
	log.Printf("Operation UPDATE USER done: %v\n", user)
	return user, nil
}

func (s *UserService) Delete(user data.User) error {
	if err := s.repo.Delete(user.Id); err != nil {
		log.Printf("Operation DELETE USER failed: %v\n", err)
		return err
	}
	log.Printf("Operation DELETE USER done: %v\n", user)
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
