package model

import (
	"fmt"
	"strings"
)

type User struct {
	Id           int64  `json:"id"`
	Email        string `json:"email"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	PasswordHash string `json:"passwordHash,omitempty"`
}

func (u User) Validate() error {
	err := validateEmail(u.Email)
	if err != nil {
		return fmt.Errorf("could not validate email: %v", err)
	}

	if len(u.Firstname) <= 0 {
		return fmt.Errorf("no firstname.")
	}

	if len(u.Lastname) <= 0 {
		return fmt.Errorf("no lastname.")
	}

	if len(u.PasswordHash) <= 0 {
		return fmt.Errorf("no password hash.")
	}

	return nil
}

func validateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email is empty.")
	}

	if !strings.Contains(email, "@") {
		return fmt.Errorf("email does not contain any %q", "@")
	}

	if !strings.Contains(email, ".") {
		return fmt.Errorf("email does not contain any %q", ".")
	}

	return nil
}
