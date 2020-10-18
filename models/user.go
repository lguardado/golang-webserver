package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new users cannot have an Id")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return *user, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' was not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, user := range users {
		if user.ID == u.ID {
			users[i] = &u
			return *users[i], nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' was not found", u.ID)
}

func RemoveUserByID(id int) (string, error) {
	for k, v := range users {
		if v.ID == id {
			users = append(users[:k], users[k+1:]...)
			return "user removed", nil
		}
	}
	return "", fmt.Errorf("Cannot find user with id '%v'", id)
}
