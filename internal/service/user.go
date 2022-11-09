package service

import (
	"github.com/knvovk/copypass/internal/data"
	"github.com/knvovk/copypass/internal/domain"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	repo domain.UserRepository
	log  *logrus.Logger
}

func NewUserService(r domain.UserRepository, log *logrus.Logger) *UserService {
	return &UserService{
		repo: r,
		log:  log,
	}
}

func (s *UserService) Create(user data.User) (data.User, error) {
	passwordHash := user.Password
	_user := domain.User{
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: passwordHash,
	}
	inserted, err := s.repo.Insert(_user)
	if err != nil {
		s.log.Errorf("Operation CREATE USER failed: %v", err)
		return data.User{}, err
	}
	s.log.Infof("Operation CREATE USER done: %s", inserted.Id)
	user.Id = inserted.Id
	user.Password = ""
	return user, nil
}

func (s *UserService) Get(id string, unsafe bool) (data.User, error) {
	_user, err := s.repo.Find(id)
	if err != nil {
		s.log.Errorf("Operation GET USER failed: %v", err)
		s.log.Errorf("Requested id: %s", id)
		return data.User{}, err
	}
	user := data.User{
		Id:       _user.Id,
		Username: _user.Username,
		Email:    _user.Email,
	}
	if unsafe {
		user.Password = _user.PasswordHash
	}
	s.log.Debugf("Operation GET USER done: %v", user)
	return user, nil
}

func (s *UserService) GetAll(limit, offset int) ([]data.User, error) {
	_users, err := s.repo.FindAll(limit, offset)
	if err != nil {
		s.log.Errorf("Operation GET USERS failed: %v", err)
		return nil, err
	}
	users := make([]data.User, 0)
	for _, _user := range _users {
		user := data.User{
			Id:       _user.Id,
			Username: _user.Username,
			Email:    _user.Email,
		}
		users = append(users, user)
	}
	s.log.Debugf("Operation GET USERS done. Total: %d", len(users))
	return users, nil
}

func (s *UserService) Update(user data.User) (data.User, error) {
	_user := domain.User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
	updated, err := s.repo.Update(_user)
	if err != nil {
		s.log.Errorf("Operation UPDATE USER failed: %v", err)
		return data.User{}, err
	}
	user.Username = updated.Username
	user.Email = updated.Email
	user.Password = ""
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
