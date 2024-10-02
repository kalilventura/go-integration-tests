package main

import "gorm.io/gorm"

type User struct {
	ID   int64  `gorm:"primary_key"`
	Name string `gorm:"type:varchar(255)"`
}

func (User) TableName() string {
	return "users"
}

type Repository struct {
	client *gorm.DB
}

func NewRepository(client *gorm.DB) *Repository {
	return &Repository{client: client}
}

func (itself *Repository) ListAll() ([]User, error) {
	var users []User

	err := itself.client.Model(&User{}).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
