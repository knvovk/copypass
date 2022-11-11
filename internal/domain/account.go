package domain

type Account struct {
	Id          string
	User        User
	Name        string
	Description string
	Url         string
	Username    string
	Password    string
}

type AccountRepository interface {
	Insert(a Account) (Account, error)
	Find(id string) (Account, error)
	FindByUser(user User) (Account, error)
	FindByName(name string) (Account, error)
	FindAll(limit, offset int) ([]Account, error)
	Update(a Account) (Account, error)
	Delete(id string) error
}
