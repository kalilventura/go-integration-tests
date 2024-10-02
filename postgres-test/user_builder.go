package main

import "github.com/brianvoe/gofakeit/v7"

type UserBuilder struct{}

func (itself UserBuilder) Build() User {
	return User{
		Name: gofakeit.Name(),
	}
}

func (itself UserBuilder) BuildMany(quantity int) []User {
	var users []User

	for range quantity {
		users = append(users, itself.Build())
	}

	return users
}
