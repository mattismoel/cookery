package model

import (
	"testing"
)

func TestUserValidate(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr bool
	}{
		{
			name: "email: valid email",
			user: User{
				Id:           0,
				Email:        "test@mail.com",
				Firstname:    "John",
				Lastname:     "Doe",
				PasswordHash: "###############",
			},
			wantErr: false,
		},
		{
			name: "email: no @",
			user: User{
				Id:           0,
				Email:        "testmail.com",
				Firstname:    "John",
				Lastname:     "Doe",
				PasswordHash: "###############",
			},
			wantErr: true,
		},
		{
			name: "firstname: no firstname",
			user: User{
				Id:           0,
				Email:        "test@mail.com",
				Firstname:    "",
				Lastname:     "Doe",
				PasswordHash: "###############",
			},
			wantErr: true,
		},
		{
			name: "lastname: no lastname",
			user: User{
				Id:           0,
				Email:        "test@mail.com",
				Firstname:    "John",
				Lastname:     "",
				PasswordHash: "###############",
			},
			wantErr: true,
		},
		{
			name: "hash: no hash",
			user: User{
				Id:           0,
				Email:        "test@mail.com",
				Firstname:    "John",
				Lastname:     "Doe",
				PasswordHash: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Logf("case %q", tt.name)

		err := tt.user.Validate()

		if tt.wantErr && err == nil {
			t.Fatalf("wanted an error, got none")
		}

		if !tt.wantErr && err != nil {
			t.Fatalf("did not want an error. got %v", err)
		}

	}
}
