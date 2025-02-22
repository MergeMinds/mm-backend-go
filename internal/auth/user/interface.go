package user

type Repo interface {
	Create(user *CreateModel) (*Model, error)
	GetByUsername(username string) (*Model, error)
	GetById(id int64) (*Model, error)
	DeleteById(id int64) error
}
