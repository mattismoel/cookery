package service

import (
	"context"
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

func (s UserService) Register(ctx context.Context, email, firstname, lastname, password, passwordRepeat string) error {
	if password != passwordRepeat {
		log.Fatalf("passwords do not match")
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("could not generate hash from password: %v", err)
	}

	fmt.Println(string(hashBytes))

	user := model.User{Email: email, Firstname: firstname, Lastname: lastname, PasswordHash: string(hashBytes)}

	err = user.Validate()
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.storage.AddUser(ctx, user)
	if err != nil {
		return fmt.Errorf("could not add user to storage: %v", err)
	}

	return nil
}

func (s UserService) Login(ctx context.Context, email, password string) (model.User, error) {
	usr, err := s.storage.UserByEmail(ctx, email)
	if err != nil {
		return model.User{}, fmt.Errorf("could not get user from storage: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.PasswordHash), []byte(password))
	if err != nil {
		return model.User{}, fmt.Errorf("passwords not comparable: %v", err)
	}

	return usr, nil
}
