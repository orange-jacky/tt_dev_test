package main

func DbListAllUsers() ([]User, error) {
	db := Contect()
	users, err := GetAllUsers(db)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func DbCreateNewUser(name string) (User, error) {
	db := Contect()
	err := CreateRelationSchema(db)
	if err != nil {
		return User{}, err
	}
	err = CreateUserSchema(db)
	if err != nil {
		return User{}, err
	}
	user, err := AddNewUser(db, name)
	if err != nil {
		return user, err
	}
	return user, nil
}

func DbListUserAllRelationships(userid int64) ([]Relationship, error) {
	db := Contect()
	r, err := GetUserAllRelationships(db, userid)
	if err != nil {
		return nil, err
	}
	return r, nil

}

func DbInsertUserRelationshipState(userid int64, other_userid int64, state string) (Relationship, error) {
	db := Contect()
	err := CreateUserSchema(db)
	if err != nil {
		return Relationship{}, err
	}
	err = CreateRelationSchema(db)
	if err != nil {
		return Relationship{}, err
	}
	r, err := InsertUserRelationshipState(db, userid, other_userid, state)
	if err != nil {
		return r, err
	}
	return r, nil
}
