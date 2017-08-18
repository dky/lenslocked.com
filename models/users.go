package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type User struct {
	gorm.Model
	Name string
	Email string `gorm:"not null;unique_index"`
}

//Provides methods for querying, creating and updating users.
type UserGorm struct {
	//db *gorm.DB
	*gorm.DB
}

func NewUserGorm(connectionInfo string) (*UserGorm, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	return &UserGorm{db}, nil
}

func (ug *UserGorm) ByID(id uint) *User {
	return nil
}

func (ug *UserGorm) ByEmail(email string) *User {
	return nil
}

func (ug *UserGorm) Create(user *User) {
}

func (ug *UserGorm) Update(user *User) {
}

func (ug *UserGorm) Delete(id uint) {
}
