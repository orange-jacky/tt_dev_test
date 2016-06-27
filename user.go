package main

import (
	"fmt"
	"gopkg.in/pg.v4"
)

type User struct {
	Id            int64          `json:"id,string"`
	Name          string         `json:"name"`
	Type          string         `json:"type"`
	Relationships []Relationship `json:"-"`
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %s>", u.Id, u.Name, u.Type)
}

func Contect() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "tantan",
		Database: "tantan",
	})
	return db
}

func CreateUserSchema(db *pg.DB) error {
	query := "CREATE TABLE  IF NOT EXISTS users (id serial, name text, type text)"
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func AddNewUser(db *pg.DB, name string) (User, error) {
	user := User{
		Name: name,
		Type: "user",
	}
	err := db.Create(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetAllUsers(db *pg.DB) ([]User, error) {
	var users []User
	err := db.Model(&users).Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}
