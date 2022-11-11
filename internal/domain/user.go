package domain

type User struct {
	Id           string
	Username     string
	Email        string
	PasswordHash string
}

type UserRepository interface {
	Insert(u User) (User, error)
	Find(id string) (User, error)
	FindByUsername(username string) (User, error)
	FindByEmail(email string) (User, error)
	FindAll(limit, offset int) ([]User, error)
	Update(u User) (User, error)
	Delete(id string) error
}
