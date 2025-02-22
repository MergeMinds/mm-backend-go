package session

import "github.com/google/uuid"

type Seconds = int

type Repo interface {
	Create(userId int64, lifetime Seconds) (*Model, error)
	GetById(id uuid.UUID) (*Model, error)
	DeleteById(id uuid.UUID) error
}
