package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var userPwPepper = "secret-random-string"

type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"_"` //- tells gorm to not save Password value into the DB.
	PasswordHash string `gorm:"not null"`
}

type UserService interface {
	ByID(id uint) *User
	ByEmail(email string) *User
	Authenticate(email, password string) *User
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
}

type UserGorm struct {
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
	/*
		ret := &User{}
		err := ug.DB.First(ret, id).Error
		switch err {
		case nil:
			return ret
		case gorm.ErrRecordNotFound:
			return nil
		default:
			panic(err)
		}
	*/
	return ug.byQuery(ug.DB.Where("id =?", id))
}

func (ug *UserGorm) ByEmail(email string) *User {
	/*
		ret := &User{}
		err := ug.DB.Where("email =?", email).First(ret).Error
		switch err {
		case nil:
			return ret
		case gorm.ErrRecordNotFound:
			return nil
		default:
			panic(err)
		}
	*/
	return ug.byQuery(ug.DB.Where("email =?", email))
}

func (ug *UserGorm) byQuery(query *gorm.DB) *User {
	ret := &User{}
	err := query.First(ret).Error
	switch err {
	case nil:
		return ret
	case gorm.ErrRecordNotFound:
		return nil
	default:
		panic(err)
	}
}

func (ug *UserGorm) Create(user *User) error {
	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password+userPwPepper), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	return ug.DB.Create(user).Error
}

func (ug *UserGorm) Update(user *User) error {
	return ug.DB.Save(user).Error
}

func (ug *UserGorm) Delete(id uint) error {
	user := &User{Model: gorm.Model{ID: id}}
	return ug.DB.Delete(user).Error
}

func (ug *UserGorm) DestructiveReset() {
	ug.DropTableIfExists(&User{})
	//ug.AutoMigrate(&User{})
	ug.AutoMigrate()
}

func (ug *UserGorm) AutoMigrate() {
	ug.DB.AutoMigrate(&User{})
}

func (ug *UserGorm) Authenticate(email, password string) *User {
	foundUser := ug.ByEmail(email)
	if foundUser == nil {
		//No user found with that email address
		return nil
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(foundUser.PasswordHash),
		[]byte(password+userPwPepper))
		if err != nil {
			//invalid password
			return nil
		}
		return foundUser
}
