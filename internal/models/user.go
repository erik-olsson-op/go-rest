package models

import (
	"context"
	"errors"
	"github.com/erik-olsson-op/go-rest/internal/database"
	"github.com/erik-olsson-op/go-rest/internal/util"
	"time"
)

type User struct {
	Id          int64       `json:"id"`
	Credentials Credentials `json:"credentials"`
	Registered  time.Time   `json:"registered"`
}

type Credentials struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=25"`
}

func (c *Credentials) Save() (int64, error) {
	query := "INSERT INTO user(email,password,registered) VALUES (?,?,?)"
	prepare, err := database.Connection.PrepareContext(context.Background(), query)
	if err != nil {
		return -1, err
	}
	defer prepare.Close()
	hash, err := util.HashPassword(c.Password)
	if err != nil {
		return -1, err
	}
	result, err := prepare.Exec(c.Email, hash, time.Now())
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func ValidateCredentials(c *Credentials) (int64, string, error) {
	query := "SELECT id,password FROM user WHERE email = ?"
	prepare, err := database.Connection.PrepareContext(context.Background(), query)
	if err != nil {
		return -1, "", err
	}
	defer prepare.Close()
	row := prepare.QueryRowContext(context.Background(), c.Email)
	var u User
	err = row.Scan(&u.Id, &u.Credentials.Password)
	if err != nil {
		return -1, "", err
	}

	ok := util.ValidatePassword(u.Credentials.Password, c.Password)
	if ok {
		return u.Id, c.Email, nil
	}

	return -1, "", errors.New("wrong username or password")
}
