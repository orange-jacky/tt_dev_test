package main

import (
	"fmt"
	"gopkg.in/pg.v4"
)

type Relationship struct {
	Id          int64  `json:"-"`
	OtherUserId int64  `json:"-"`
	UserId      int64  `json:"user_id,string"`
	State       string `json:"state"`
	Type        string `json:"type"`
}

func (r Relationship) String() string {
	return fmt.Sprintf("relationship<%d %d %s %s>", r.Id, r.UserId, r.State, r.Type)
}

func CreateRelationSchema(db *pg.DB) error {
	query := "CREATE TABLE  IF NOT EXISTS relationships (id serial, other_user_id bigint, user_id bigint, state text, type text)"
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func InsertUserRelationshipState(db *pg.DB, userid int64, other_userid int64, state string) (Relationship, error) {
	r1 := Relationship{}
	db.Model(&r1).
		Where("user_id = ?", userid).
		Where("other_user_id = ?", other_userid).
		Select()

	r2 := Relationship{}
	db.Model(&r2).
		Where("user_id = ?", other_userid).
		Where("other_user_id = ?", userid).
		Select()

	if r1.State == "" && r2.State == "" {
		r2.OtherUserId = userid
		r2.UserId = other_userid
		r2.State = state
		r2.Type = "relationship"
		err := db.Create(&r2)
		if err != nil { //create fail
			return r2, err
		}
		return r2, nil
	} else if r1.State == "" && r2.State != "" {
		r2.State = state
		db.Model(&r2).
			Set("state = ?", state).
			Where("id = ?", r2.Id).
			Update()
		return r2, nil

	} else if r1.State != "" {
		if r1.State == "liked" && state == "liked" {
			r2.OtherUserId = userid
			r2.UserId = other_userid
			r2.Type = "relationship"
			r2.State = "matched"
			_, err := db.Model(&r2).
				Create()
			if err != nil { //create fail
				return r2, err
			}
			return r2, nil
		} else {
			r2.State = state
			db.Model(&r2).
				Set("state = ?", state).
				Where("id = ?", r2.Id).
				Update()
			return r2, nil
		}

	}
	return r2, nil
}

func GetUserAllRelationships(db *pg.DB, userid int64) ([]Relationship, error) {
	var user User
	err := db.Model(&user).
		Column("user.*", "Relationships").
		Where("id = ?", userid).
		Select()
	if err != nil {
		return nil, err
	}
	for i, _ := range user.Relationships {
		user.Relationships[i].UserId = user.Relationships[i].OtherUserId
	}
	return user.Relationships, nil
}
