package memory

import (
	"fmt"
	"math/rand"

	"github.com/mattismoel/cookery/internal/model"
)

var users []model.User = []model.User{
	{Id: 0, Email: "johndoe@mail.com", Firstname: "John", Lastname: "Doe", PasswordHash: "12345678"},
}

func (s memoryStorage) UserById(id int64) (model.User, error) {
	for _, user := range users {
		return user, nil
	}

	return model.User{}, fmt.Errorf("could not find user with id: %d", id)
}

func (s memoryStorage) AddUser(u model.User) (int64, error) {
	users = append(users, u)
	id := int64(rand.Int())
	return id, nil
}
