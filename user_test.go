package main

import "testing"

func TestContect(t *testing.T) {
	Contect()
}

func TestCreateUserSchema(t *testing.T) {
	db := Contect()
	CreateRelationSchema(db)
}

func TestAddNewUser(t *testing.T) {
	db := Contect()
	err := CreateRelationSchema(db)
	if err == nil {
		err = CreateUserSchema(db)
		if err == nil {
			AddNewUser(db, "testuser")
		}
	}
}

func TestGetAllUsers(t *testing.T) {
	db := Contect()
	GetAllUsers(db)
}
