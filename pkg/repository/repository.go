package repository

import (
	"database/sql"

	"github.com/ranecsutu/testing/pkg/data"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllUsers() ([]*data.User, error)
	DeleteUser(id int) error
	GetUser(id int) (*data.User, error)
	GetUserByEmail(email string) (*data.User, error)
	InsertUser(user data.User) (int, error)
	InsertUserImage(i data.UserImage) (int, error)
	ResetPassword(id int, password string) error
	UpdateUser(u data.User) error
}
