package user

import (
	"fmt"

	entity "iprotector.github.com/entities/user"
	"iprotector.github.com/store"
)

// user repo only cares about the implementation of the store
func New(store store.Storer) *User {
	return &User{
		store: store,
	}
}

type User struct {
	store store.Storer
}

func (u *User) Get(id int) entity.User {
	value := u.store.Get(id)
	fmt.Println(value)
	return entity.User{
		ID:   1,
		Name: "HI",
	}
}
