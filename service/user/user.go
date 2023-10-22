package user

import (
	entity "iprotector.github.com/entities/user"
)

type UserRepository interface {
	Get(id int) entity.User
}

// service user only cares about the repo implementation
func New(repo UserRepository) *User {
	return &User{
		repo: repo,
	}
}

type User struct {
	repo UserRepository
}

func (u *User) Get(id int) entity.User {
	return u.repo.Get(id)
}
