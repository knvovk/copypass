package services

import (
	"github.com/knvovk/copypass/internal/dto"
	"github.com/knvovk/copypass/internal/models"
	"github.com/knvovk/copypass/internal/storages"
	"github.com/knvovk/copypass/pkg/passwd"
	"log"
)

type UserService struct {
	userStorage *storages.UserStorage
}

func NewUserService(storage *storages.UserStorage) *UserService {
	return &UserService{userStorage: storage}
}

func (s *UserService) Create(user dto.User) (dto.User, error) {
	passwordHash, err := passwd.HashPassword(user.Password)
	if err != nil {
		log.Printf("Operation CREATE USER failed: %v\n", err)
		return dto.User{}, err
	}
	_user := userModel(user)
	_user.PasswordHash = passwordHash
	inserted, err := s.userStorage.Insert(_user)
	if err != nil {
		log.Printf("Operation CREATE USER failed: %v\n", err)
		return dto.User{}, err
	}
	user = userDto(inserted, false)
	log.Printf("Operation CREATE USER done: %s\n", user.Id)
	return user, nil
}

func (s *UserService) GetOne(id string, unsafe bool) (dto.User, error) {
	_user, err := s.userStorage.Find(id)
	if err != nil {
		log.Printf("Operation GET USER failed: %v\n", err)
		log.Printf("Requested id: %s\n", id)
		return dto.User{}, err
	}
	user := userDto(_user, unsafe)
	log.Printf("Operation GET USER done: %v\n", user)
	return user, nil
}

func (s *UserService) GetMany(limit, offset int) ([]dto.User, error) {
	_users, err := s.userStorage.FindAll(limit, offset)
	if err != nil {
		log.Printf("Operation GET USERS failed: %v\n", err)
		return nil, err
	}
	users := make([]dto.User, 0)
	for _, _user := range _users {
		user := userDto(_user, false)
		users = append(users, user)
	}
	log.Printf("Operation GET USERS done. Total: %d\n", len(users))
	return users, nil
}

func (s *UserService) Update(user dto.User) (dto.User, error) {
	passwordHash, err := passwd.HashPassword(user.Password)
	if err != nil {
		log.Printf("Operation UPDATE USER failed: %v\n", err)
		return dto.User{}, err
	}
	user.Password = passwordHash
	_user := userModel(user)
	updated, err := s.userStorage.Update(_user)
	if err != nil {
		log.Printf("Operation UPDATE USER failed: %v\n", err)
		return dto.User{}, err
	}
	user = userDto(updated, false)
	log.Printf("Operation UPDATE USER done: %v\n", user)
	return user, nil
}

func (s *UserService) Delete(user dto.User) error {
	if err := s.userStorage.Delete(user.Id); err != nil {
		log.Printf("Operation DELETE USER failed: %v\n", err)
		return err
	}
	log.Printf("Operation DELETE USER done: %v\n", user)
	return nil
}

func userDto(user models.User, unsafe bool) dto.User {
	_user := dto.User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
	if unsafe {
		_user.Password = user.PasswordHash
	}
	return _user
}

func userModel(user dto.User) models.User {
	return models.User{
		Id:           user.Id,
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.Password,
	}
}
