package domain

type Account struct {
	Id          string
	User        User
	Name        string
	Description string
	Url         string
	Login       string
	Password    string
}

type AccountRepository interface {
	Insert(account Account) (Account, error)
	Find(id string) (Account, error)
	FindByUser(user User) (Account, error)
	FindByName(name string) (Account, error)
	FindAll(limit, offset int) ([]Account, error)
	Update(user Account) (Account, error)
	Delete(id string) error
}
