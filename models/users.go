package models

import (
	"chat-app/db"
	user "chat-app/structs"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)


func RegisterUser(newUser user.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	if err != nil {
		return fmt.Errorf("could not hash password: %w", err)
	}

	newUser.Password = string(hashedPassword) 

	dbConn := db.DbConnect()
	defer dbConn.Close() 

	_, err = dbConn.Model(&newUser).Insert() 

	if err != nil {
		return fmt.Errorf("failed to insert user into the database: %w", err)
	}

	return nil

}

func LoginUser(loginData user.LoginUser) (*user.User, error) {
	var userDB user.User

	dbConn := db.DbConnect()
	defer dbConn.Close()

	err := dbConn.Model(&userDB).Where("email = ?", loginData.Email).Select()

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(loginData.Password))
	if err != nil {
		return nil, err
	}

	return &userDB, nil 
}
