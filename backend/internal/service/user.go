package service

import (
	"fmt"
	"log"

	"github.com/mattismoel/cookery/internal/model"
	"github.com/mattismoel/cookery/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	storage storage.Storage
}

func NewUserService(storage storage.Storage) *UserService {
	return &UserService{storage}
}

func (s UserService) Register(email, firstname, lastname, password, passwordRepeat string) error {
	if password != passwordRepeat {
		log.Fatalf("passwords do not match")
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("could not generate hash from password: %v", err)
	}

	user := model.User{Email: email, Firstname: firstname, Lastname: lastname, PasswordHash: string(hashBytes)}

	err = user.Validate()
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.storage.AddUser(user)
	if err != nil {
		return fmt.Errorf("could not add user to storage: %v", err)
	}

	return nil
}
