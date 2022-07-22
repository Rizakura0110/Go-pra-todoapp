package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users(
		uuid,
		name,
		email,
		password,
		created_at) values (?, ? ,? , ?, ?)`
	_, err = Db.Exec(cmd,
		cretaeUUID(),
		u.Name,
		u.Email,
		Encrpt(u.Password),
		time.Now(),
	)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}